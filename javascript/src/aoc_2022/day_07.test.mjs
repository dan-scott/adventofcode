import assert from "node:assert/strict";
import test from "node:test";
import * as day from "./day_07.mjs";

test("Day 07 example", async (t) => {
  const input = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`;

  await t.test(`Part 1`, async () => {
    const result = day.part1(input);
    assert.strictEqual(result, 95437);
  });

  await t.test(`Part 2`, async () => {
    const result = day.part2(input);
    assert.strictEqual(result, 24933642);
  });
});
