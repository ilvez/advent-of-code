package main

import "fmt"
import "os"
import "bufio"

func main() {
  inputs := FileToLines("input2")
  counts := make([]int, len(inputs[0]))
  for _, inputString := range inputs {
    for i, char := range inputString {
      if char == '0' { counts[i] -= 1 }
      if char == '1' { counts[i] += 1 }
    }
    fmt.Println(inputString, counts)
  }
  gammaRate := CalculateGamma(counts)
  fmt.Println("Result", gammaRate)
}

func CalculateGamma(counts []int) int {
  gammaSlice := make([]int, len(counts))
  for i, value := range counts {
    if value < 0 { gammaSlice[i] = '0' }
    if value >= 0 { gammaSlice[i] = '1' }
  }
  fmt.Println(string(gammaSlice.Bytes())
  return len(counts)
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

