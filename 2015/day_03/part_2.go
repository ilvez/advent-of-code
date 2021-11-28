package main

import "fmt"
import "io/ioutil"

type Coordinate struct {
  x, y int
}
func main() {
  directions, _ := ioutil.ReadFile("input")

  santaCoord := Coordinate { 0, 0 }
  roboCoord := Coordinate { 0, 0 }
  currentCoord := Coordinate { 0, 0 }

  allHouses := [] Coordinate { currentCoord }

  for index, direction := range(directions) {
    if index % 2 == 0 {
      santaCoord = NextCoordinate(direction, santaCoord)
      currentCoord = santaCoord
    } else {
      roboCoord = NextCoordinate(direction, roboCoord)
      currentCoord = roboCoord
    }
    if !ElementExists(allHouses, currentCoord) {
      allHouses = append(allHouses, currentCoord)
    }
  }
  fmt.Println("Result:", len(allHouses))
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
