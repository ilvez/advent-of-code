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
const rope = [
  [0, 0],
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

for (const i in commands) {
  for (let step = 0; step < commands[i][1]; step += 1) {
    switch (commands[i][0]) {
      case 'R':
        rope[0] = [rope[0][0], rope[0][1] + 1];
        break;
      case 'U':
        rope[0] = [rope[0][0] - 1, rope[0][1]];
        break;
      case 'L':
        rope[0] = [rope[0][0], rope[0][1] - 1];
        break;
      case 'D':
        rope[0] = [rope[0][0] + 1, rope[0][1]];
        break;
    }
    for (let i = 0; i < 9; i += 1) {
      rope[i + 1] = newTailLocation(rope[i], rope[i + 1]);
    }
    part1Locations.push(ld.clone(rope[1]));
    part2Locations.push(ld.clone(rope[9]));
  }
}

console.log('Day 09 part 1:', ld.uniqWith(part1Locations, ld.isEqual).length);
console.log('Day 09 part 1:', ld.uniqWith(part2Locations, ld.isEqual).length);
