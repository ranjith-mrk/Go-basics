package main 

import "crypto/sha1"
import "fmt"

func main() {
  str := "Sample String"

  h := sha1.New()
  h.Write([]byte(str))

  sum := h.Sum(nil)

  fmt.Println(str)
  fmt.Println(sum)
  fmt.Printf("%x\n",sum)
}