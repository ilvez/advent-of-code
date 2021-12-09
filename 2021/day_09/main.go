package main

import(
  aoc "aocgo/aochelper"
  "fmt"
  _ "strings"
)


func main() {
  parseMap("input2")
}

func parseMap(fileName string) {
  inputLines := aoc.FileToLines(fileName)
  fmt.Println(inputLines)
}
