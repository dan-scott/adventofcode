export const dayNumber = 14;

class Cave {
  #fixed = new Set();

  #id([x, y]) {
    return `${x}-${y}`;
  }

  #next([x, y]) {
    return [
      [x, y + 1],
      [x - 1, y + 1],
      [x + 1, y + 1],
    ];
  }

  #isFixed(pos) {
    return pos[1] === this.lowPoint + 2 || this.#fixed.has(this.#id(pos));
  }

  constructor(input) {
    const lineCoords = input.trim().split("\n").map(parseLine);
    this.lowPoint = Math.max(...lineCoords.flatMap((l) => l.map((p) => p[1])));
    for (const line of lineCoords) {
      let current = line.shift();
      while (line.length) {
        let next = line.shift();
        let minX = Math.min(current[0], next[0]);
        let maxX = Math.max(current[0], next[0]);
        let minY = Math.min(current[1], next[1]);
        let maxY = Math.max(current[1], next[1]);
        for (let x = minX; x <= maxX; x++) {
          for (let y = minY; y <= maxY; y++) {
            this.#fixed.add(this.#id([x, y]));
          }
        }
        current = next;
      }
    }
  }

  fill1() {
    let restCount = 0;
    let finished = false;
    while (!finished) {
      let grain = [500, 0];
      let moving = true;
      while (moving) {
        moving = false;
        for (const pos of this.#next(grain)) {
          if (!this.#isFixed(pos)) {
            if (pos[1] >= this.lowPoint) {
              return restCount;
            }
            grain = pos;
            moving = true;
            break;
          }
        }
      }
      restCount++;
      this.#fixed.add(this.#id(grain));
    }
    return restCount;
  }

  fill2() {
    let restCount = 0;
    let finished = false;
    while (!finished) {
      let grain = [500, 0];
      let moving = true;
      while (moving) {
        moving = false;
        for (const pos of this.#next(grain)) {
          if (!this.#isFixed(pos)) {
            grain = pos;
            moving = true;
            break;
          }
        }
      }
      restCount++;
      this.#fixed.add(this.#id(grain));
      if (grain[0] === 500 && grain[1] === 0) {
        finished = true;
        return restCount;
      }
    }
    return restCount;
  }
}

function parseLine(line) {
  return line.split(" -> ").map(parsePair);
}

function parsePair(pair) {
  return pair.split(",").map((p) => parseInt(p, 10));
}

export function part1(input) {
  const cave = new Cave(input);
  return cave.fill1();
}

export function part2(input) {
  const cave = new Cave(input);
  return cave.fill2();
}
