const decideOnABC = (they, we) => {
  const points = +(we.replace(/A/, 1).replace(/B/, 2).replace(/C/, 3));
  if (they == we) return points + 3;
  if (add(we) == they) return points;
  return points + 6;
};

const part1DecideResult = (pair) => {
  const we = pair[1].replace(/X/, 'A').replace(/Y/, 'B').replace(/Z/, 'C');

  return decideOnABC(pair[0], we);
};

const add = (
  choice,
) => (String.fromCharCode(choice.charCodeAt(0) + 1 > 67 ? 65 : choice.charCodeAt(0) + 1));
const sub = (
  choice,
) => (String.fromCharCode(choice.charCodeAt(0) - 1 < 65 ? 67 : choice.charCodeAt(0) - 1));

const part2XYZToChoice = (they, we) => {
  switch (we) {
    case 'X':
      return sub(they);
    case 'Y':
      return they;
    case 'Z':
      return add(they);
  }
};

const part2DecideResult = (pair) => {
  const we = part2XYZToChoice(pair[0], pair[1]);

  return decideOnABC(pair[0], we);
};

const parsedInput = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)
  .filter((pair) => pair != '')
  .map((text) => text.split(/ /));

const part1 = parsedInput.map(part1DecideResult).reduce((sum, current) => sum + current, 0);
const part2 = parsedInput.map(part2DecideResult).reduce((sum, current) => sum + current, 0);
console.log('Day 02 part 1:', part1, '-> correct:', part1 == 11906);
console.log('Day 02 part 2:', part2, '-> correct:', part2 == 11186);
