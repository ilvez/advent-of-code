import { ld } from 'https://x.nest.land/deno-lodash@1.0.0/mod.ts';

const commands = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)
  .filter((line) => !line == '')
  .map((line) => line.split(' '));

const newTailLocation = (head, tail) => {
  const yDiff = head[0] - tail[0];
  const xDiff = head[1] - tail[1];
  if (yDiff < -1) {
    tail[0] -= 1;
    if (Math.abs(xDiff) == 1) tail[1] = head[1];
  } else if (yDiff > 1) {
    tail[0] += 1;
    if (Math.abs(xDiff) == 1) tail[1] = head[1];
  }

  if (xDiff < -1) {
    tail[1] -= 1;
    if (Math.abs(yDiff) == 1) tail[0] = head[0];
  } else if (xDiff > 1) {
    tail[1] += 1;
    if (Math.abs(yDiff) == 1) tail[0] = head[0];
  }

  return ld.clone(tail);
};

const part1Locations = [[0, 0]];
const part2Locations = [[0, 0]];
const tails = [
  [0, 0],
  [0, 0],
  [0, 0],
  [0, 0],
  [0, 0],
  [0, 0],
  [0, 0],
  [0, 0],
  [0, 0],
];
let head = [0, 0];

for (const i in commands) {
  for (let step = 0; step < commands[i][1]; step += 1) {
    switch (commands[i][0]) {
      case 'R':
        head = [head[0], head[1] + 1];
        break;
      case 'U':
        head = [head[0] - 1, head[1]];
        break;
      case 'L':
        head = [head[0], head[1] - 1];
        break;
      case 'D':
        head = [head[0] + 1, head[1]];
        break;
    }
    tails[0] = newTailLocation(head, tails[0]);
    for (let i = 1; i <= 8; i += 1) {
      tails[i] = newTailLocation(tails[i - 1], tails[i]);
    }
    part1Locations.push(ld.clone(tails[0]));
    part2Locations.push(ld.clone(tails[8]));
  }
}

console.log('Day 09 part 1:', ld.uniqWith(part1Locations, ld.isEqual).length);
console.log('Day 09 part 1:', ld.uniqWith(part2Locations, ld.isEqual).length);
