package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
  depths, _ := FileToLines("input")

  var currentWindow [3] int
  var previousWindow [3] int
  descents := 0

  for index, depthString := range(depths) {
    depth, _ := strconv.Atoi(depthString)

    if index > 0 {
      previousWindow[(index - 1) % 3] = currentWindow[(index - 1) % 3]
    }
    currentWindow[index % 3] = depth

    if index >= 3 && sum(previousWindow) < sum(currentWindow) {
      descents += 1
    }
  }
  fmt.Println("Result", descents)
}

func sum(array [3]int) int {
  result := 0
  for _, v := range array {
    result += v
  }
  return result
}

// Copypasted from internets
func FileToLines(filePath string) (lines []string, err error) {
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
