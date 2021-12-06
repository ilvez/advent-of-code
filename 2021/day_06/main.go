package main

import "fmt"
import "strings"
import aoc "aocgo/aochelper"


func main() {
  inputString := aoc.FileToLines("input")
  state := InputToState(inputString[0])

  fmt.Println("First state", state)
  for i := 0; i < 80; i++ {
    state = NewState(state)
  }
  fmt.Println("Result", len(state))
}

func NewState(state []int) ([]int) {
  newFishes := make([]int, 0)
  for i, fish := range(state) {
    if fish == 0 {
      state[i] = 6
      newFishes = append(newFishes, 8)
    } else {
      state[i] -= 1
    }
  }
  newState := make([]int, 0)
  newState = append(append([]int{}, state...), newFishes...)
  return newState
}

func InputToState(input string) []int {
  stringArray := strings.Split(input, ",")
  state := make([]int, len(stringArray))
  for i, numString := range(stringArray) {
    state[i] = aoc.StringToInt(numString)
  }
  return state
}
