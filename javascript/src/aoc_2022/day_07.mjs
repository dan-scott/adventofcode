export const dayNumber = 7;

function calcDirSize(dir) {
  dir.size += dir.subDirs.reduce((t, d) => t + calcDirSize(d), 0);
  return dir.size;
}

/**
 *
 * @param {string[]} lines
 */
function parseTree(lines) {
  const stack = [];
  const lookup = new Map();
  let currentDir = { name: "/", subDirs: [], size: 0 };
  lookup.set(currentDir.name, currentDir);
  for (const line of lines.slice(1)) {
    if (line === "$ ls" || line.includes("dir")) {
      continue;
    } else if (line === "$ cd ..") {
      currentDir = stack.pop();
    } else if (line.includes("$ cd")) {
      const parent = currentDir;
      stack.push(currentDir);
      currentDir = {
        name: `${parent.name}${line.split(" ")[2]}/`,
        subDirs: [],
        size: 0,
      };
      parent.subDirs.push(currentDir);
      lookup.set(currentDir.name, currentDir);
    } else if (line.length) {
      let [s] = line.split(" ");
      currentDir.size += parseInt(s, 10);
    }
  }
  calcDirSize(lookup.get("/"));
  return lookup;
}

/**
 *
 * @param {string} input
 */
export function part1(input) {
  const lines = input.split("\n");
  const lookup = parseTree(lines);
  return [...lookup.values()]
    .filter((v) => v.size <= 100000)
    .reduce((t, d) => t + d.size, 0);
}

/**
 *
 * @param {string} input
 */
export function part2(input) {
  const lines = input.split("\n");
  const lookup = parseTree(lines);
  const requred = 30000000 - (70000000 - lookup.get("/").size);
  const candidates = [...lookup.values()].filter((v) => v.size >= requred);
  candidates.sort((a, b) => a.size - b.size);
  return candidates[0].size;
}
