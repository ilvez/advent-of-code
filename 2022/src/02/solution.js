const part1XYZToABC = a => {
  switch (a) {
    case 'X': return 'A'; // rock
    case 'Y': return 'B'; // paper
    case 'Z': return 'C'; // scissors
  }
}

const part1ABCToRockPaperScissors = a => {
  switch (a) {
    case 'A': return 'rock';
    case 'B': return 'paper';
    case 'C': return 'scissors';
  }
}

const part1ABCToScore = a => {
  switch(a) {
    case 'rock': return 1;
    case 'paper': return 2;
    case 'scissors': return 3;
  }
}
function decideResult(pair) {
  const they = part1ABCToRockPaperScissors(pair[0])
  const we = part1ABCToRockPaperScissors(part1XYZToABC(pair[1]))
  const theyScore = part1ABCToScore(they)
  const weScore = part1ABCToScore(we)

  let result = 0;
  if (they == we) {
    result = weScore + 3;
  } else if ((they == 'rock' && we == 'paper') || (they == 'paper' && we == 'scissors') || (they == 'scissors' && we == 'rock')) {
    result = weScore + 6;
  } else {
    result = weScore;
  }
  return result
}

const results = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)
  .filter(pair => pair != "")
  .map(text => text.split(/ /))
  .map(decideResult)
  .reduce((sum, current) => sum + current, 0)

console.log(results)
