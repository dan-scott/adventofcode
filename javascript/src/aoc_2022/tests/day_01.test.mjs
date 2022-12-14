import assert from "node:assert/strict";
import test from "node:test";
import * as day from "../day_01.mjs";
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

test("Day 1 Part 1", () => {
  const result = day.part1(input);
  assert.strictEqual(result, 24000);
});

test("Day 1 Part 2", () => {
  const result = day.part2(input);
  assert.strictEqual(result, 45000);
});
