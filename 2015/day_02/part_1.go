package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

func main() {
  sizes, _ := FileToLines("input")

  totalArea := 0
  for _, size := range(sizes) {
    l, w, h := ParseLWH(size)
    totalArea += NeededPaperArea(l, w, h)
  }
  fmt.Println("total square feet of wrapping paper:", totalArea)
}

func NeededPaperArea(l int, w int, h int) (area int) {
  side1, side2, side3 := l*w, w*h, h*l
  return 2 * side1 + 2 * side2 + 2 * side3 + SmallestSideArea(side1, side2, side3)
}

func SmallestSideArea(l int, w int, h int) (smallest int) {
  if (l <= w && l <= h) {
    return l
  } else if (w <= l && w <= h) {
    return w
  } else {
    return h
  }
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
