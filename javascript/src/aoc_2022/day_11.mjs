export const dayNumber = 11;

function parseItems(itemsLine) {
  const [, nums] = itemsLine.split(": ");
  return nums.split(", ").map((n) => parseInt(n, 10));
}

function parseOp(opLine) {
  const [, op] = opLine.split("= ");
  return eval(`(old) => ${op}`);
}

function parseTest(testLine) {
  return parseInt(testLine.split("by ")[1], 10);
}

function parseCase(caseLine) {
  return parseInt(caseLine.split("monkey ")[1], 10);
}

class Monkey {
  constructor(input) {
    const [numL, itemsL, opL, testL, trueL, falseL] = input.trim().split("\n");
    this.num = parseInt(numL.split(" ")[1], 10);
    this.items = parseItems(itemsL);
    this.op = parseOp(opL);
    this.test = parseTest(testL);
    this.ifTrue = parseCase(trueL);
    this.ifFalse = parseCase(falseL);
    this.inspectCount = 0;
    this.worryLess = undefined;
  }

  catch(item) {
    this.items.push(item);
  }

  run(monkeyList) {
    while (this.items.length) {
      let next = this.items.shift();
      this.inspectCount++;
      next = this.op(next);
      next = this.worryLess ? next % this.worryLess : Math.floor(next / 3);
      if (next % this.test === 0) {
        monkeyList[this.ifTrue].catch(next);
      } else {
        monkeyList[this.ifFalse].catch(next);
      }
    }
  }
}

export function part1(input) {
  const monkeys = input.split("\n\n").map((i) => new Monkey(i));
  for (let i = 0; i < 20; i++) {
    for (const monkey of monkeys) {
      monkey.run(monkeys);
    }
  }
  monkeys.sort((a, b) => b.inspectCount - a.inspectCount);
  return monkeys[0].inspectCount * monkeys[1].inspectCount;
}

export function part2(input) {
  const monkeys = input.split("\n\n").map((i) => new Monkey(i));
  const lcm = monkeys.reduce((prod, m) => prod * m.test, 1);
  monkeys.forEach((m) => (m.worryLess = lcm));
  for (let i = 0; i < 10000; i++) {
    for (const monkey of monkeys) {
      monkey.run(monkeys);
    }
  }
  monkeys.sort((a, b) => b.inspectCount - a.inspectCount);
  return monkeys[0].inspectCount * monkeys[1].inspectCount;
}
