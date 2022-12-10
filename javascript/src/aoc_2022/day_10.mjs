export const dayNumber = 10;

const cycleCt = new Map([
  ["noop", 1],
  ["addx", 2],
]);

class CPU {
  #code = [];
  #pc = 0;
  cycles = 1;
  x = 1;
  log = [];
  /**
   *
   * @param {string} input
   */
  constructor(input) {
    this.#code = input.trim().split("\n");
  }

  runProgram() {
    while (!this.done) {
      const [ins, arg] = this.#code[this.#pc].split(" ");
      this.#pc++;
      this.cycles += cycleCt.get(ins);
      this.log.push([ins, arg, this.x]);
      if (ins === "addx") {
        this.log.push([ins, arg, this.x]);
        this.x += parseInt(arg, 10);
      }
    }
  }

  get done() {
    return this.#pc === this.#code.length;
  }
}

export function part1(input) {
  const cpu = new CPU(input);
  cpu.runProgram();
  let total = 0;
  for (let i = 19; i < cpu.log.length; i += 40) {
    total += cpu.log[i][2] * (i + 1);
  }
  return total;
}

export function part2(input) {
  const cpu = new CPU(input);
  cpu.runProgram();
  let display = "";
  for (let cycle = 0; cycle < cpu.log.length; cycle++) {
    const hPos = cycle % 40;
    if (hPos === 0) {
      display += "\n";
    }
    const x = cpu.log[cycle][2];
    display += Math.abs(hPos - x) < 2 ? "#" : " ";
  }
  return display;
}
