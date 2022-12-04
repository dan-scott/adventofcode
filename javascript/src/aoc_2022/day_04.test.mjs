import assert from "node:assert/strict";
import test from "node:test";
import * as day from "./day_04.mjs";

test("Day 03 example", async (t) => {
  const input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`;

  await t.test("Part 1", async () => {
    const result = day.part1(input);
    assert.strictEqual(result, 2);
  });

  await t.test("Part 2", async () => {
    const result = day.part2(input);
    assert.strictEqual(result, 4);
  });
});
