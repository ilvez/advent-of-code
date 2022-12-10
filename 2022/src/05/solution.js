import { ld } from 'https://x.nest.land/deno-lodash@1.0.0/mod.ts';

const lines = Deno.readTextFileSync(Deno.args[0]).split(/\n/);

const initializePiles = (amount) => {
  const piles = [];
  for (let i = 0; i < amount; i += 1) piles.push([]);
  return piles;
};

const maxStackSize = (lines) => {
  for (let i = 0; i < lines.length; i += 1) {
    if (lines[i] == '') return i - 1;
  }
};

const parsePiles = (lines) => {
  const stacksAmount = (lines[0].length + 1) / 4;
  const piles = initializePiles(stacksAmount);
  for (let lineNum = 0; lineNum < maxStackSize(lines); lineNum += 1) {
    for (let i = 0; i < stacksAmount; i += 1) {
      const linePos = 1 + (i * 4);
      if (lines[lineNum][linePos] != ' ') piles[i].push(lines[lineNum][linePos]);
    }
  }
  for(const p in piles) piles[p] = ld.reverse(piles[p])
  return piles;
};

const executeCrateMover9000 = (lines, piles) => {
  for (let lineNum = maxStackSize(lines) + 2; lineNum < lines.length; lineNum += 1) {
    const parts = lines[lineNum].split(' ');
    const amount = parts[1]
    const from = parts[3] - 1
    const to = parts[5] - 1
    for(let i = 0; i < amount; i+=1) {
      piles[to].push(piles[from].pop())
    }
  }
  return piles
};

const executeCrateMover9001 = (lines, piles) => {
  for (let lineNum = maxStackSize(lines) + 2; lineNum < lines.length; lineNum += 1) {
    const parts = lines[lineNum].split(' ');
    const amount = parts[1]
    const from = parts[3] - 1
    const to = parts[5] - 1

    const movedStack = []
    for(let i = 0; i < amount; i+=1) {
      movedStack.push(piles[from].pop())
    }
    piles[to] = ld.concat(piles[to], ld.reverse(movedStack))
  }
  return piles
};

console.log('Day 05 part 1:', executeCrateMover9000(lines, parsePiles(lines)).map(pile => pile.pop()).join(''))
console.log('Day 05 part 2:', executeCrateMover9001(lines, parsePiles(lines)).map(pile => pile.pop()).join(''))
