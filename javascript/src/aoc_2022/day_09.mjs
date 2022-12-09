export const dayNumber = 9;

class Knot {
  x = 0;
  y = 0;
  tail = undefined;
  visited = new Set(["0-0"]);

  move({ x, y }) {
    this.x += x;
    this.y += y;
    this.visited.add(`${this.x}-${this.y}`);
    if (this.tail) {
      const xDiff = this.x - this.tail.x;
      const yDiff = this.y - this.tail.y;
      if (Math.abs(xDiff) < 2 && Math.abs(yDiff) < 2) {
        return;
      }
      const mv = {
        x: Math.max(-1, Math.min(1, xDiff)),
        y: Math.max(-1, Math.min(1, yDiff)),
      };
      this.tail.move(mv);
    }
  }
}

const directions = {
  R: { x: 1, y: 0 },
  U: { x: 0, y: -1 },
  L: { x: -1, y: 0 },
  D: { x: 0, y: 1 },
};

/**
 *
 * @param {string} input
 */
export function part1(input) {
  const head = new Knot();
  head.tail = new Knot();
  for (const instruction of input.trim().split("\n")) {
    const dir = directions[instruction[0]];
    const count = parseInt(instruction.slice(2), 10);
    for (let i = 0; i < count; i++) {
      head.move(dir);
    }
  }
  return head.tail.visited.size;
}

export function part2(input) {
  const head = new Knot();
  let tail = head;
  for (let i = 1; i < 10; i++) {
    tail.tail = new Knot();
    tail = tail.tail;
  }

  for (const instruction of input.trim().split("\n")) {
    const dir = directions[instruction[0]];
    const count = parseInt(instruction.slice(2), 10);
    for (let i = 0; i < count; i++) {
      head.move(dir);
    }
  }
  return tail.visited.size;
}
