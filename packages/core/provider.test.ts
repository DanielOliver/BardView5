import { fail } from "assert";
import { randomUUID } from "crypto";
import { Bv5ProviderFileSystem } from "./provider";

describe("FileProvider", () => {
  const provider = new Bv5ProviderFileSystem("test_out");

  beforeAll(() => {
    const entries = provider.getEntries("").as();
    if (entries.ok) {
      entries.value.forEach((entry) => {
        provider.removeEntry(entry);
      });
    } else {
      fail(entries.err);
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
      t: "session",
      name: "A new day",
    });
    if (!entrySave.ok) {
      fail(entrySave.err);
    }

    const entrySearch2 = provider.getEntry(entryId).as();
    if (!entrySearch2.ok) {
      fail(entrySearch2.err);
    }
    expect(entrySearch2.value.id).toBe(entryId);
  });

  test("Create and get entry in path", () => {
    const entryId = randomUUID();

    const entrySearch = provider.getEntry(entryId);
    if (entrySearch.ok) {
      fail();
    }
    expect(entrySearch.err).toBe("404");

    const entrySave = provider.saveEntry({
      id: entryId,
      t: "session",
      name: "Test",
    });
    if (!entrySave.ok) {
      fail(entrySave.err);
    }

    const entrySearch2 = provider.getEntry(entryId).as();
    if (!entrySearch2.ok) {
      fail(entrySearch2.err);
    }
    expect(entrySearch2.value.id).toBe(entryId);
  });
});
