package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {
  positionH := 0
  positionV := 0
  for _, inputString := range FileToLines("input") {
    direction := NewDirection(inputString)
    if direction.direction == "u" {
      positionV -= direction.step
    } else if direction.direction == "d" {
      positionV += direction.step
    } else if direction.direction == "f" {
      positionH += direction.step
    }
    fmt.Println(direction, positionH, positionV)
  }
  fmt.Println("Result", positionH * positionV)
}

type Direction struct {
  direction string
  step int
}

func NewDirection(inputString string) Direction {
  parts := strings.Split(inputString, " ")
  step, _ := strconv.Atoi(parts[1])
  return Direction { direction: string(parts[0][0]), step: step }
}

// Copypasted from internets
func FileToLines(filePath string) (lines []string) {
  file, err := os.Open(filePath)
  if err != nil {
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  err = scanner.Err()
  return
}

