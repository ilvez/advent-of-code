package main

import "fmt"
import "os"
import "bufio"
import "strconv"
import "regexp"

func main() {
  inputs := FileToLines("input2")
  bingoNumbers := BingoStringToNumbers(inputs[0])
  fmt.Println("Result", bingoNumbers, BingoStringToNumbers(inputs[3]))
}

func BingoStringToNumbers(line string) []int {
  outputs := make([]int, 0)
  re := regexp.MustCompile(`\d+`)
  for _, numberString := range re.FindAllString(line, -1) {
    number, _ := strconv.Atoi(numberString)
    outputs = append(outputs, number)
  }
  return outputs
}

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

