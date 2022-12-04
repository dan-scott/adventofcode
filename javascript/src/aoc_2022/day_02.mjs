export const dayNumber = 2;

//0 = R
//1 = P
//2 = S
const score = [
  [4, 8, 3],
  [1, 5, 9],
  [7, 2, 6],
];

// 0 = lose
// 1 = draw
// 2 = win
const decode = [
  [3, 4, 8],
  [1, 5, 9],
  [2, 6, 7],
];

const chrA = "A".charCodeAt(0);
const chrX = "X".charCodeAt(0);

export function part1(input) {
  let total = 0;
  for (const round of input.trim().split("\n")) {
    const elfHand = round.charCodeAt(0) - chrA;
    const yourHand = round.charCodeAt(2) - chrX;
    total += score[elfHand][yourHand];
  }
  return total;
}

export function part2(input) {
  let total = 0;
  for (const round of input.trim().split("\n")) {
    const elfHand = round.charCodeAt(0) - chrA;
    const strategy = round.charCodeAt(2) - chrX;
    total += decode[elfHand][strategy];
  }
  return total;
}
