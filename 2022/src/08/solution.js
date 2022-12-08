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

console.log('Day 08 part 1:', visibleTrees.length);
