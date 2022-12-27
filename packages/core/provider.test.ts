import { fail } from "assert";
import { randomUUID } from "crypto";
import { Bv5ProviderFileSystem } from "./provider";

describe("FileProvider", () => {
  const provider = new Bv5ProviderFileSystem("test_out");

  beforeAll(() => {
    const entries = provider.getEntries("Creature");
    if (entries.ok) {
      entries.success.forEach((entry) => {
        provider.removeEntry({
          id: entry.id,
          t: entry.t,
        });
      });
    } else {
      console.log(entries.err);
    }
  });

  test("Create and get entry", () => {
    const entryId = randomUUID();

    const entrySearch = provider.getEntry(entryId);
    if (entrySearch.ok) {
      fail();
    }
    expect(entrySearch.err).toBe("404");

    const entrySave = provider.saveEntry({
      id: entryId,
      t: "Creature",
    });
    if (!entrySave.ok) {
      fail(entrySave.err);
    }

    const entrySearch2 = provider.getEntry(entryId);
    if (!entrySearch2.ok) {
      fail(entrySearch2.err);
    }
    expect(entrySearch2.success.id).toBe(entryId);
    expect(entrySearch2.success.t).toBe("Creature");
    expect(entrySearch2.success.c).toBeUndefined();
  });
});
