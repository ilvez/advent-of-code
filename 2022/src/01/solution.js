import { SortService, Direction } from "https://deno.land/x/sort@v1.1.1/mod.ts"

const getCarryAmounts = (content) => {
  const result = []
  let current = 0
  content.forEach((item) => {
    if (item == 0) {
      result.push(current)
      current = 0
    } else {
      current += item
    }
  })
  return result
}
const numArray = Deno
  .readTextFileSync(Deno.args[0])
  .split(/\n/)
  .map((i) => +i);

const amounts = SortService.sort(getCarryAmounts(numArray), Direction.DESCENDING)
console.log(amounts[0])
console.log(amounts.slice(0, 3).reduce((sum, current) => sum + current, 0))
