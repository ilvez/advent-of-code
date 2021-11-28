package main

import "fmt"
import "crypto/md5"
import "encoding/hex"

func main() {
  input := "yzbqklnj"
  i := 0
  for {
    currentInput := fmt.Sprintf("%s%d", input, i)
    hash := MD5(currentInput)
    if hash[:6] == "000000" {
      fmt.Println(fmt.Sprintf("%d", i))
      break
    }

    i += 1
  }
}

// MD5 hashes using md5 algorithm
func MD5(text string) string {
  hash := md5.Sum([]byte(text))
  return hex.EncodeToString(hash[:])
}
