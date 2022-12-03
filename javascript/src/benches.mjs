import { opendir } from "node:fs/promises";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import { getInput } from "./common/inputs.mjs";
import Benchmark from "benchmark";

const __dirname = fileURLToPath(new URL(".", import.meta.url));

async function runYear(year) {
  const path = join(__dirname, `aoc_${year}`);
  const yearDir = await opendir(path);
  const suite = new Benchmark.Suite();

  for await (const entry of yearDir) {
    if (entry.name.includes("test")) {
      continue;
    }
    const dayMod = await import(join(path, entry.name));
    if (dayMod.skipBench) {
      console.log(`Skipping ${entry.name}!`);
      return;
    }
    const day = parseInt(entry.name.split("_")[1].split(".")[0], 10);
    const input = await getInput(year, day);
    if (dayMod.part1) {
      suite.add(`${year}, Day ${day}, Part 1`, function () {
        dayMod.part1(input);
      });
    }
    if (dayMod.part2) {
      suite.add(`${year}, Day ${day}, Part 2`, function () {
        dayMod.part2(input);
      });
    }
  }

  suite
    .on("cycle", function (event) {
      console.log(String(event.target));
    })
    .run();
}

await runYear(2022);
