import { parseMarkdown } from "./index";
import { readFileSync, writeFileSync } from "fs";
import path from "path";
import { Bv5, ICreatureTypeOptions } from "./classes";

const overwriteCases = false;

function dataFile(subfolder: string, filename: string): string {
  return path.join(
    path.dirname(path.dirname(module.filename)),
    "testdata",
    subfolder,
    filename
  );
}

describe("parsedCreature", () => {
  const parsedCreatureTests = [["owlbear.md", "owlbear.json"]];

  test.each(parsedCreatureTests)("%j should parse", (input, expected) => {
    const filename = dataFile("creatures", input);
    const data = readFileSync(filename, "utf8");
    const parsed = parseMarkdown<ICreatureTypeOptions & Bv5>(data);

    const checkFilename = dataFile("creatures", expected);
    if (overwriteCases) {
      writeFileSync(checkFilename, JSON.stringify(parsed, null, 2));
    }
    const checkData = JSON.parse(readFileSync(checkFilename, "utf8"));

    expect(parsed).toStrictEqual(checkData);
  });
});

describe("parsedSession", () => {
  const parsedSessionTests = [["20230119.md", "20230119.json"]];

  test.each(parsedSessionTests)("%j should parse", (input, expected) => {
    const filename = dataFile("session", input);
    const data = readFileSync(filename, "utf8");
    const parsed = parseMarkdown<ICreatureTypeOptions & Bv5>(data);

    const checkFilename = dataFile("session", expected);
    if (overwriteCases) {
      writeFileSync(checkFilename, JSON.stringify(parsed, null, 2));
    }
    const checkData = JSON.parse(readFileSync(checkFilename, "utf8"));

    expect(parsed).toStrictEqual(checkData);
  });
});
