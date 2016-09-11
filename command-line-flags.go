package main

import "fmt"
import "flag"

func main() {
  word    := flag.String("word", "hello", "Enter string")
  number  := flag.Int("number", 10, "Enter number")
  truth   := flag.Bool("truth", false, "Enter bool")

  fmt.Println("word:", *word)
  fmt.Println("number:", *number)
  fmt.Println("truth:", *truth)
  fmt.Println("other args:", flag.Args())
}