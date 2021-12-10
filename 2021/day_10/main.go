package main

import(
  aoc "aocgo/aochelper"
  "fmt"
  "os"
)
func main() {
  chunk, _ := parseChunks("<{([([[(<>()){}]>(<<{{")
  chunk.printChunk()
  // c := newChunk("(")
  // c.printChunk()
  // c.newSubChunk("[")
  // c.closeChunk(")")
  // c.printChunk()
  //printPart1Solution()
}

func parseChunks(instructions string) (chunk Chunk, illegal rune) {
  for _, i := range instructions {
    command := string(i)
    if chunk.opening == "" {
      chunk.opening = command
      continue
    }
    if isOpening(command) {
      chunk.newSubChunk(command)
    }
  }
  return chunk, illegal
}

type Chunk struct {
  opening, closing string
  chunks []Chunk
}

func newChunk(opening string) Chunk {
  return Chunk { opening: "(" }
}

func (c *Chunk) newSubChunk(opening string) (*Chunk) {
  if !c.isOpen() {
    fmt.Println("Trying to add subchunk to closed chunk", c.toString(), string(opening))
    os.Exit(1)
  }
  subChunk := newChunk(opening)
  c.chunks = append(c.chunks, subChunk)
  return &subChunk
}

func (c *Chunk) isOpen() bool {
  return c.closing == ""
}

func (c *Chunk) toString() string {
  chunks := ""
  for _, chunk := range c.chunks {
    chunks += chunk.toString()
  }
  return fmt.Sprintf("{'%s%s', [%s], %t}", c.opening, c.closing, chunks, c.isOpen())
}

func (c *Chunk) printChunk() {
  fmt.Println(c.toString())
}

func (c *Chunk) closeChunk(closing string) {
  if matchingOpening(closing) == c.opening {
    c.closing = closing
  } else {
    fmt.Println("Trying to close", c, "with", string(closing))
    os.Exit(1)
  }
}

func printPart1Solution() {
  score := 0
  for _, instructions := range aoc.FileToLines("input2") {
    //fmt.Println(instructions)
    _, illegal := parseChunks(instructions)
    if illegal > 0 {
      i := illegalToPoints(illegal)
      //fmt.Println(string(i))
      score += i
    }
  }
  fmt.Println(score)
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

func illegalToPoints(illegal rune) int {
  if illegal == ')' {
    return 3
  } else if illegal == ']' {
    return 57
  } else if illegal == '}' {
    return 1197
  } else if illegal == '>' {
    return 25137
  }
  return 0
}
