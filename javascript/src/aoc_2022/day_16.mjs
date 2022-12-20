export const dayNumber = 16;

function parseGraph(input) {
  const nodeMap = new Map();
  for (const line of input.trim().split("\n")) {
    const [idPart, tunnelPart] = line.split("valve");
    const id = idPart.slice(6, 8);
    const flow = parseInt(idPart.split("rate=")[1].split(";")[0], 10);
    const linked = tunnelPart.slice(1).trim().split(", ");
    nodeMap.set(id, {
      id,
      flow,
      linked,
      paths: {},
    });
  }

  const flowNodes = new Set(
    [...nodeMap.values()].filter((n) => n.flow > 0).map((n) => n.id)
  );

  for (const node of nodeMap.values()) {
    const queue = node.linked.map((l) => ({
      id: l,
      length: 1,
      parents: [node.id],
    }));

    const found = new Set();

    while (found.size < flowNodes.size && queue.length) {
      const next = queue.shift();
      const parents = [...next.parents, next.id];
      if (!found.has(next.id) && flowNodes.has(next.id)) {
        node.paths[next.id] = {
          flow: nodeMap.get(next.id).flow,
          cost: next.length + 1,
          path: parents.slice(1),
        };
        found.add(next.id);
      }
      queue.push(
        ...nodeMap
          .get(next.id)
          .linked.filter((l) => parents.indexOf(l) < 0)
          .map((l) => ({
            id: l,
            length: next.length + 1,
            parents,
          }))
      );
    }
  }

  return nodeMap;
}

function getMaxPressure(current, timeLeft, flowNodes, nodeMap) {
  const node = nodeMap.get(current);
  let maxPressure = 0;

  for (const flowNode of flowNodes) {
    const nextFlows = new Set(flowNodes.values());
    nextFlows.delete(flowNode);
    const nextTime = timeLeft - node.paths[flowNode].cost;
    if (nextTime < 0) {
      continue;
    }
    const totalPressure =
      nextTime * node.paths[flowNode].flow +
      getMaxPressure(flowNode, nextTime, nextFlows, nodeMap);
    if (totalPressure > maxPressure) {
      maxPressure = totalPressure;
    }
  }
  return maxPressure;
}

export function part1(input) {
  const nodeMap = parseGraph(input);
  const flowNodes = new Set(
    [...nodeMap.values()].filter((f) => f.flow > 0).map((f) => f.id)
  );

  return getMaxPressure("AA", 30, flowNodes, nodeMap);
}

export function part2(input) {}
