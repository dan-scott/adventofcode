import assert from "node:assert/strict";
import test from "node:test";
import * as day from "./day_06.mjs";

const inputs = [
  ["bvwbjplbgvbhsrlpgdmjqwftvncz", 5],
  ["nppdvjthqldpwncqszvftbrmjlhg", 6],
  ["nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10],
  ["zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11],
];

for (const [input, expected] of inputs) {
  test(`Day 6 Part 1 Example ${input}`, () => {
    const result = day.part1(input);
    assert.strictEqual(result, expected);
  });
}
