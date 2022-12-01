import assert from "node:assert/strict";
import test from "node:test";
import { part1, part2 } from "./day_01.mjs";

test("Day 01 example", async (t) => {
  const input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`;

  await t.test("Part 1", async () => {
    const result = await part1(input);
    assert.strictEqual(result, 24000);
  });

  await t.test("Part 2", async () => {
    const result = await part2(input);
    assert.strictEqual(result, 45000);
  });
});
