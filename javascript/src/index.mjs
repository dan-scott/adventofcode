import { opendir } from "node:fs/promises";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import { getInput } from "./common/inputs.mjs";

const __dirname = fileURLToPath(new URL(".", import.meta.url));

async function runYear(year) {
  const path = join(__dirname, `aoc_${year}`);
  const yearDir = await opendir(path);
  for await (const entry of yearDir) {
    if (entry.name.includes("test")) {
      continue;
    }
    const { part1, part2, skip } = await import(join(path, entry.name));
    if (skip) {
      console.log(`Skipping ${entry.name}!`);
      return;
    }
    const day = parseInt(entry.name.split("_")[1].split(".")[0], 10);
    const input = await getInput(year, day);
    if (part1) {
      console.log(entry.name, "part 1:", part1(input));
    } else {
      console.log(entry.name, "part 1 not found.");
    }
    if (part2) {
      console.log(entry.name, "part 2:", part2(input));
    } else {
      console.log(entry.name, "part 2 not found.");
    }
    console.log("---------------------------------------------------");
  }
}

await runYear(2022);