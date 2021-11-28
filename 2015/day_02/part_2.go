package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"
import "sort"

func main() {
  sizes, _ := FileToLines("input")

  var totalLength int64 = 0
  for _, size := range(sizes) {
    l, w, h := ParseLWH(size)
    a, b := TwoSmallestOfThree(l, w, h)
    currentLength := 2 * a + 2 * b + Volume(l, w, h)
    fmt.Println("  inputs:", l, w, h)
    fmt.Println("  TwoSmallestOfThree:", a, b)
    fmt.Println("  currentLength:", currentLength)
    totalLength += int64(currentLength)
    fmt.Println("Result:", totalLength)
  }
}

func TwoSmallestOfThree(l int, w int, h int) (a int, b int) {
  arr := [3] int { l, w, h}

  slice := arr[:]

  sort.Ints(slice)

  return arr[0], arr[1]
}

func Volume(l int, w int, h int) (volume int) {
  return l * w * h
}

func ParseLWH(size string) (l int, w int, h int) {
  split := strings.Split(size, "x")
  return StringToInt(split[0]), StringToInt(split[1]), StringToInt(split[2])
}

func StringToInt(valueString string) (value int) {
  value, _ = strconv.Atoi(valueString)

  return
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
