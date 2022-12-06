export const dayNumber = 6;

/**
 *
 * @param {string} input
 */
export function part1(input) {
  for (let i = 4; i <= input.length; i++) {
    if (new Set(input.slice(i - 4, i)).size === 4) {
      return i;
    }
  }
  return -1;
}

export function part2(input) {
  for (let i = 14; i <= input.length; i++) {
    if (new Set(input.slice(i - 14, i)).size === 14) {
      return i;
    }
  }
  return -1;
}
