import { ld } from 'https://x.nest.land/deno-lodash@1.0.0/mod.ts';

const priority = (ch) => {
  if (ch == '') return 0;
  const asciiValue = ch.charCodeAt(0);
  if (asciiValue >= 97) return asciiValue - 96;
  else return asciiValue - 64 + 26;
};

const findCommon = (...packs) => {
  const common = ld.flatten(ld.intersection(...packs.map((p) => p.split(''))));
  if (common.length == 0) return 0;
  return common[0];
};

const groupByN = (n, data) => {
  const result = [];
  for (let i = 0; i < data.length; i += n) result.push(data.slice(i, i + n));
  return result;
};

const input = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)
  .filter((a) => a.length > 1);

const solution = (group) =>
  group.map((pack) => priority(findCommon(...pack))).reduce(
    (sum, current) => sum + current,
    0,
  );

const part1Input = input.map(
  (content) => [
    content.slice(0, content.length / 2),
    content.slice(content.length / 2, content.length),
  ],
);

console.log('Day 03 part 1:', solution(part1Input));
console.log('Day 03 part 2:', solution(groupByN(3, input)));
