const input = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)
  .filter((line) => line.length > 0)
  .map(line => line.split(' '))

window.x = 1
window.cycleCount = 0
window.cumulativeSignalStrength = 0

const cycle = () => {
  cycleCount += 1
  if (cycleCount == 20 || cycleCount == 60 || cycleCount == 100 || cycleCount == 140 || cycleCount == 180 || cycleCount == 220) {
    cumulativeSignalStrength += cycleCount * x
  }
}
const processInstruction = (command) => {
  cycle()
  if (command[0] == 'addx') {
    cycle()
    x += +command[1]
  }
}

input.forEach(processInstruction)
console.log(cumulativeSignalStrength)
