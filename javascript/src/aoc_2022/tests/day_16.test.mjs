import assert from "node:assert/strict";
import test from "node:test";
import * as day from "../day_16.mjs";

const input = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`;

test("Day 16 Part 1", () => {
  const result = day.part1(input, 10);
  assert.strictEqual(result, 1651);
});

// test("Day 16 Part 2", () => {
//   const result = day.part2(input, 20);
//   assert.strictEqual(result, 56000011);
// });
