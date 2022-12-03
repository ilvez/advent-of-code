import { ld } from 'https://x.nest.land/deno-lodash@1.0.0/mod.ts';

const priority = (ch) => {
  const asciiValue = ch.charCodeAt(0);
  if (asciiValue >= 97) return asciiValue - 96;
  else return asciiValue - 64 + 26;
};

const findCommon = (...packs) => {
  const common = ld.flatten(ld.intersection(...packs.map((p) => p.split(''))));
  return common.length == 1 ? common[0] : 0;
};

const groupByThree = (input) => {
  const result = [];
  for (let i = 0; i < input.length; i += 3) result.push(input.slice(i, i + 3));
  return result;
};

const groupIntoTwoPacks = (input) =>
  input.map(
    (content) => [
      content.slice(0, content.length / 2),
      content.slice(content.length / 2, content.length),
    ],
  );

const solution = (group) =>
  group.map((pack) => priority(findCommon(...pack))).reduce(
    (sum, current) => sum + current,
    0,
  );

const input = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)
  .filter((a) => a.length > 1);

console.log('Day 03 part 1:', solution(groupIntoTwoPacks(input)));
console.log('Day 03 part 2:', solution(groupByThree(input)));
