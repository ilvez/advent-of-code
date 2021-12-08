package main

import(
  aoc "aocgo/aochelper"
  "fmt"
  "strings"
)


const (
  LEN_1 = 2 // 1 cf
  LEN_2 = 5 // 2 acdeg
  LEN_3 = 5 // 3 acdfg
  LEN_4 = 4 // 4 bcdf
  LEN_5 = 5 // 5 abdfg
  LEN_6 = 6 // 6 abdefg
  LEN_7 = 3 // 7 acf
  LEN_8 = 7 // 8 abcdefg
  LEN_9 = 6 // 9 abcdfg
  LEN_0 = 6 // 0 abcefg
)

func main() {
  var result int
  for _, inputLine := range aoc.FileToLines("input") {
    _, values := parseSignalsAndValuesString(inputLine)
    for _, valueString := range values {
      if parseValue(valueString) >= 0 {
        result += 1
      }
    }
  }
  fmt.Println(result)
}

func parseValue(valueString string) (value int){
  switch len(valueString) {
    case LEN_1: value = 1
    case LEN_4: value = 4
    case LEN_7: value = 7
    case LEN_8: value = 8
    default: value = -1
  }
  return
}

func parseSignalsAndValuesString(input string) (signals, values []string) {
  splitted := strings.Split(input, " | ")
  signals = strings.Split(splitted[0], " ")
  values = strings.Split(splitted[1], " ")
  return
}
