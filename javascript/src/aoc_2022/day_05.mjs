export const dayNumber = 5;

function readInput(input) {
  const lines = input.split("\n");
  const stacks = [];
  while (lines.length) {
    const line = lines.shift();
    if (line[1] === "1") {
      break;
    }
    for (let i = 0; i < line.length; i += 4) {
      if (line[i] === "[") {
        const col = i / 4;
        if (!stacks[col]) {
          stacks[col] = [];
        }
        stacks[col].unshift(line[i + 1]);
      }
    }
  }
  lines.shift();
  const instructions = [];
  while (lines.length) {
    if (!lines[0].length) {
      break;
    }
    let [move, places] = lines.shift().split(" from ");
    move = parseInt(move.slice(5), 10);
    let [from, to] = places.split(" to ");
    from = parseInt(from, 10) - 1;
    to = parseInt(to, 10) - 1;
    instructions.push([move, from, to]);
  }
  return {
    stacks,
    instructions,
  };
}

export function part1(input) {
  const { stacks, instructions } = readInput(input);
  for (const [move, from, to] of instructions) {
    for (let i = 0; i < move; i++) {
      stacks[to].push(stacks[from].pop());
    }
  }
  return stacks.reduce((code, stack) => (code += stack.pop() ?? ""), "");
}

export function part2(input) {
  const { stacks, instructions } = readInput(input);
  for (const [move, from, to] of instructions) {
    const crates = stacks[from].splice(stacks[from].length - move);
    stacks[to].push(...crates);
  }
  return stacks.reduce((code, stack) => (code += stack.pop() ?? ""), "");
}
