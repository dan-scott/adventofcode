import assert from "node:assert/strict";
import test from "node:test";
import * as day from "../day_12.mjs";

const input = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`;

test("Day 12 Part 1", () => {
  const result = day.part1(input);
  assert.strictEqual(result, 31);
});

test("Day 12 Part 2", () => {
  const result = day.part2(input);
  assert.strictEqual(result, 29);
});
