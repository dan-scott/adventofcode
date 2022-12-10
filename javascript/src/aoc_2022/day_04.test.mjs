import assert from "node:assert/strict";
import test from "node:test";
import * as day from "./day_04.mjs";

const input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`;

test("Day 4 Part 1", () => {
  const result = day.part1(input);
  assert.strictEqual(result, 2);
});

test("Day 4 Part 2", () => {
  const result = day.part2(input);
  assert.strictEqual(result, 4);
});
