//0 = R
//1 = P
//2 = S

let score = [
  [3, 6, 0],
  [0, 3, 6],
  [6, 0, 3],
];

// 0 = lose
// 1 = draw
// 2 = win
let decode = [
  [2, 0, 1],
  [0, 1, 2],
  [1, 2, 0],
];

const chrA = "A".charCodeAt(0);
const chrX = "X".charCodeAt(0);

export async function part1(input) {
  let total = 0;
  for (const round of input.split("\n")) {
    if (!round.length) {
      continue;
    }
    const elfHand = round.charCodeAt(0) - chrA;
    const yourHand = round.charCodeAt(2) - chrX;
    total += score[elfHand][yourHand] + yourHand + 1;
  }
  return total;
}

export async function part2(input) {
  let total = 0;
  for (const round of input.split("\n")) {
    if (!round.length) {
      continue;
    }
    const elfHand = round.charCodeAt(0) - chrA;
    const strategy = round.charCodeAt(2) - chrX;
    const yourHand = decode[elfHand][strategy];
    total += score[elfHand][yourHand] + yourHand + 1;
  }
  return total;
}
