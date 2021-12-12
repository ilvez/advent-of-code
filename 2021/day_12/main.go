package main

import(
  aoc "aocgo/aochelper"
  "fmt"
  "strings"
)
func main() {
  inputFile := "input2"
  fmt.Println(parseEdges(inputFile))
}

func parseEdges(inputFile string) (edges []Edge) {
  lines := aoc.FileToLines(inputFile)
  for _, line := range lines {
    nodes := strings.Split(line, "-")
    edges = append(edges, Edge { a: nodes[0], b: nodes[1] })
  }
  return
}

type Edge struct {
  a, b string
}
