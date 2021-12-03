package main

import "fmt"
import "os"
import "bufio"
import "strconv"

func main() {
  inputs := FileToLines("input")
  fmt.Println(
    "Result",
    BinaryStringToInteger(OxygenGeneratorRatingBinaryString(inputs)) *
      BinaryStringToInteger(CarbonDioxideScrubberBinaryString(inputs)))
}

func BinaryStringToInteger(input string) int {
  result, _ := strconv.ParseInt(input, 2, 32)
  return int(result)
}

func OxygenGeneratorRatingBinaryString(inputs []string) string {
  if len(inputs) == 1 {
    return ""
  } else {
    commonValue, survivors := MostCommon(inputs)
    if len(survivors) == 1 {
      return commonValue + survivors[0]
    }

    return commonValue + OxygenGeneratorRatingBinaryString(survivors)
  }
}

func CarbonDioxideScrubberBinaryString(inputs []string) string {
  if len(inputs) == 1 {
    return ""
  } else {
    commonValue, survivors := LeastCommon(inputs)
    if len(survivors) == 1 {
      return commonValue + survivors[0]
    }

    return commonValue + CarbonDioxideScrubberBinaryString(survivors)
  }
}

func LeastCommon(inputs []string) (commonValue string, survivors []string) {
  zeroStarting := NumberStartingStrings(inputs, "0")
  oneStarting := NumberStartingStrings(inputs, "1")

  if len(zeroStarting) <= len(oneStarting) {
    return "0", zeroStarting
  } else {
    return "1", oneStarting
  }
}
func MostCommon(inputs []string) (commonValue string, survivors []string) {
  zeroStarting := NumberStartingStrings(inputs, "0")
  oneStarting := NumberStartingStrings(inputs, "1")

  if len(zeroStarting) > len(oneStarting) {
    return "0", zeroStarting
  } else {
    return "1", oneStarting
  }
}

func NumberStartingStrings(inputs []string, number string) (outputs []string) {
  outputs = make([]string, 0)

  for _, inputString := range inputs {
    if string(inputString[0]) == number {
      outputs = append(outputs, inputString[1:])
    }
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

