export const dayNumber = 4;

function elfRange(elf) {
  const [s, e] = elf.split("-");
  return [parseInt(s, 10), parseInt(e, 10)];
}

function contains([as, ae], [bs, be]) {
  return as <= bs && ae >= be;
}

function inside(as, [bs, be]) {
  return bs <= as && be >= as;
}

export function part1(input) {
  let count = 0;
  for (const line of input.trim().split("\n")) {
    const [l, r] = line.split(",");
    const lr = elfRange(l);
    const rr = elfRange(r);
    if (contains(lr, rr) || contains(rr, lr)) {
      count++;
    }
  }
  return count;
}

export function part2(input) {
  let count = 0;
  for (const line of input.trim().split("\n")) {
    const [l, r] = line.split(",");
    const lr = elfRange(l);
    const rr = elfRange(r);
    if (inside(lr[0], rr) || inside(rr[0], lr)) {
      count++;
    }
  }
  return count;
}
