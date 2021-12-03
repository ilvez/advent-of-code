package main

import "fmt"
import "os"
import "bufio"
import "strconv"

func main() {
  inputs := FileToLines("input")
  counts := CountBites(inputs)
  fmt.Println("Result", CalculateGamma(counts) * CalculateEpsilon(counts))
}

func CalculateEpsilon(counts []int) int {
  slice := ""
  for _, value := range counts {
    if value < 0 { slice += "1" }
    if value >= 0 { slice += "0" }
  }
  result, _ := strconv.ParseInt(slice, 2, 32)
  return int(result)
}

func CalculateGamma(counts []int) int {
  slice := ""
  for _, value := range counts {
    if value < 0 { slice += "0" }
    if value >= 0 { slice += "1" }
  }
  result, _ := strconv.ParseInt(slice, 2, 32)
  return int(result)
}

func CountBites(inputs []string) (counts []int) {
  counts = make([]int, len(inputs[0]))
  for _, inputString := range inputs {
    for i, char := range inputString {
      if char == '0' { counts[i] -= 1 }
      if char == '1' { counts[i] += 1 }
    }
    fmt.Println(inputString, counts)
  }

  return
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

