package main

import(
  aoc "aocgo/aochelper"
  "fmt"
  _ "strings"
)

type direction int

const (
  up direction = 1
  down direction = 0
)

type Point struct {
  value int
  n, s, e, w direction
}

type Map [][]Point

func main() {
  depthMap := parseMap("input")
  parseDirections(&depthMap)
  printPart1Result(&depthMap)
}


func printPart1Result(depthMap *Map) {
  riskLevel := 0
  for _, row := range(*depthMap) {
    for _, point := range(row) {
      if point.n == up && point.e == up && point.s == up && point.w == up {
        riskLevel += 1 + point.value
      }
    }
  }
  fmt.Println("Part 1 result:", riskLevel)
}

func parseDirections(depthMap *Map) {
  rows := len(*depthMap)
  for y := 0; y < rows; y++ {
    cols := len((*depthMap)[y])
    for x := 0; x < cols; x++ {
      currentPoint := (*depthMap)[y][x]
      if y - 1 >= 0 {
        if currentPoint.value < (*depthMap)[y-1][x].value {
          (*depthMap)[y][x].n = up
        }
      } else {
        (*depthMap)[y][x].n = up
      }
      if x + 1 < cols {
        if currentPoint.value < (*depthMap)[y][x+1].value {
          (*depthMap)[y][x].e = up
        }
      } else {
        (*depthMap)[y][x].e = up
      }
      if y + 1 < rows {
        if currentPoint.value < (*depthMap)[y+1][x].value {
          (*depthMap)[y][x].s = up
        }
      } else {
        (*depthMap)[y][x].s = up
      }
      if x - 1 >= 0 {
        if currentPoint.value < (*depthMap)[y][x-1].value {
          (*depthMap)[y][x].w = up
        }
      } else {
        (*depthMap)[y][x].w = up
      }
    }
  }
}

func parseMap(fileName string) (depthMap Map) {
  lines := aoc.FileToLines(fileName)
  depthMap = make([][]Point, len(lines))
  for i, line := range lines {
    depthMap[i] = make([]Point, len(line))
    for j, pointString := range line {
      depthMap[i][j] = Point {
        value: aoc.StringToInt(string(pointString)),
      }
    }
  }
  return
}

func printPoint(point Point) {
  fmt.Print(
    fmt.Sprint(
      "[",
      point.value,
      ", nesw:",
      point.n,
      point.e,
      point.s,
      point.w,
      "]",
    ),
  )
}

func printMap(depthMap *Map) {
  for _, row := range(*(depthMap)) {
    for _, point := range(row) {
      printPoint(point)
    }
    fmt.Println()
  }
}
