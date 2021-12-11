package main

import(
  aoc "aocgo/aochelper"
  "fmt"
  _"os"
)

type Point struct {
  value int
  illuminated bool
  loc Loc
}

type Map [][]Point
type Loc struct { x, y int }

var globalFlashCount int
var globalStepFlashCount int

func main() {
  inputFile := "input"

  printPart1Solution(inputFile)
  printPart2Solution(inputFile)
}

func iterateStep(m *Map) {
  globalStepFlashCount = 0
  for y, row := range(*m) {
    for x := range(row) {
      p := &(*m)[y][x]
      if increasePoint(p) {
        illuminateClose(m, (*p).loc)
      }
    }
  }
  // Reset illuminated since we have completed this step
  for y, row := range(*m) {
    for x := range(row) {
      p := &(*m)[y][x]
      p.illuminated = false
    }
  }
}

func increasePoint(p *Point) (flash bool) {
  if !p.illuminated {
    if p.value == 9 {
      p.value = 0
      p.illuminated = true
      globalFlashCount += 1
      globalStepFlashCount += 1
      return true
    }
    p.value += 1
  }
  return false
}

func illuminateClose(m *Map, dst Loc) {
  isWithinRowLower := dst.x - 1 >= 0
  isWithinRowUpper := dst.x + 1 < len((*m)[dst.y])
  if dst.y - 1 >= 0 {
    if isWithinRowLower {
      illuminateDirection(m, &(*m)[dst.y - 1][dst.x - 1])
    }

    illuminateDirection(m, &(*m)[dst.y - 1][dst.x])

    if isWithinRowUpper {
      illuminateDirection(m, &(*m)[dst.y - 1][dst.x + 1])
    }
  }

  if isWithinRowUpper {
    illuminateDirection(m, &(*m)[dst.y][dst.x + 1])
  }

  if dst.y + 1 < len(*m) {
    if isWithinRowUpper {
      illuminateDirection(m, &(*m)[dst.y + 1][dst.x + 1])
    }

    illuminateDirection(m, &(*m)[dst.y + 1][dst.x])

    if isWithinRowLower {
      illuminateDirection(m, &(*m)[dst.y + 1][dst.x - 1])
    }
  }

  if isWithinRowLower {
    illuminateDirection(m, &(*m)[dst.y][dst.x - 1])
  }
}

func illuminateDirection(m *Map, p *Point) {
  if !p.illuminated {
    if increasePoint(p) {
      illuminateClose(m, (*p).loc)
    }
  }
}

func parseMap(fileName string) (m Map) {
  globalFlashCount = 0
  lines := aoc.FileToLines(fileName)
  m = make([][]Point, len(lines))
  for i, line := range lines {
    m[i] = make([]Point, len(line))
    for j, pointString := range line {
      m[i][j] = Point {
        value: aoc.StringToInt(string(pointString)),
        loc: Loc { x: j, y: i },
      }
    }
  }
  return
}

func printMap(m *Map) {
  for _, row := range(*(m)) {
    for _, p := range(row) {
      printPoint(p)
    }
    fmt.Println()
  }
  fmt.Println("------")
}

func printPoint(p Point) {
  var color string
  switch {
  case p.value == 0 : color = colorRed
  case p.value <= 5: color = colorWhite
  case p.value < 9: color = colorGreen
  case p.value == 9 : color = colorYellow
  default: color = colorGreen
  }
  fmt.Print(
    fmt.Sprint(
      color,
      p.value,
      colorReset,
    ),
  )
}

func printPart1Solution(inputFile string) {
  m := parseMap(inputFile)
  for i := 0; i < 100; i++ {
    iterateStep(&m)
  }
  fmt.Println("Part 1 solution:", globalFlashCount)
}

func printPart2Solution(inputFile string) {
  m := parseMap(inputFile)
  mapSize := len(m) * len(m[0])
  i := 0
  for ; globalStepFlashCount < mapSize; i++ {
    iterateStep(&m)
  }
  fmt.Println("Part 2 solution:", i)
}

const (
  colorReset = "\033[0m"
  colorRed = "\033[31m"
  colorGreen = "\033[32m"
  colorYellow = "\033[33m"
  colorBlue = "\033[34m"
  colorPurple = "\033[35m"
  colorCyan = "\033[36m"
  colorWhite = "\033[37m"
)
