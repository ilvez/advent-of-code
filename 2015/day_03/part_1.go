package main

import "fmt"
import "io/ioutil"

type Coordinate struct {
  x, y int
}
func main() {
  directions, _ := ioutil.ReadFile("input")

  currentCoord := Coordinate { 0, 0 }

  coordinates := [] Coordinate { currentCoord }
  for _, direction := range(directions) {
    currentCoord = NextCoordinate(direction, currentCoord)
    if !ElementExists(coordinates, currentCoord) {
      coordinates = append(coordinates, currentCoord)
    }
    fmt.Println(string(direction), currentCoord)
  }
  fmt.Println("Result:", len(coordinates))
}

func ElementExists(coordinates []Coordinate, currentCoord Coordinate) bool {
  for _, coordinate := range coordinates {
    if coordinate == currentCoord {
      return true
    }
  }
  return false
}
func NextCoordinate(direction byte, currentCoord Coordinate) (Coordinate) {
  switch direction {
  case '>': return Coordinate { currentCoord.x + 1, currentCoord.y }
  case '<': return Coordinate { currentCoord.x - 1, currentCoord.y }
  case '^': return Coordinate { currentCoord.x, currentCoord.y + 1 }
  case 'v': return Coordinate { currentCoord.x, currentCoord.y - 1 }
  }
  return currentCoord
}
