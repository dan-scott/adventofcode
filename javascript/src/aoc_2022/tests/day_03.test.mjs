import assert from "node:assert/strict";
import test from "node:test";
import * as day from "../day_03.mjs";

const input = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`;

test("Day 3 Part 1", () => {
  const result = day.part1(input);
  assert.strictEqual(result, 157);
});

test("Day 3 Part 2", () => {
  const result = day.part2(input);
  assert.strictEqual(result, 70);
});
