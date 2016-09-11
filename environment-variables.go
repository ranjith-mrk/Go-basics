package main 

import(
  "fmt"
  "os"
  "strings"
)

func main() {
  os.Setenv("HELLO", "1")

  fmt.Println("HELLO:", os.Getenv("HELLO"))
  fmt.Println("WORLD:", os.Getenv("WORLD"))

  fmt.Println()
  for _, e := range os.Environ() {
      pair := strings.Split(e, "=")
      fmt.Println(pair[0])
  }
}