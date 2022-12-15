import { opendir } from "node:fs/promises";
import { join } from "node:path";
import { hrtime } from "node:process";
import { fileURLToPath } from "node:url";
import { getInput } from "./common/inputs.mjs";

const srcDir = fileURLToPath(new URL(".", import.meta.url));

export async function run({ year, days = [], benchMode = false }) {
  const results = await runDays(year, days);

  if (benchMode) {
    reportBenchmarks(results);
  } else {
    reportResults(results);
  }
}

async function getDaySolvers(year) {
  const path = join(srcDir, `aoc_${year}`);
  const yearDir = await opendir(path);
  const days = [];
  for await (const entry of yearDir) {
    if (entry.name.includes("test")) {
      continue;
    }

    const day = await import(join(path, entry.name));
    if (!day.dayNumber) {
      throw new Error("dayNumber not exported by", entry.name);
    }

    days.push(day);
  }

  days.sort((a, b) => a.dayNumber - b.dayNumber);
  return days;
}

async function runDays(year, days) {
  const solvers = await getDaySolvers(year);

  const results = [];

  for (const { skip, dayNumber, part1, part2 } of solvers) {
    if (skip) {
      console.log(`Skipping Day ${dayNumber}`);
      continue;
    }
    if (days.length && !days.includes(dayNumber)) {
      continue;
    }
    const input = await getInput(year, dayNumber);
    const p1start = hrtime.bigint();
    let p1Result;
    let p2Result;
    if (part1) {
      p1Result = part1(input);
    }
    const p1end = hrtime.bigint();
    if (part2) {
      p2Result = part2(input);
    }
    const p2end = hrtime.bigint();
    results.push({
      dayNumber,
      time: p2end - p1start,
      p1: {
        result: p1Result,
        time: p1end - p1start,
      },
      p2: {
        result: p2Result,
        time: p2end - p1end,
      },
    });
  }
  return results;
}

function reportBenchmarks(results) {
  const f = new Intl.NumberFormat();
  const format = (hrnum) => {
    if (typeof hrnum !== "bigint") {
      return hrnum + "";
    }
    let num = Number(hrnum);
    let unit = "us";
    num /= 1000;
    if (num > 1000) {
      unit = "ms";
      num /= 1000;
    }
    if (num > 1000) {
      unit = "s";
      num /= 1000;
    }
    return f.format(num) + unit;
  };

  const padTime = Math.max(
    ...results
      .flatMap(({ time, p1, p2 }) => [time, p1.time, p2.time])
      .map((t) => format(t).length)
  );

  const header =
    "|   Day | " +
    "Part 1".padStart(padTime, " ") +
    " | " +
    "Part 2".padStart(padTime, " ") +
    " | " +
    "Total".padStart(padTime, " ") +
    " |";

  const divider = "-".repeat(header.length);

  console.log(divider);
  console.log(header);
  console.log(divider);

  const printDay = ({ dayNumber, time, p1, p2 }) => {
    const line = [
      [dayNumber, 5],
      [p1.time, padTime],
      [p2.time, padTime],
      [time, padTime],
    ]
      .map(([v, p]) => format(v).padStart(p, " "))
      .join(" | ");
    console.log("|", line, "|");
  };

  results.forEach(printDay);

  console.log(divider);

  let totals = results.pop();
  totals = results.reduce((t, r) => {
    t.p1.time += r.p1.time;
    t.p2.time += r.p2.time;
    t.time += r.time;
    return t;
  }, totals);
  totals.dayNumber = "Total";
  printDay(totals);

  console.log(divider);
}

function reportResults(results) {
  console.log("------------------------------------------------------------");
  for (const { dayNumber, p1, p2 } of results) {
    console.log(`Day ${dayNumber}, Part 1`, p1.result);
    console.log(`Day ${dayNumber}, Part 2`, p2.result);
    console.log("------------------------------------------------------------");
  }
}
