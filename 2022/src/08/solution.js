const trees = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)
  .filter((line) => line.length > 0)
  .map((line) => line.split('').map((number) => +number));

const rowsCount = trees.length;
const colsCount = trees[0].length;
const visibleTrees = [];
let previousTreeValue = -1;

const treeExists = (visibleTrees, row, col) => {
  return visibleTrees.filter((loc) => loc[0] == row && loc[1] == col).length >= 1;
};

for (let row = 0; row < rowsCount; row += 1) {
  for (let col = 0; col < colsCount; col += 1) {
    if (trees[row][col] > previousTreeValue) {
      previousTreeValue = trees[row][col];
      if (!treeExists(visibleTrees, row, col)) {
        visibleTrees.push([row, col]);
      }
    } else if (previousTreeValue >= 9) {
      col = colsCount;
      previousTreeValue = -1;
    }
  }
  previousTreeValue = -1;
}

for (let col = colsCount - 1; col >= 0; col -= 1) {
  for (let row = 0; row < rowsCount; row += 1) {
    if (trees[row][col] > previousTreeValue) {
      previousTreeValue = trees[row][col];
      if (!treeExists(visibleTrees, row, col)) {
        visibleTrees.push([row, col]);
      }
    } else if (previousTreeValue >= 9) {
      row = rowsCount;
      previousTreeValue = -1;
    }
  }
  previousTreeValue = -1;
}

for (let col = colsCount - 1; col >= 0; col -= 1) {
  for (let row = rowsCount - 1; row >= 0; row -= 1) {
    if (trees[row][col] > previousTreeValue) {
      previousTreeValue = trees[row][col];
      if (!treeExists(visibleTrees, row, col)) {
        visibleTrees.push([row, col]);
      }
    } else if (previousTreeValue >= 9) {
      row = 0;
      previousTreeValue = -1;
    }
  }
  previousTreeValue = -1;
}

for (let row = rowsCount - 1; row >= 0; row -= 1) {
  for (let col = colsCount - 1; col >= 0; col -= 1) {
    if (trees[row][col] > previousTreeValue) {
      previousTreeValue = trees[row][col];
      if (!treeExists(visibleTrees, row, col)) {
        visibleTrees.push([row, col]);
      }
    } else if (previousTreeValue >= 9) {
      col = 0;
      previousTreeValue = -1;
    }
  }
  previousTreeValue = -1;
}

const scenicValue = (tree, row, col) => {
  let upTrees = 0;
  let leftTrees = 0;
  let rightTrees = 0;
  let downTrees = 0;

  if (row > 0) {
    for (let rowValue = row - 1; rowValue >= 0; rowValue -= 1) {
      upTrees += 1;
      if (tree <= trees[rowValue][col]) break;
    }
  } else {
    upTrees = 0;
  }

  if (col > 0) {
    for (let colValue = col - 1; colValue >= 0; colValue -= 1) {
      leftTrees += 1;
      if (tree <= trees[row][colValue]) break;
    }
  } else {
    leftTrees = 0;
  }

  if (col < colsCount) {
    for (let colValue = col + 1; colValue < colsCount; colValue += 1) {
      rightTrees += 1;
      if (tree <= trees[row][colValue]) break;
    }
  } else {
    rightTrees = 0;
  }

  if (row < rowsCount) {
    for (let rowValue = row + 1; rowValue < rowsCount; rowValue += 1) {
      downTrees += 1;
      if (tree <= trees[rowValue][col]) {
        break;
      }
    }
  } else {
    downTrees = 0;
  }

  return upTrees * leftTrees * rightTrees * downTrees;
};

let maxScenicValue = -1;
for (let row = 0; row < rowsCount; row += 1) {
  for (let col = 0; col < colsCount; col += 1) {
    const currentValue = scenicValue(trees[row][col], row, col);
    if (currentValue > maxScenicValue) {
      maxScenicValue = currentValue;
    }
  }
}

console.log('Day 08 part 1:', visibleTrees.length);
console.log('Day 08 part 2:', maxScenicValue);
