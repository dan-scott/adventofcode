import { opendir } from "node:fs/promises";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import { dayNumber } from "./aoc_2022/day_04.mjs";
import { getInput } from "./common/inputs.mjs";

const __dirname = fileURLToPath(new URL(".", import.meta.url));

const allowList = [5];

async function runYear(year) {
  const path = join(__dirname, `aoc_${year}`);
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

    days[day.dayNumber - 1] = day;
  }

  console.log("---------------------------------------------------");
  for (const day of days) {
    if (!day) {
      continue;
    }
    if (day.skip) {
      console.log(`Skipping Day ${day.dayNumber}`);
      continue;
    }
    if (allowList.length && !allowList.includes(day.dayNumber)) {
      continue;
    }
    const input = await getInput(year, day.dayNumber);
    if (day.part1) {
      console.log(`Day ${day.dayNumber} part 1:`, day.part1(input));
    } else {
      console.log(`Day ${dayNumber} part 1 not found`);
    }
    if (day.part2) {
      console.log(`Day ${day.dayNumber} part 2:`, day.part2(input));
    } else {
      console.log(`Day ${dayNumber} part 2 not found`);
    }
    console.log("---------------------------------------------------");
  }
}

await runYear(2022);
