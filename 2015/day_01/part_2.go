package main

import "fmt"
import "io/ioutil"
import "os"

func main() {
  directions, _ := ioutil.ReadFile("input")
  current_floor := 0
  fmt.Println(fmt.Sprintf("Steps count: %d", len(directions)))

  for i := 0; i < len(directions); i++ {
    if current_floor == -1 {
      fmt.Println(fmt.Sprintf("Basement at step: %d", i))
      os.Exit(0)
    }
    if directions[i] == '(' {
      current_floor += 1
    } else if directions[i] == ')' {
      current_floor -= 1
    }
  }
}
