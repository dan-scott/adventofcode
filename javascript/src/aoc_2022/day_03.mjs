const chra = "a".charCodeAt(0);
const chrA = "A".charCodeAt(0);

function codeToPoints(code) {
  if (code >= chra) {
    return code - chra + 1;
  } else {
    return code - chrA + 27;
  }
}

export function part1(input) {
  let points = new Array(53).fill(0);
  for (const rucksack of input.trim().split("\n")) {
    const l = rucksack.slice(0, rucksack.length / 2);
    const r = rucksack.slice(rucksack.length / 2);
    for (const item of l) {
      if (r.includes(item)) {
        points[codeToPoints(item.charCodeAt(0))]++;
        break;
      }
    }
  }
  return points.reduce((t, c, i) => t + c * i, 0);
}

export function part2(input) {
  let points = new Array(53).fill(0);
  const rucksacks = input.trim().split("\n");
  for (let i = 0; i < rucksacks.length; i += 3) {
    const [a, b, c] = rucksacks.slice(i, i + 3);
    for (const item of a) {
      if (b.includes(item) && c.includes(item)) {
        points[codeToPoints(item.charCodeAt(0))]++;
        break;
      }
    }
  }
  return points.reduce((t, c, i) => t + c * i, 0);
}
