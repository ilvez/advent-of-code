package main

import "fmt"
import aoc "aocgo/aochelper"
import "strings"

type Line []Pos

type LineCoordinates struct {
  start, end Pos
}

type Pos struct {
  x, y int
}

type Crossing struct {
  pos Pos
  count int
}

func main() {
  inputs := aoc.FileToLines("input")
  positions := ParseInputPositions(inputs)

  crossings := make([]Crossing, 0)
  for _, position := range positions {
    for _, pos := range DrawLine(position) {
      i := FindCrossing(crossings, pos)
      if (i < 0) {
        crossings = append(crossings, Crossing { pos, 0 } )
      } else {
        crossings[i] = Crossing { crossings[i].pos, crossings[i].count + 1 }
      }
    }
  }
  result := 0
  for _, crossing := range crossings {
    if crossing.count > 0 {
      result += 1
    }
  }
  fmt.Println("Result:", result)
}

func FindCrossing(crossings []Crossing, pos Pos) int {
  for i, existingCrossing := range crossings {
    if existingCrossing.pos == pos { return i }
  }
  return -1
}

func DrawLine(coords LineCoordinates) Line {
  line := make(Line, 0)
  if coords.start.x == coords.end.x {
    for i:=0 ; i <= aoc.Abs(coords.start.y-coords.end.y); i++ {
      y := findDirectionFunction(coords.start.y, coords.end.y)(coords.start.y, i)
      line = append(line, Pos { coords.start.x, y })
    }
  }
  if coords.start.y == coords.end.y {
    for i:=0 ; i <= aoc.Abs(coords.start.x-coords.end.x); i++ {
      x := findDirectionFunction(coords.start.x, coords.end.x)(coords.start.x, i)
      line = append(line, Pos { x, coords.start.y })
    }
  }
  return line
}

func findDirectionFunction(start, end int) func(int, int) int {
  if start <= end { return Add }
  return Sub
}

func Add(a, b int) int { return a + b }
func Sub(a, b int) int { return a - b }

func ParseInputPositions(inputs []string) (positions []LineCoordinates) {
  positions = make([]LineCoordinates, 0)
  for _, input := range inputs {
    start, end := ParseLine(input)
    positions = append(positions, LineCoordinates { start, end })
  }
  return
}

func ParseLine(line string) (Pos, Pos) {
  unparsedPositions := strings.Split(line, " -> ")
  return ParsePos(unparsedPositions[0]), ParsePos(unparsedPositions[1])
}

func ParsePos(unparsedPosition string) Pos {
  posStrings := strings.Split(unparsedPosition, ",")
  return Pos { aoc.StringToInt(posStrings[0]), aoc.StringToInt(posStrings[1])}
}

