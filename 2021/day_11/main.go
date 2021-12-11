package main

import(
  aoc "aocgo/aochelper"
  "fmt"
  _"os"
)

type Point struct {
  value int
}

type Map [][]Point

func main() {
  inputFile := "input2"
  // printPart1Solution(inputFile)
  // printPart2Solution(inputFile)

  m := parseMap(inputFile)
  printMap(&m)
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

func printMap(depthMap *Map) {
  for _, row := range(*(depthMap)) {
    for _, point := range(row) {
      printPoint(point)
    }
    fmt.Println()
  }
}

func printPoint(point Point) {
  var color string
  switch {
  case point.value <= 5: color = colorWhite
  case point.value < 9: color = colorYellow
  case point.value == 9 : color = colorRed
  case point.value == 0 : color = colorWhite
  default: color = colorGreen
  }
  fmt.Print(
    fmt.Sprint(
      color,
      point.value,
      colorReset,
    ),
  )
}

func printPart1Solution(inputFile string) {
  fmt.Println("Part 1 solution:", inputFile)
}

func printPart2Solution(inputFile string) {
  fmt.Println("Part 2 solution:", inputFile)
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
