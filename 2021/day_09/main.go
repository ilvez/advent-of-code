package main

import(
  aoc "aocgo/aochelper"
  "fmt"
)

type direction int

const (
  up direction = 1
  down direction = 0
)

type Point struct {
  value int
  n, s, e, w direction
  bottom bool
  basinSize int
}

type L struct { x, y int }
type Map [][]Point

func main() {
  depthMap := parseMap("input2")
  parseDirections(&depthMap)
  detectBottoms(&depthMap)
  printPart1Result(&depthMap)
  walkBasin(&depthMap, L {x: 9, y: 0}, 0)
  printMap(&depthMap)
}

func walkBasin(depthMap *Map, dst L, basinSize int) int {
  fmt.Println(dst, (*depthMap)[dst.y][dst.x].value, basinSize)
  (*depthMap)[dst.y][dst.x].basinSize = basinSize + 1
  if dst.y - 1 >= 0 && (*depthMap)[dst.y - 1][dst.x].basinSize == 0 && (*depthMap)[dst.y - 1][dst.x].value < 9  {
    basinSize = walkBasin(depthMap, L{x: dst.x, y: dst.y - 1}, basinSize)
  }
  if dst.x + 1 < len((*depthMap)[dst.y]) && (*depthMap)[dst.y][dst.x+1].basinSize == 0 && (*depthMap)[dst.y][dst.x+1].value < 9  {
    basinSize = walkBasin(depthMap, L{x: dst.x+1, y: dst.y}, basinSize)
  }
  if dst.y + 1 < len(*depthMap) && (*depthMap)[dst.y+1][dst.x].basinSize == 0 && (*depthMap)[dst.y+1][dst.x].value < 9  {
    basinSize = walkBasin(depthMap, L{x: dst.x, y: dst.y+1}, basinSize)
  }
  if dst.x - 1 >= 0 && (*depthMap)[dst.y][dst.x-1].basinSize == 0 && (*depthMap)[dst.y][dst.x-1].value < 9  {
    basinSize = walkBasin(depthMap, L{x: dst.x-1, y: dst.y}, basinSize)
  }
  (*depthMap)[dst.y][dst.x].basinSize = basinSize + 1
  return basinSize + 1
}

func printPart1Result(depthMap *Map) {
  riskLevel := 0
  for _, row := range(*depthMap) {
    for _, point := range(row) {
      if point.bottom {
        riskLevel += 1 + point.value
      }
    }
  }
  fmt.Println("Part 1 result:", riskLevel)
}

func detectBottoms(depthMap *Map) {
  for y := 0; y < len(*depthMap); y++ {
    for x := 0; x < len((*depthMap)[y]); x++ {
      point := (*depthMap)[y][x]
      if point.n == up && point.e == up && point.s == up && point.w == up {
        (*depthMap)[y][x].bottom = true
      }
    }
  }
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
      ", b:",
      point.basinSize,
      "] ",
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
