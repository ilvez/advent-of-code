const input = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)
  .filter((line) => line.length > 0)
  .map((line) => line.split(' '));

window.x = 1;
window.cycleCount = 0;
window.cumulativeSignalStrength = 0;

const cycle = () => {
  const line = Math.floor(cycleCount / 40);
  const pos = cycleCount % 40;
  if (x - 1 == pos || x == pos || x + 1 == pos) {
    screen[line][pos] = '#';
  }
  cycleCount += 1;
  if (
    cycleCount == 20 || cycleCount == 60 || cycleCount == 100 || cycleCount == 140 ||
    cycleCount == 180 || cycleCount == 220
  ) {
    cumulativeSignalStrength += cycleCount * x;
  }
};

const processInstruction = (command) => {
  cycle();
  if (command[0] == 'addx') {
    cycle();
    x += +command[1];
  }
};

const emptyScreen = () => {
  window.screen = [[], [], [], [], [], []];
  for (let i = 0; i < 6; i += 1) {
    for (let j = 0; j < 40; j += 1) {
      screen[i].push('.');
    }
  }
};

const printScreen = () => {
  for (let i = 0; i < 6; i += 1) {
    console.log(screen[i].join(''));
  }
};

emptyScreen();
input.forEach(processInstruction);
console.log(cumulativeSignalStrength);
printScreen();
