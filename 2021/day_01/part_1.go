package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
  depths, _ := FileToLines("input")

  previousDepth := 0
  descents := 0

  for _, depthString := range(depths) {
    depth, _ := strconv.Atoi(depthString)
    if previousDepth != 0 && previousDepth < depth {
      descents += 1
    }
    previousDepth = depth
  }
  fmt.Println(descents)
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
