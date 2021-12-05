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
  lineCoordinates := ParseInputPositions(inputs)

  crossings := make([]Crossing, 0)
  for _, position := range lineCoordinates {
    line := DrawLine(position)
    for _, pos := range line {
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
    var length int
    if aoc.Abs(coords.start.x-coords.end.x) > aoc.Abs(coords.start.y-coords.end.y) {
      length = aoc.Abs(coords.start.x-coords.end.x)
    } else {
      length = aoc.Abs(coords.start.y-coords.end.y)
    }
    for i:=0 ; i <= length; i++ {
      x := DirectionFunc(coords.start.x, coords.end.x)(coords.start.x, i)
      y := DirectionFunc(coords.start.y, coords.end.y)(coords.start.y, i)
      line = append(line, Pos { x, y })
    }
  return line
}

func DirectionFunc(start, end int) func(int, int) int {
  if start < end {
    return Add
  } else if start == end {
    return Noop
  }
  return Sub
}

func Add(a, b int) int { return a + b }
func Sub(a, b int) int { return a - b }
func Noop(a, b int) int { return a }

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

