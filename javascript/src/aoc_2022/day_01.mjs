import { getInput } from "../common/inputs.mjs";

async function getCalCounts() {
  const input = getInput(2022, 1);
  const counts = [];
  let calCt = 0;
  for (const line of (await input).split("\n")) {
    if (line.length === 0) {
      counts.push(calCt);
      calCt = 0;
    } else {
      calCt += parseInt(line, 10);
    }
  }

  return counts;
}

export async function part1() {
  const counts = await getCalCounts();
  return Math.max(...counts);
}

export async function part2() {
  const counts = await getCalCounts();
  counts.sort((a, b) => b - a);
  return counts[0] + counts[1] + counts[2];
}
