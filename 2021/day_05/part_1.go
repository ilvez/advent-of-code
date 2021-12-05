package main

import "fmt"
import aoc "aocgo/aochelper"
import "strings"

func main() {
  inputs := aoc.FileToLines("input2")
  for _, input := range inputs {
    start, end := ParseLine(input)
    fmt.Println("Result:", start, end)
  }
}

func ParseLine(line string) (Pos, Pos) {
  unparsedPositions := strings.Split(line, " -> ")
  return ParsePos(unparsedPositions[0]), ParsePos(unparsedPositions[1])
}

func ParsePos(unparsedPosition string) Pos {
  posStrings := strings.Split(unparsedPosition, ",")
  return Pos { aoc.StringToInt(posStrings[0]), aoc.StringToInt(posStrings[1])}
}

type Pos struct {
  x, y int
}
