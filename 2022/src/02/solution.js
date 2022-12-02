const abcToChoice = (a) => {
  switch (a) {
    case 'A':
      return 'rock';
    case 'B':
      return 'paper';
    case 'C':
      return 'scissors';
  }
};

const choicePoints = (a) => {
  switch (a) {
    case 'rock':
      return 1;
    case 'paper':
      return 2;
    case 'scissors':
      return 3;
  }
};

const roundToScore = (they, we) => {
  if (they == we) return choicePoints(we) + 3;
  else if (
    (they == 'rock' && we == 'paper') ||
    (they == 'paper' && we == 'scissors') ||
    (they == 'scissors' && we == 'rock')
  ) return choicePoints(we) + 6;
  else return choicePoints(we);
};

const part1XYZToABC = (a) => {
  switch (a) {
    case 'X': // rock
      return 'A';
    case 'Y': // paper
      return 'B';
    case 'Z': // scissors
      return 'C';
  }
};

const part1DecideResult = (pair) => {
  const they = abcToChoice(pair[0]);
  const we = abcToChoice(part1XYZToABC(pair[1]));

  return roundToScore(they, we);
};

const part2XYZToChoice = (they, xyz) => {
  switch (xyz) {
    case 'X': //lose
      if (they == 'rock') return 'scissors';
      if (they == 'paper') return 'rock';
      if (they == 'scissors') return 'paper';
      break;
    case 'Y': // draw
      return they;
    case 'Z': // win
      if (they == 'rock') return 'paper';
      if (they == 'paper') return 'scissors';
      if (they == 'scissors') return 'rock';
      break;
  }
};

const part2DecideResult = (pair) => {
  const they = abcToChoice(pair[0]);
  const we = part2XYZToChoice(they, pair[1]);

  return roundToScore(they, we);
};

const parsedInput = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)
  .filter((pair) => pair != '')
  .map((text) => text.split(/ /));

console.log(
  'Day 02 part 1:',
  parsedInput.map(part1DecideResult).reduce((sum, current) => sum + current, 0),
);
console.log(
  'Day 02 part 2:',
  parsedInput.map(part2DecideResult).reduce((sum, current) => sum + current, 0),
);
