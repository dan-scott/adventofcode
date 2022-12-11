import assert from "node:assert/strict";
import test from "node:test";
import * as day from "./day_11.mjs";

const input = `Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`;

test("Day 11 Part 1", () => {
  const result = day.part1(input);
  assert.strictEqual(result, 10605);
});

test("Day 11 Part 2", () => {
  const result = day.part2(input);
  assert.strictEqual(result, 2713310158);
});
