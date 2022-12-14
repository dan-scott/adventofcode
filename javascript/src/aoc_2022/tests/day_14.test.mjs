import assert from "node:assert/strict";
import test from "node:test";
import * as day from "../day_14.mjs";

const input = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`;

test("Day 14 Part 1", () => {
  const result = day.part1(input);
  assert.strictEqual(result, 24);
});

test("Day 14 Part 2", () => {
  const result = day.part2(input);
  assert.strictEqual(result, 93);
});
