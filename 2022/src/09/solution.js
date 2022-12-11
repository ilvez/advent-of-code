const commands = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)
  .filter(line => !line == '')
  .map(line => line.split(' '));

const updateHeadLocation = () => {
  if headLocation[]
  locations.push(tailLocation)
}

const locations = [[0,0]]
let tailLocation = [0,0]
let headLocation = [0,0]

for(const i in commands) {
  for(let step = 0; step < commands[i][1]; step += 1) {
    switch(commands[i][0]) {
      case 'R':
        headLocation = [headLocation[0], headLocation[1] + 1]
        break;
      case 'U':
        headLocation = [headLocation[0] - 1, headLocation[1]]
        break;
      case 'L':
        headLocation = [headLocation[0], headLocation[1] - 1]
        break;
      case 'D':
        headLocation = [headLocation[0] + 1, headLocation[1]]
        break;
    }
    updateHeadLocation()
  }
}
