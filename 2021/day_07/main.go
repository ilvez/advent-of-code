package main

import "fmt"
import "math"
import aoc "aocgo/aochelper"
import "io/ioutil"


func main() {
  data, err := ioutil.ReadFile("7-1000000-2.in")
  if err != nil { fmt.Println(err) }
  locations := aoc.StringToInt64Array(string(data))
  fmt.Println("Start calculations")
  fmt.Println("Part 1 result:", FindCheapestFuel(locations, FindFuelPart1))
  fmt.Println("Part 2 result:", FindCheapestFuel(locations, FindFuelPart2))
}

func FindCheapestFuel(locations []int64, fuelFunc func([]int64, int64) int64) int64 {
  var cheapestFuel int64 = math.MaxInt64 - 1
  min, max := FindMinMaxLocations(locations)
  for min != max && min + 1 != max {
    cen := (min + max) / 2
    minVal, maxVal, cenVal := CalculateForBinarySearch(locations, min, max, cen, fuelFunc)
    if minVal < cenVal && maxVal > cenVal {
      max = cen
      cheapestFuel = minVal
    } else if minVal > cenVal && maxVal < cenVal {
      min = cen
      cheapestFuel = maxVal
    } else if minVal > maxVal {
      cheapestFuel = maxVal
      min = cen
    } else if minVal < maxVal {
      cheapestFuel = minVal
      max = cen
    } else {
      cheapestFuel = cenVal
      min = cen
    }

    if cheapestFuel > cenVal { cheapestFuel = cenVal }
  }
  return cheapestFuel
}

func CalculateForBinarySearch(locations []int64, min, max, cen int64, fuelFunc func([]int64, int64) int64) (int64, int64, int64) {
  return fuelFunc(locations, min), fuelFunc(locations, max), fuelFunc(locations, cen)
}

func FindFuelPart1(locations []int64, goToLocation int64) (fuelAmount int64) {
  fuelAmount = 0
  for _, location := range locations {
    fuelAmount += aoc.AbsInt64(goToLocation - location)
  }
  return
}

func FindFuelPart2(locations []int64, goToLocation int64) (fuelAmount int64) {
  fuelAmount = 0
  for _, location := range locations {
    distance := aoc.AbsInt64(goToLocation - location)
    fuelAmount += (distance * (distance + 1)) / 2
  }
  return
}

func FindMinMaxLocations(a []int64) (min, max int64) {
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

