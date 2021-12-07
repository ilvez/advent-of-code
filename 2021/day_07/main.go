package main

import "fmt"
import aoc "aocgo/aochelper"


func main() {
  inputString := aoc.FileToLines("input2")
  inputArray := aoc.StringToIntArray(inputString[0])
  fmt.Println(inputArray)
}
