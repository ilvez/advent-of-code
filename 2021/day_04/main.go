package main

import(
  aoc "aocgo/aochelper"
  "fmt"
  "strconv"
  "regexp"
)

func main() {
  inputs := aoc.FileToLines("input")
  printPart1Solution(inputs)
  printPart2Solution(inputs)
}

func printPart1Solution(inputs []string) {
  winningResult := 0
  bestBoardPosition := 99999

  bingoNumbers := parseIntArray(inputs[0])
  boards := parseBoards(inputs[2:])
  for _, board := range(boards) {
    position, result := SolveBoard(board, bingoNumbers)
    if position < bestBoardPosition {
      bestBoardPosition = position
      winningResult = result
    }
  }

  fmt.Println("Part 1 result:", winningResult)
}

func printPart2Solution(inputs []string) {
  bingoNumbers := parseIntArray(inputs[0])
  bestBoardPosition := 0
  winningResult := 0

  boards := parseBoards(inputs[2:])
  for _, board := range(boards) {
    position, result := SolveBoard(board, bingoNumbers)
    if position > bestBoardPosition {
      bestBoardPosition = position
      winningResult = result
    }
  }

  fmt.Println("Part 2 result:", winningResult)
}

func printBoard(board Board) {
  fmt.Println("----------------")
  for _, row := range board {
    for _, number := range row {
      fmt.Print(number, "|")
    }
    fmt.Println()
  }
  fmt.Println("----------------")
}

func SolveBoard(board Board, bingoNumbers []int) (winningPosition int, solution int) {
  boardSize := len(board)
  markedColumns := make([]int, boardSize)
  markedRows := make([]int, boardSize)
  for position, bingoNumber := range bingoNumbers {
    rowPos, colPos := NumberPositionOnBoard(board, bingoNumber)
    if rowPos == -1 || colPos == -1 { continue }
    markedRows[rowPos] += 1
    markedColumns[colPos] += 1
    board[rowPos][colPos] = -1
    if (IsBingo(markedRows, boardSize) || IsBingo(markedColumns, boardSize)) {
      return position, FindSolution(board, bingoNumber)
    }
  }
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

func IsBingo(markedPositions []int, boardSize int) bool {
  for _, posCount := range markedPositions { if posCount == boardSize { return true} }
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

type Board [][]int

func parseBoards(inputs []string) []Board {
  boards := make([]Board, 0)
  var board Board
  for _, line := range inputs {
    bingoNumbers := parseIntArray(line)
    if len(bingoNumbers) == 0 {
      boards = append(boards, board)
      board = make(Board, 0)
      continue
    }
    board = append(board, bingoNumbers)
  }
  boards = append(boards, board)
  return boards
}

func parseIntArray(line string) []int {
  outputs := make([]int, 0)
  re := regexp.MustCompile(`\d+`)
  for _, numberString := range re.FindAllString(line, -1) {
    number, _ := strconv.Atoi(numberString)
    outputs = append(outputs, number)
  }
  return outputs
}
