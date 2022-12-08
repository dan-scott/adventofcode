import assert from "node:assert/strict";
import test from "node:test";
import * as day from "./day_08.mjs";

test("Day 08 example", async (t) => {
  const input = `30373
25512
65332
33549
35390`;

  await t.test(`Part 1`, async () => {
    const result = day.part1(input);
    assert.strictEqual(result, 21);
  });

  await t.test(`Part 2`, async () => {
    const result = day.part2(input);
    assert.strictEqual(result, 8);
  });
});
