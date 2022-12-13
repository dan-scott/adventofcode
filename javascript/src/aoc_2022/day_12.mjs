export const dayNumber = 12;

const chrA = "a".charCodeAt(0);

class Grid {
  #cells = new Map();
  #start;
  #end;
  constructor(rows, start, end) {
    this.#start = this.#id(start);
    this.#end = this.#id(end);
    for (let x = 0; x < rows[0].length; x++) {
      for (let y = 0; y < rows.length; y++) {
        const id = this.#id([x, y]);
        this.#cells.set(id, {
          id,
          pos: [x, y],
          height: rows[y].charCodeAt(x) - chrA,
          dist: Number.MAX_VALUE,
          prev: undefined,
        });
      }
    }
    this.#cells.get(this.#end).dist = 0;
  }

  #id([x, y]) {
    return `${x}-${y}`;
  }

  #neighbours({ pos: [x, y] }) {
    return [
      this.#id([x + 1, y]),
      this.#id([x - 1, y]),
      this.#id([x, y + 1]),
      this.#id([x, y - 1]),
    ]
      .map((id) => this.#cells.get(id))
      .filter((cell) => !!cell);
  }

  getShortestDistance(testHeight = false) {
    let leaves = [this.#end];
    while (leaves.length) {
      const leafKey = leaves.shift();
      const leaf = this.#cells.get(leafKey);
      for (const next of this.#neighbours(leaf)) {
        if (
          (next.height >= leaf.height || next.height === leaf.height - 1) &&
          next.dist > leaf.dist + 1
        ) {
          next.prev = leafKey;
          next.dist = leaf.dist + 1;
          if (
            (!testHeight && next.id === this.#start) ||
            (testHeight && next.height === 0)
          ) {
            return [next.id, next.dist];
          }
          leaves.push(next.id);
        }
      }
    }
    return ["", -1];
  }
}

/**
 *
 * @param {string} input
 */
function parseGrid(input) {
  const rows = input.trim().split("\n");
  let start;
  let end;
  for (let y = 0; y < rows.length; y++) {
    if (rows[y].includes("S")) {
      start = [rows[y].indexOf("S"), y];
      rows[y] = rows[y].replace(/S/, "a");
    }
    if (rows[y].includes("E")) {
      end = [rows[y].indexOf("E"), y];
      rows[y] = rows[y].replace(/E/, "z");
    }
  }
  return new Grid(rows, start, end);
}

export function part1(input) {
  const grid = parseGrid(input);
  const [, dist] = grid.getShortestDistance(false);
  return dist;
}

export function part2(input) {
  const grid = parseGrid(input);
  const [, dist] = grid.getShortestDistance(true);
  return dist;
}
