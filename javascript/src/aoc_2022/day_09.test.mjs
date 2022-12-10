import assert from "node:assert/strict";
import test from "node:test";
import * as day from "./day_09.mjs";

test(`Day 9 Part 1`, () => {
  const input1 = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`;
  const result = day.part1(input1);
  assert.strictEqual(result, 13);
});

test(`Day 9 Part 2`, () => {
  const input = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`;
  const result = day.part2(input);
  assert.strictEqual(result, 36);
});
