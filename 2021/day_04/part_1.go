package main

import "fmt"
import "os"
import "bufio"
import "strconv"
import "regexp"

func main() {
  inputs := FileToLines("input")
  bingoNumbers := BingoStringToNumbers(inputs[0])

  bestBoardPosition := 99999
  winningResult := 0

  for i := 1; i <= NumberOfBoards(inputs); i++ {
    board := ParseBoard(inputs, i)
    position, result := SolveBoard(board, bingoNumbers)
    if position < bestBoardPosition {
      bestBoardPosition = position
      winningResult = result
    }
  }

  fmt.Println("Result:", winningResult)
}

func SolveBoard(board [][]int, bingoNumbers []int) (winningPosition int, solution int) {
  markedColumns := make([]int, 5)
  markedRows := make([]int, 5)
  for position, bingoNumber := range bingoNumbers {
    rowPos, colPos := NumberPositionOnBoard(board, bingoNumber)
    if rowPos == -1 || colPos == -1 { continue }
    markedRows[rowPos] += 1
    markedColumns[colPos] += 1
    board[rowPos][colPos] = -1
    if (IsBingo(markedRows) || IsBingo(markedColumns)) {
      return position, FindSolution(board, bingoNumber)
    }
  }
  fmt.Println("-------", markedColumns, markedRows)
  return -1, -1
}

func FindSolution(board [][]int, winningNumber int) int {
  solution := 0
  for _, row := range board {
    for _, number := range row {
      if number == -1 { continue }
      solution += number
    }
  }

  return solution * winningNumber
}

func IsBingo(markedPositions []int) bool {
  for _, posCount := range markedPositions { if posCount == 5 { return true} }
  return false
}

func NumberPositionOnBoard(board [][]int, number int) (rowPos int, colPos int) {
  for rowPos, row := range board {
    for colPos, boardNumber := range row {
      if boardNumber == number {
        return rowPos, colPos
      }
    }
  }
  return -1, -1
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

