import { ld } from 'https://x.nest.land/deno-lodash@1.0.0/mod.ts';

const input = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)
  .filter((a) => a.length > 1);

const fullRange = (line) => {
  const assignments = line.split(',');
  const a = assignments[0].split('-');
  const b = assignments[1].split('-');
  return [ld.range(+a[0], +a[1] + 1), ld.range(+b[0], +b[1] + 1)];
};

const fullyContains = (input) => {
  return input.filter((line) => {
    const range = fullRange(line);
    const intersection = ld.intersection(range[0], range[1]);
    return intersection.length == range[0].length || intersection.length == range[1].length;
  }).length;
};

const partiallyContains = (input) => {
  return input.filter((line) => {
    const range = fullRange(line);
    return ld.intersection(range[0], range[1]).length > 0;
  }).length;
};

console.log('Day 03 part 1:', fullyContains(input));
console.log('Day 03 part 1:', partiallyContains(input));
