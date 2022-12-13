export const dayNumber = 13;

function cmp(a, b) {
  let aT = typeof a;
  let bT = typeof b;

  if (aT !== bT) {
    if (aT === "number") {
      a = [a];
      aT = typeof a;
    } else {
      b = [b];
      bT = typeof b;
    }
  }
  if (aT === "number") {
    return a - b;
  }
  const aL = [...a];
  const bL = [...b];
  while (aL.length && bL.length) {
    let res = cmp(aL.shift(), bL.shift());
    if (res !== 0) {
      return res;
    }
  }
  return a.length - b.length;
}

function isInOrder(pairStr) {
  const [a, b] = pairStr.split("\n").map((p) => JSON.parse(p));
  return cmp(a, b) < 0;
}

export function part1(input) {
  return input
    .trim()
    .split("\n\n")
    .map(isInOrder)
    .reduce((t, is, i) => (is ? t + i + 1 : t), 0);
}

export function part2(input) {
  const d1 = [[2]];
  const d2 = [[6]];
  const msg = [
    d1,
    d2,
    ...input
      .trim()
      .replace(/\n\n/g, "\n")
      .split("\n")
      .map((l) => JSON.parse(l)),
  ];
  msg.sort(cmp);
  return (msg.indexOf(d1) + 1) * (msg.indexOf(d2) + 1);
}
