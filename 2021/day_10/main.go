package main

import(
  aoc "aocgo/aochelper"
  "errors"
  "fmt"
  "os"
)
func main() {
  printPart1Solution()
  printPart2Solution()
}

func printPart1Solution() {
  score := 0
  for _, instructions := range aoc.FileToLines("input") {
    _, err := parseChunks(instructions)
    if err != nil {
      i := illegalToPoints(err)
      score += i
    }
  }
  fmt.Println("Part1 solution:", score)
}

func printPart2Solution() {
  for _, instructions := range aoc.FileToLines("input2") {
    incomplete, err := parseChunks(instructions)
    if err != nil {
      continue
    }
    fmt.Println(incomplete.toString())
    fmt.Println(completeChunkScore(incomplete))
  }
}

func completeChunkScore(chunk *Chunk) (points int) {
  currentChunk := chunk
  for (*currentChunk).parent != nil {
    points = points * 5 + completePoints((*currentChunk).closingCommand())
    currentChunk = (*currentChunk).parent
  }
  points = points * 5 + completePoints((*currentChunk).closingCommand())
  return
}

func parseChunks(instructions string) (*Chunk, error) {
  chunk := Chunk {}
  currentChunk := &chunk
  for _, i := range instructions {
    command := string(i)
    if (*currentChunk).opening == "" {
      (*currentChunk).opening = command
      continue
    }
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

func newChunk(opening string) Chunk {
  return Chunk { opening: opening }
}

func (c *Chunk) newSubChunk(opening string) (*Chunk) {
  if !c.isOpen() {
    fmt.Println("Trying to add subchunk to closed chunk", c.toString(), string(opening))
    os.Exit(1)
  }
  subChunk := newChunk(opening)
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
    colorYellow = "\033[33m"
    colorBlue = "\033[34m"
    colorPurple = "\033[35m"
    colorCyan = "\033[36m"
    colorWhite = "\033[37m"
    bold = "\033[38m"
)
