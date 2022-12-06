import assert from "node:assert/strict";
import test from "node:test";
import * as day from "./day_06.mjs";

test("Day 06 example", async (t) => {
  const inputs = [
    ["bvwbjplbgvbhsrlpgdmjqwftvncz", 5],
    ["nppdvjthqldpwncqszvftbrmjlhg", 6],
    ["nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10],
    ["zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11],
  ];

  for (const [input, expected] of inputs) {
    await t.test(`Part 1 Example ${input}`, async () => {
      const result = day.part1(input);
      assert.strictEqual(result, expected);
    });
  }
});
