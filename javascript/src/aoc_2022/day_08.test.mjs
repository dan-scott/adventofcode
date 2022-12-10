import assert from "node:assert/strict";
import test from "node:test";
import * as day from "./day_08.mjs";

const input = `30373
25512
65332
33549
35390`;

test(`Day 8 Part 1`, () => {
  const result = day.part1(input);
  assert.strictEqual(result, 21);
});

test(`Day 8 Part 2`, () => {
  const result = day.part2(input);
  assert.strictEqual(result, 8);
});
