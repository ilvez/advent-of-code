package main

import(
  aoc "aocgo/aochelper"
  "errors"
  "fmt"
  "os"
  "sort"
)
func main() {
  inputFile := "input"
  printPart1Solution(inputFile)
  printPart2Solution(inputFile)
}

func printPart1Solution(inputFile string) {
  score := 0
  for _, instructions := range aoc.FileToLines(inputFile) {
    _, err := parseChunks(instructions)
    if err != nil {
      i := illegalToPoints(err)
      score += i
    }
  }
  fmt.Println("Part 1 solution:", score)
}

func printPart2Solution(inputFile string) {
  lines := aoc.FileToLines(inputFile)
  scores := make([]int, 0)
  for _, instructions := range lines {
    incomplete, err := parseChunks(instructions)
    if err != nil {
      continue
    }
    scores = append(scores, completeChunkScore(incomplete))
  }
  sort.Ints(scores)
  fmt.Println("Part 2 solution:", scores[len(scores)/2])
}

func completeChunkScore(chunk *Chunk) (points int) {
  currentChunk := chunk
  for (*currentChunk).parent != nil {
    points = points * 5 + completePoints((*currentChunk).closingCommand())
    currentChunk = (*currentChunk).parent
  }
  if (*currentChunk).closingCommand() != "" {
    points = points * 5 + completePoints((*currentChunk).closingCommand())
  }
  return
}

func parseChunks(instructions string) (*Chunk, error) {
  chunk := Chunk {}
  currentChunk := &chunk
  for _, i := range instructions {
    command := string(i)
    if isOpening(command) {
      currentChunk = (*currentChunk).newSubChunk(command)
    } else {
      err := currentChunk.closeChunk(command)
      if err != nil {
        return currentChunk, err
      }
      if currentChunk.parent != nil {
        currentChunk = currentChunk.parent
      } else {
        return currentChunk, nil
      }
    }
  }
  return currentChunk, nil
}

type Chunk struct {
  opening, closing string
  parent *Chunk
  chunks []Chunk
}

func (c *Chunk) newSubChunk(opening string) (*Chunk) {
  if !c.isOpen() {
    fmt.Println("Trying to add subchunk to closed chunk", c.toString(), string(opening))
    os.Exit(1)
  }
  subChunk := Chunk { opening: opening }
  subChunk.parent = c
  c.chunks = append(c.chunks, subChunk)
  return &subChunk
}

func (c *Chunk) closingCommand() (match string) {
  switch c.opening {
    case "{": match = "}"
    case "(": match = ")"
    case "[": match = "]"
    case "<": match = ">"
  }
  return
}

func (c *Chunk) isOpen() bool {
  return c.closing == ""
}

func (c *Chunk) toString() string {
  chunks := ""
  for _, chunk := range c.chunks {
    chunks += chunk.toString()
  }
  parent := "R"
  if c.parent != nil {
    parent = (*c.parent).opening
  }

  return fmt.Sprintf(
    "<%s%s%s%s [%s] %s%s%s>",
    colorGreen,
    c.opening,
    c.closing,
    colorReset,
    chunks,
    colorRed,
    parent,
    colorReset,
  )
}

func (c *Chunk) printChunk() {
  fmt.Println(c.toString())
}

func (c *Chunk) closeChunk(closing string) error {
  if matchingOpening(closing) == c.opening {
    c.closing = closing
  } else {
    return errors.New(closing)
  }
  return nil
}

func isOpening(i string) bool {
  return i == "{" || i == "(" || i == "[" || i == "<"
}

func matchingOpening(i string) (match string) {
  switch i {
    case "}": match = "{"
    case ")": match = "("
    case "]": match = "["
    case ">": match = "<"
    default: os.Exit(1)
  }

  return
}

func completePoints(i string) int {
  if i == ")" {
    return 1
  } else if i == "]" {
    return 2
  } else if i == "}" {
    return 3
  } else if i == ">" {
    return 4
  }
  return 0
}

func illegalToPoints(err error) int {
  illegal := fmt.Sprint(err)
  if illegal == ")" {
    return 3
  } else if illegal == "]" {
    return 57
  } else if illegal == "}" {
    return 1197
  } else if illegal == ">" {
    return 25137
  }
  return 0
}

const (
    colorReset = "\033[0m"
    colorRed = "\033[31m"
    colorGreen = "\033[32m"
)
