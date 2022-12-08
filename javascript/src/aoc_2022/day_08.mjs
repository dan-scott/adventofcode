export const dayNumber = 8;

const cp0 = "0".charCodeAt(0);

/**
 *
 * @param {string} input
 */
export function part1(input) {
  const { size, grid } = parseGrid(input);

  let visTrees = new Set();
  for (let i = 0; i < size; i++) {
    let maxDown = -1;
    let maxUp = -1;
    let maxLeft = -1;
    let maxRight = -1;
    let xDown = i;
    let xUp = i;
    let yLeft = i;
    let yRight = i;
    for (let j = 0; j < size; j++) {
      let yDown = j;
      let yUp = size - j - 1;
      let xLeft = j;
      let xRight = size - j - 1;
      if (grid[xDown][yDown] > maxDown) {
        visTrees.add(`x${xDown}y${yDown}`);
        maxDown = grid[xDown][yDown];
      }
      if (grid[xUp][yUp] > maxUp) {
        visTrees.add(`x${xUp}y${yUp}`);
        maxUp = grid[xUp][yUp];
      }
      if (grid[xLeft][yLeft] > maxLeft) {
        visTrees.add(`x${xLeft}y${yLeft}`);
        maxLeft = grid[xLeft][yLeft];
      }
      if (grid[xRight][yRight] > maxRight) {
        visTrees.add(`x${xRight}y${yRight}`);
        maxRight = grid[xRight][yRight];
      }
    }
  }

  return visTrees.size;
}

function parseGrid(input) {
  const lines = input.trim().split("\n");
  const size = lines.length;
  const grid = [...Array(size).keys()].map(() => []);
  for (let y = 0; y < size; y++) {
    for (let x = 0; x < size; x++) {
      grid[x][y] = lines[y].charCodeAt(x) - cp0;
    }
  }
  return { size, grid };
}

export function part2(input) {
  const { size, grid } = parseGrid(input);
  let maxVis = 0;
  for (let x = 1; x < size - 1; x++) {
    for (let y = 1; y < size - 1; y++) {
      const up = { found: false, dist: 0 };
      const down = { found: false, dist: 0 };
      const left = { found: false, dist: 0 };
      const right = { found: false, dist: 0 };
      for (let i = 1; i < size; i++) {
        if (!up.found && down.found && left.found && right.found) {
          break;
        }
        if (!up.found) {
          if (y - i >= 0) {
            up.dist = i;
            if (grid[x][y - i] >= grid[x][y]) {
              up.found = true;
            }
          } else {
            up.found = true;
          }
        }
        if (!down.found) {
          if (y + i < size) {
            down.dist = i;
            if (grid[x][y + i] >= grid[x][y]) {
              down.found = true;
            }
          } else {
            down.found = true;
          }
        }
        if (!left.found) {
          if (x - i >= 0) {
            left.dist = i;
            if (grid[x - i][y] >= grid[x][y]) {
              left.found = true;
            }
          } else {
            left.found = true;
          }
        }
        if (!right.found) {
          if (x + i < size) {
            right.dist = i;
            if (grid[x + i][y] >= grid[x][y]) {
              right.found = true;
            }
          } else {
            right.found = true;
          }
        }
      }
      const vis = up.dist * down.dist * left.dist * right.dist;
      if (vis > maxVis) {
        maxVis = vis;
      }
    }
  }
  return maxVis;
}
