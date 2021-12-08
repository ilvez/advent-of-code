package main

import(
  aoc "aocgo/aochelper"
  "fmt"
  "strings"
)


const (     //   ABCDEFG
  LEN_1 = 2 // 1   c  f
  LEN_2 = 5 // 2 a cde g
  LEN_3 = 5 // 3 a cd fg
  LEN_4 = 4 // 4  bcd f
  LEN_5 = 5 // 5 ab d fg
  LEN_6 = 6 // 6 ab defg
  LEN_7 = 3 // 7 a c  f
  LEN_8 = 7 // 8 abcdefg
  LEN_9 = 6 // 9 abcd fg
  LEN_0 = 6 // 0 abc efg
)

// Identifying steps:
// These are already identified: 1, 4, 7, 8
// C is only segment shared between 1 & 2
// F is missing only in 2
// 2 is identified with not having F in 5 segment numbers
// 5 is only 5 segment digit without C
// 3 is identified when 5 and 2 are identified
// 6 is only 6 segment digit without C
// E is only segment that 2 has but other 5 segment numbers don't have
// 9 is identified with not having E and being 6 segments
// 0 is identified when 9 and 6 are identified


func main() {
  var result int
  result = solveInput()
  fmt.Println(result)
}

func solveInput() (result int) {
  for _, inputLine := range aoc.FileToLines("input") {
    _, values := parseSignalsAndValuesString(inputLine)
    for _, valueString := range values {
      if parseValue(valueString) >= 0 {
        result += 1
      }
    }
  }
  return result
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

// Not useful notes:
// B is only char that 4 & 5 share which is not in other 5 segments numbers 2 & 3
// C is missing in 5 and 6
// A is diff between 1 and 7
// D is missing in 1, 7, 0
// B,E is are in only one of the 5 chars (2,3,7)
