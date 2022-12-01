import { SortService, Direction } from "https://deno.land/x/sort@v1.1.1/mod.ts"

const numArray = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n\n/)
  .map((text) => text.split(/\n/))
  .map((array) => array.map((i) => +i).reduce((sum, current) => sum + current, 0))

const amounts = SortService.sort(numArray, Direction.DESCENDING)
console.log(`Day 01 part 1: ${amounts[0]}`)
console.log(`Day 01 part 2: ${amounts.slice(0, 3).reduce((sum, current) => sum + current, 0)}`)
