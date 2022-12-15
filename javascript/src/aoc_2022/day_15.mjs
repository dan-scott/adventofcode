export const dayNumber = 15;

function id([x, y]) {
  return `${x}-${y}`;
}

function dist([x1, y1], [x2, y2]) {
  return Math.abs(x1 - x2) + Math.abs(y1 - y2);
}

function parsePart(part) {
  const [x, y] = part.split(", ");
  return [parseInt(x.substring(2), 10), parseInt(y.substring(2), 10)];
}

function parseInput(input) {
  let beacons = [];
  const sensors = [];
  for (const line of input.trim().split("\n")) {
    const [s, b] = line.split(": closest beacon");
    const bLoc = parsePart(b.split("at ")[1]);
    const sLoc = parsePart(s.split("at ")[1]);
    beacons.push(bLoc);
    sensors.push({
      loc: sLoc,
      dist: dist(sLoc, bLoc),
      beacon: bLoc,
    });
  }
  const mp = beacons.reduce((m, b) => {
    m.set(id(b), b);
    return m;
  }, new Map());
  beacons = [...mp.values()];
  let xs = [...beacons.map(([x]) => x), ...sensors.map(({ loc }) => loc[0])];
  let ys = [...beacons.map(([, y]) => y), ...sensors.map(({ loc }) => loc[0])];
  return {
    beacons: mp,
    sensors,
    minX: Math.min(...xs),
    maxX: Math.max(...xs),
    minY: Math.max(...ys),
    maxY: Math.max(...ys),
  };
}

function inSquare([x, y], sq) {
  return x >= 0 && x <= sq && y >= 0 && y <= sq;
}

export function part1(input, line = 2000000) {
  const { beacons, sensors, ...dims } = parseInput(input);
  let ct = 0;
  for (let x = dims.minX; x <= dims.maxX; x++) {
    const gridPos = [x, line];
    if (beacons.has(id(gridPos))) {
      continue;
    }
    if (sensors.filter((s) => dist(gridPos, s.loc) <= s.dist).length) {
      ct++;
    }
  }
  return ct;
}

export function part2(input, max = 4000000) {
  const { sensors } = parseInput(input);
  for (const sn of sensors) {
    const [x, y] = sn.loc;
    let s1 = [x + sn.dist + 1, y];
    let s2 = [x, y + sn.dist + 1];
    let s3 = [x - sn.dist - 1, y];
    let s4 = [x, y - sn.dist - 1];
    for (let i = 0; i <= sn.dist + 1; i++) {
      let check = [s1, s2, s3, s4].filter((s) => inSquare(s, max));
      if (check.length) {
        for (const sidePoint of check) {
          if (!sensors.filter((s) => dist(sidePoint, s.loc) <= s.dist).length) {
            return sidePoint[0] * 4000000 + sidePoint[1];
          }
        }
      }
      s1[0]--;
      s1[1]++;
      s2[0]--;
      s2[1]--;
      s3[0]++;
      s3[1]--;
      s4[0]++;
      s4[1]++;
    }
  }
}
