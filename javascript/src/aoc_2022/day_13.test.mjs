import assert from "node:assert/strict";
import test from "node:test";
import * as day from "./day_13.mjs";

const input = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`;

test("Day 13 Part 1", () => {
  const result = day.part1(input);
  assert.strictEqual(result, 13);
});

test("Day 13 Part 2", () => {
  const result = day.part2(input);
  assert.strictEqual(result, 140);
});
