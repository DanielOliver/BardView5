import path from "path";
import { Bv5 } from "./open/classes";
import { extractData } from "./open/index";
import { Err, Ok, Result } from "./result";
import glob from "glob";
import { mkdirSync, readFileSync, writeFileSync, rmSync } from "fs";

export interface IBv5Provider {
  removeEntry(item: Bv5): Result<number, string>;
  saveEntry(item: Bv5): Result<number, string>;
  getEntry(id: string): Result<Bv5, string>;
  getEntryFromPath(file: string): Result<Bv5, string>;
  getEntries(t: string): Result<Bv5[], string>;
}

export class Bv5ProviderFileSystem implements IBv5Provider {
  private _fullPath: string;

  constructor(location: string) {
    this._fullPath = path.resolve(location);
  }
  removeEntry(item: Bv5): Result<number, string> {
    try {
      const fullPath = path.join(this._fullPath, item.id + ".json");
      rmSync(fullPath);
      return Ok(1);
    } catch (error) {
      return Err(`${error}`);
    }
  }
  saveEntry(item: Bv5): Result<number, string> {
    const savePath = path.join(this._fullPath, item.id + ".json");
    try {
      mkdirSync(path.dirname(savePath), { recursive: true });
      writeFileSync(savePath, JSON.stringify(item, null, 2));
      return Ok(1);
    } catch (error) {
      return Err(`Failed to write file at ${savePath}; because ${error}`);
    }
  }
  getEntry(id: string): Result<Bv5, "404" | string> {
    try {
      const files = glob.sync(id + ".@(json|md)", {
        matchBase: true,
        absolute: true,
        cwd: this._fullPath,
      });

      if (files.length === 0) {
        return Err("404");
      }

      const fileContents = readFileSync(files[0], "utf-8");
      const ext = path.extname(files[0]);
      if (ext === ".json") {
        const json = JSON.parse(fileContents);
        return Ok({
          id: json.id,
          t: json.t,
          name: json.name,
        });
      }
      if (ext === ".md") {
        return Ok(extractData<Bv5>(fileContents));
      }

      return Err(`Unknown file extension ${ext}`);
    } catch (error) {
      return Err("Failed!");
    }
  }
  getEntryFromPath(file: string): Result<Bv5, "404" | string> {
    try {
      if (!path.isAbsolute(file)) {
        file = path.join(this._fullPath, file);
      }

      const fileContents = readFileSync(file, "utf-8");
      const ext = path.extname(file);
      if (ext === ".json") {
        const json = JSON.parse(fileContents);
        return Ok({
          id: json.id,
          t: json.t,
          name: json.name,
        });
      }
      if (ext === ".md") {
        return Ok(extractData<Bv5>(fileContents));
      }

      return Err(`Unknown file extension ${ext}`);
    } catch (error) {
      return Err("Failed!");
    }
  }
  getEntries(t: string): Result<Bv5[], string> {
    let globPath = "**/" + "*.@(json|md)";
    if (t.trim().length > 0 && !t.trim().startsWith("/")) {
      globPath = t.toLowerCase() + "/**/" + "*.@(json|md)";
    }
    const files = glob.sync(globPath, {
      matchBase: true,
      absolute: true,
      cwd: this._fullPath,
    });

    const failures: string[] = [];
    const successes: Bv5[] = [];
    files.forEach((filePath) => {
      this.getEntryFromPath(filePath).chain(
        (value) => {
          successes.push(value);
        },
        (err) => {
          failures.push(`Failure loading entry ${path} with ${err}`);
        }
      );
    });

    if (failures.length > 0) {
      return Err(failures.join("; "));
    }

    return Ok(successes);
  }
}
