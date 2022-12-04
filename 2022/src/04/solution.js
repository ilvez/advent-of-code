import { ld } from 'https://x.nest.land/deno-lodash@1.0.0/mod.ts';

const input = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)
  .filter((a) => a.length > 1);

const part1Solution = (input) => {
  return input.filter((line) => {
    const assignments = line.split(',')
    const a = assignments[0].split('-')
    const b = assignments[1].split('-')
    const aRange = ld.range(+a[0], +a[1] + 1)
    const bRange = ld.range(+b[0], +b[1] + 1)
    const intersection = ld.intersection(aRange, bRange)
    return intersection.length == aRange.length || intersection.length == bRange.length
  }).length
}

console.log('Day 03 part 1:', part1Solution(input));
