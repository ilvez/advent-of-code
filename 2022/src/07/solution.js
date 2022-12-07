import { ld } from 'https://x.nest.land/deno-lodash@1.0.0/mod.ts';

const input = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/);

window.lineNumber = -1;

const parseTree = (input) => {
  return treeTraverse(input, 0);
};

const treeTraverse = (input) => {
  lineNumber += 1;
  const currentTree = { size: 0 };
  for (; lineNumber < input.length; lineNumber += 1) {
    const line = input[lineNumber];
    const command = line.split(' ');
    if (command[1] == 'cd') {
      if (command[2] == '..') {
        return currentTree;
      } else {
        currentTree[command[2]] = treeTraverse(input);
        currentTree.size += currentTree[command[2]].size;
      }
    } else if (+command[0] > 0) {
      currentTree.size += +command[0];
    }
  }
  return currentTree;
};

function flattenObject(object, prefix = '', result = {}) {
  if (ld.isNumber(object)) {
    result[prefix] = object;
  } else {
    for (const element in object) flattenObject(object[element], prefix + '/' + element, result);
  }
  return result;
}

const fileTree = parseTree(input);

const flattenedArray = Object.entries(flattenObject(fileTree));
console.log(
  `Day 7 part 1:`,
  flattenedArray.filter((entry) => entry[1] < 100000).map((entry) => entry[1]).reduce((
    sum,
    current,
  ) => sum + current),
);
