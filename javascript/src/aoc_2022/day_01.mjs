function getCalCounts(input) {
  const counts = [];
  let calCt = 0;
  for (const line of input.split("\n")) {
    if (line.length === 0) {
      counts.push(calCt);
      calCt = 0;
    } else {
      calCt += parseInt(line, 10);
    }
  }

  return counts;
}

export function part1(input) {
  const counts = getCalCounts(input);
  return Math.max(...counts);
}

export function part2(input) {
  const counts = getCalCounts(input);
  counts.sort((a, b) => b - a);
  return counts[0] + counts[1] + counts[2];
}
