import * as fs from 'fs'
import * as path from 'path'
import * as _ from 'lodash'

function fileToStringArray(fileName: string):string[] {
  const projectDir = process.env.PWD || '.'
  const filePath = path.join(projectDir, fileName)
  return fs.readFileSync(filePath, 'utf-8').split(/\n/)
}

function getNumberArray(content:string[]):number[] {
  return content.map(function(item:string) {
    return +item
  })
}

function getCarryAmounts(content:number[]):number[] {
  var result:number[] = []
  var current:number = 0
  content.forEach((item:number) => {
    if (item == 0) {
      result.push(current)
      current = 0
    } else {
      current += item
    }
  })

  return result
}

const amounts:number[] = getCarryAmounts(getNumberArray(fileToStringArray('src/01/input')))

console.log(_.max(amounts))
console.log(_.sum(amounts.sort((a: number, b:number) => b - a).slice(0, 3)))
