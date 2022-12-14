import assert from "node:assert/strict";
import test from "node:test";
import * as day from "../day_02.mjs";

const input = `A Y
B X
C Z`;

test("Day 2 Part 1", () => {
  const result = day.part1(input);
  assert.strictEqual(result, 15);
});

test("Day 2 Part 2", () => {
  const result = day.part2(input);
  assert.strictEqual(result, 12);
});
