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

func cleanDeadEdges(edges *[]Edge) {
}

func parseEdges(inputFile string) (edges []Edge) {
  lines := aoc.FileToLines(inputFile)
  for _, line := range lines {
    nodes := strings.Split(line, "-")
    a := Node { value: nodes[0], isBig: isBig(nodes[0])}
    b := Node { value: nodes[1], isBig: isBig(nodes[1])}
    edges = append(edges, Edge { a: a, b: b })
  }
  return
}

type Node struct {
  value string
  isBig bool
}

type Edge struct {
  a, b Node
}

func isBig(s string) bool {
  return strings.ToUpper(s) == s
}
