package main

import "fmt"
import "strings"
import aoc "aocgo/aochelper"


func main() {
  inputString := aoc.FileToLines("input")
  state := InputToState(inputString[0])

  for day := 1; day <= 256; day++ {
    state = NewState(state, day)
  }
  fmt.Println("Result", FishCount(state))
}

func FishCount(state map[int]int) (fishCount int) {
  fishCount = 0
  for _, fishAmount := range state {
    fishCount += fishAmount
  }
  return
}

func NewState(state map[int]int, day int) (map[int]int) {
  newState := make(map[int]int, 0)
  newbornAmount := 0
  for fishAge, amount := range state {
    if fishAge == 0 {
      newState[6] += amount
      newbornAmount = amount
    } else {
      newState[fishAge-1] += amount
    }
  }
  if newbornAmount > 0 {
    newState[8] = newbornAmount
  }
  return newState
}

func InputToState(input string) map[int]int {
  stringArray := strings.Split(input, ",")
  state := make(map[int]int, len(stringArray))
  for _, numString := range(stringArray) {
    fishAge := aoc.StringToInt(numString)
    state[fishAge] +=1
  }
  return state
}
