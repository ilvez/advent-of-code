package main

import(
  aoc "aocgo/aochelper"
  "fmt"
)
func main() {
  inputFile := "input2"
  lines := aoc.FileToLines(inputFile)
  fmt.Println(lines)
}

type Edge struct {
  a, b string
}
