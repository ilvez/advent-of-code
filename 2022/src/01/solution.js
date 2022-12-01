import * as fs from 'fs'
import * as path from 'path'
import * as _ from 'lodash'

function fileToStringArray(fileName) {
  const projectDir = process.env.PWD || '.'
  const filePath = path.join(projectDir, fileName)
  return fs.readFileSync(filePath, 'utf-8').split(/\n/)
}

function getNumberArray(content){
  return content.map(function(item) {
    return +item
  })
}

function getCarryAmounts(content) {
  result = []
  current = 0
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

const amounts = getCarryAmounts(getNumberArray(fileToStringArray('src/01/input')))

console.log(_.max(amounts))
console.log(_.sum(amounts.sort((a, b) => b - a).slice(0, 3)))
