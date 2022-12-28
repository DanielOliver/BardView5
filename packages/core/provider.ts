import path from "path";
import { Bv5Obj } from "./dnd5e/classes";
import { extractData } from "./dnd5e/index";
import { Err, Ok, Result } from "./result";
import glob from "glob";
import { mkdirSync, readFileSync, writeFileSync, rmSync } from "fs";

export interface IBv5Provider {
  removeEntry(item: Bv5Obj): Result<number, string>;
  saveEntry(item: Bv5Obj): Result<number, string>;
  getEntry(id: string): Result<Bv5Obj, string>;
  getEntries(t: string): Result<Bv5Obj[], string>;
}

export class Bv5ProviderFileSystem implements IBv5Provider {
  private _fullPath: string;

  constructor(location: string) {
    this._fullPath = path.resolve(location);
  }
  removeEntry(item: Bv5Obj): Result<number, string> {
    try {
      const fullPath = path.join(this._fullPath, item.t, item.id + ".json");
      rmSync(fullPath);
      return Ok(1);
    } catch (error) {
      return Err(`${error}`);
    }
  }
  saveEntry(item: Bv5Obj): Result<number, string> {
    const savePath = path.join(
      this._fullPath,
      item.t.toLowerCase(),
      item.id + ".json"
    );
    try {
      mkdirSync(path.dirname(savePath), { recursive: true });
      writeFileSync(savePath, JSON.stringify(item, null, 2));
      return Ok(1);
    } catch (error) {
      return Err(`Failed to write file at ${savePath}; because ${error}`);
    }
  }
  getEntry(id: string): Result<Bv5Obj, "404" | string> {
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
          c: json.c,
        });
      }
      if (ext === ".md") {
        return Ok(extractData<Bv5Obj>(fileContents));
      }

      return Err(`Unknown file extension ${ext}`);
    } catch (error) {
      return Err("Failed!");
    }
  }
  getEntries(t: string): Result<Bv5Obj[], string> {
    const files = glob.sync(t.toLowerCase() + "/**/" + "*.@(json|md)", {
      matchBase: true,
      absolute: true,
      cwd: this._fullPath,
    });

    return Ok(
      files.map((value) => ({
        id: path.parse(value).name,
        t: t,
      }))
    );
  }
}
