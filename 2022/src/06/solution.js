import { ld } from 'https://x.nest.land/deno-lodash@1.0.0/mod.ts';

const findMarker = (input, size) => {
  const inputArray = input.split('');
  for (let i = size; i <= inputArray.length; i++) {
    if (Object.keys(ld.groupBy(inputArray.slice(i - size, i))).length == size) {
      return i;
    }
  }
};

const input = Deno.readTextFileSync(Deno.args[0]);

console.log(`Day 6 part 1: ${findMarker(input, 4)}`);
console.log(`Day 6 part 2: ${findMarker(input, 14)}`);
