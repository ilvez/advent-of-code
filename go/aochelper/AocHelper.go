package aochelper

import "os"
import "bufio"
import "strconv"
import "regexp"

func Abs(x int) int {
  if x < 0 { return -x }
  return x
}
func StringToInt(input string) int {
  a, _ := strconv.Atoi(input)
  return a
}

func StringToIntArray(line string) []int {
  outputs := make([]int, 0)
  re := regexp.MustCompile(`\d+`)
  for _, numberString := range re.FindAllString(line, -1) {
    outputs = append(outputs, StringToInt(numberString))
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
