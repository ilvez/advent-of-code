package main

import "fmt"
import "math"
import aoc "aocgo/aochelper"


func main() {
  inputString := aoc.FileToLines("input")
  crabLocations := aoc.StringToInt64Array(inputString[0])
  fmt.Println("Part 1 result:", FindCheapestFuel(crabLocations, FindFuelPart1))
  fmt.Println("Part 2 result:", FindCheapestFuel(crabLocations, FindFuelPart2))
}

func FindCheapestFuel(crabLocations []int64, fuelFunc func([]int64, int64) int64) int64 {
  var cheapestFuel int64 = math.MaxInt64 - 1
  minLoc, maxLoc := findMinAndMax(crabLocations)
  for i := minLoc; i <= maxLoc; i++ {
    fuel := fuelFunc(crabLocations, i)
    // fmt.Println(i, fuel)
    if fuel < cheapestFuel {
      cheapestFuel = fuel
    }
  }
  return cheapestFuel
}

func FindFuelPart1(crabLocations []int64, goToLocation int64) (fuelAmount int64) {
  fuelAmount = 0
  for _, location := range crabLocations {
    fuelAmount += aoc.AbsInt64(goToLocation - location)
  }
  return
}

func FindFuelPart2(crabLocations []int64, goToLocation int64) (fuelAmount int64) {
  fuelAmount = 0
  for _, location := range crabLocations {
    distance := aoc.AbsInt64(goToLocation - location)
    fuelAmount += (distance * (distance + 1)) / 2
  }
  return
}

func findMinAndMax(a []int64) (min, max int64) {
  min = a[0]
  max = a[0]
  for _, value := range a {
    if value < min {
      min = value
    }
    if value > max {
      max = value
    }
  }
  return
}

