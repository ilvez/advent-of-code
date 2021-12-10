package main

import(
  aoc "aocgo/aochelper"
  "errors"
  "fmt"
  "os"
)
func main() {
  printPart1Solution()
}

func printPart1Solution() {
  score := 0
  for _, instructions := range aoc.FileToLines("input") {
    _, illegal := parseChunks(instructions)
    if illegal != "" {
      i := illegalToPoints(illegal)
      score += i
    }
  }
  fmt.Println("Part1 solution:", score)
}

func parseChunks(instructions string) (Chunk, string) {
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
        return *currentChunk, command
      }
      if currentChunk.parent != nil {
        currentChunk = currentChunk.parent
      } else {
        return *currentChunk, ""
      }
    }
  }
  return *currentChunk, ""
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

func (c *Chunk) isOpen() bool {
  return c.closing == ""
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
func (c *Chunk) toString() string {
  chunks := ""
  for _, chunk := range c.chunks {
    chunks += chunk.toString()
  }
  parent := (*c.parent).opening
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
    msg := fmt.Sprintf("Trying to close", c.toString(), "with", closing)
    return errors.New(msg)
  }
  return nil
}


func isOpening(i string) bool {
  return i == "{" || i == "(" || i == "[" || i == "<"
}

func isClosing(i string) bool {
  return i == "}" || i == ")" || i == "]" || i == ">"
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

func printCommands(stack aoc.Stack) {
  fmt.Println(stackAsString(stack))
}

func stackAsString(stack aoc.Stack) (result string) {
  for _, c := range stack.Data() {
    result += string(c.(rune))
  }
  return
}

func illegalToPoints(illegal string) int {
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
