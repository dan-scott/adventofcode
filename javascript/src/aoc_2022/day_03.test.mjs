import assert from "node:assert/strict";
import test from "node:test";
import * as day from "./day_03.mjs";

test("Day 03 example", async (t) => {
  const input = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`;

  await t.test("Part 1", async () => {
    const result = day.part1(input);
    assert.strictEqual(result, 157);
  });

  await t.test("Part 2", async () => {
    const result = day.part2(input);
    assert.strictEqual(result, 70);
  });
});
