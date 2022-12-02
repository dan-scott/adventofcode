import assert from "node:assert/strict";
import test from "node:test";
import * as day from "./day_02.mjs";

test("Day 02 example", async (t) => {
  const input = `A Y
B X
C Z`;

  await t.test("Part 1", async () => {
    const result = await day.part1(input);
    assert.strictEqual(result, 15);
  });

  await t.test("Part 2", async () => {
    const result = await day.part2(input);
    assert.strictEqual(result, 12);
  });
});
