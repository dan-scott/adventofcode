import assert from "node:assert/strict";
import test from "node:test";
import * as day from "../day_05.mjs";

const input = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`;

test("Day 5 Part 1", () => {
  const result = day.part1(input);
  assert.strictEqual(result, "CMZ");
});

test("Day 5 Part 2", () => {
  const result = day.part2(input);
  assert.strictEqual(result, "MCD");
});
