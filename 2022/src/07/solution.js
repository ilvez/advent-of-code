const input = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)

window.lineNumber = -1

const parseTree = (input) => {
  return treeTraverse(input, 0)
}

const treeTraverse = (input) => {
  lineNumber += 1
  const currentTree = {size: 0}
  for (; lineNumber < input.length; lineNumber += 1) {
    const line = input[lineNumber]
    const command = line.split(' ')
    if (command[1] == "cd") {
      if (command[2] == '..') {
        return currentTree
      } else {
        currentTree[command[2]] = treeTraverse(input)
        currentTree.size += currentTree[command[2]].size
      }
    } else if (+command[0] > 0) {
      currentTree.size += +command[0]
    }
  }
  return currentTree
}

const directories = (directory) => {
  return Object
    .keys(directory).filter(entry => entry != 'size')
    .reduce((obj, subDirectory) => {
      return { ...obj, [subDirectory]: directory[subDirectory].size }
    }, {})
}

const flatten = (tree) => {
  const dirs = directories(tree)
  const subDirectories = Object.keys(dirs).flatMap((subDir) => flatten(tree[subDir]))
  console.log(subDirectories)
  return {
    ...dirs,
    ...subDirectories[0]
  }
}

const fileTree = parseTree(input)

console.log(`Day 7 part 1:`, Object.entries(flatten(fileTree)).filter(entry => entry[1] < 100000).map(entry => entry[1]).reduce((sum, current) => sum + current))
