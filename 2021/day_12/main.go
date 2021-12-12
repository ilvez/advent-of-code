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

func cleanDeadEdges(es []Edge) (cleanedEdges []Edge) {
  for _, e := range es {
    if !e.isDeadEnd(es) {
      cleanedEdges = append(cleanedEdges, e)
    }
  }
  return
}

func (edge *Edge) isDeadEnd(es []Edge) bool {
  fmt.Println("E", edge.toString())
  if edge.b.isBig || edge.a.isBig {
    return false
  }
  for _, e := range es {
    if e.a == edge.b {
      return false
    }
  }
  return true
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
