package main

import "fmt"
import "regexp"

func main() {
  fmt.Println(ContainsDouble("ugknbfddgicrmpn"))
//  fmt.Println("ugknbfddgicrmopn", ContainsVowels("ugknbfddgicrmopn"))
}

// `func NiceString(input string) bool {
// `  return ContainsVowels(input) && ContainsDouble(input) && !ContainsNoughty(input)
// `}

func ContainsVowels(input string) bool {
  return len(regexp.MustCompile(`[aoeui]`).FindAll([] byte(input), 3)) == 3
}

func ContainsDouble(input string) bool {
  var char lastChar:= ""
  for _, char := range input {
    if lastChar == '' { lastChar = char }
  }
  return false
}

func ContainsNoughty(input string) bool {
  return false
}
