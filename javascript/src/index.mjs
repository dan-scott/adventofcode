import { opendir } from "node:fs/promises";
import { fileURLToPath } from "node:url";
import { run } from "./runner.mjs";

const srcDir = fileURLToPath(new URL(".", import.meta.url));

const args = process.argv.slice(2);

const num = new RegExp(/\d+/);

const validArgs = ["--year", "-y", "--day", "-d", "--bench", "-b"];

const years = [];
const days = [];
let benchMode = false;

while (args.length) {
  const arg = args.shift();
  if (!validArgs.includes(arg)) {
    console.log(
      `usage: ${process.argv[0]} ${process.argv[1]} [-y|--year (number)]* [-d|--day (number)]*`
    );
    process.exit(1);
  }
  if (arg === "-b" || arg === "--bench") {
    benchMode = true;
    continue;
  }
  while (num.test(args[0])) {
    const val = parseInt(args.shift(), 10);
    switch (arg) {
      case "-y":
      case "--year":
        years.push(val);
        break;
      case "-d":
      case "--day":
        days.push(val);
        break;
    }
  }
}

if (!years.length) {
  const yearDir = await opendir(srcDir);
  for await (const ent of yearDir) {
    if (!ent.name.includes("aoc_")) {
      continue;
    }
    years.push(parseInt(ent.name.split("_")[1], 10));
  }
  years.sort((a, b) => a - b);
}

for (const year of years) {
  await run({ year, days, benchMode });
}
