package main

import "fmt"
import "regexp"
import "strings"
import "os"
import "bufio"

func main() {
  niceStrings := 0
  for _, inputString := range FileToLines("input") {
    if NiceString(inputString) { niceStrings += 1 }
  }
  fmt.Println(niceStrings)
}

func NiceString(input string) bool {
  return ContainsVowels(input) && ContainsDouble(input) && !ContainsNoughty(input)
}

func ContainsVowels(input string) bool {
  return len(regexp.MustCompile(`[aoeui]`).FindAll([] byte(input), 3)) == 3
}

func ContainsDouble(input string) bool {
  var lastChar rune = 0
  for _, char := range input {
    if lastChar == char { return true }
    lastChar = char
  }
  return false
}

func ContainsNoughty(input string) bool {
  return strings.Contains(input, "ab") ||
    strings.Contains(input, "cd") ||
    strings.Contains(input, "pq") ||
    strings.Contains(input, "xy")
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
