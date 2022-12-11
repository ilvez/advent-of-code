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

const locations = [[0, 0]];
let tail = [0, 0];
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
    locations.push(newTailLocation(head, tail));
  }
}

console.log('Day 09 part 1:', ld.uniqWith(locations, ld.isEqual).length);
