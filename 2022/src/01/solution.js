const numArray = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n\n/)
  .map((text) => text.split(/\n/))
  .map((array) => array.map((i) => +i).reduce((sum, current) => sum + current, 0));

const amounts = numArray.sort((a, b) => b - a);
console.log('Day 01 part 1:', amounts[0]);
console.log('Day 01 part 2:', amounts.slice(0, 3).reduce((sum, current) => sum + current, 0));
