import assert from "node:assert/strict";
import test from "node:test";
import * as day from "./day_05.mjs";

test("Day 05 example", async (t) => {
  const input = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`;

  await t.test("Part 1", async () => {
    const result = day.part1(input);
    assert.strictEqual(result, "CMZ");
  });

  await t.test("Part 2", async () => {
    const result = day.part2(input);
    assert.strictEqual(result, "MCD");
  });
});
