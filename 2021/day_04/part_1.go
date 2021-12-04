package main

import "fmt"
import "os"
import "bufio"
import "strconv"
import "regexp"

func main() {
  inputs := FileToLines("input2")
  bingoNumbers := BingoStringToNumbers(inputs[0])

  for i := 1; i <= NumberOfBoards(inputs); i++ {
    fmt.Println(ParseBoard(inputs, i))
  }

  fmt.Println("Result", bingoNumbers)
}

func ParseBoard(inputs []string, boardNumber int) [][]int {
  board := make([][]int, 0)
  boardEndLine := boardNumber * 5 + boardNumber
  for i := boardEndLine - 4; i <= boardEndLine; i++ {
    board = append(board, BingoStringToNumbers(inputs[i]))
  }
  return board
}

func NumberOfBoards(inputs []string) int {
  return (len(inputs) - 1) / 6
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

