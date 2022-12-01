import { open } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = fileURLToPath(new URL(".", import.meta.url));

function buildPath(year, number) {
  return path.join(
    __dirname,
    "../../../",
    "inputs",
    `${year}`,
    `${number}.txt`
  );
}

export async function getInput(year, number) {
  const path = buildPath(year, number);
  let fh;
  try {
    fh = await open(path, "r");
    const buf = await fh.readFile();
    return buf.toString();
  } finally {
    fh?.close();
  }
}
