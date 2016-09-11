package main 

import "fmt"
import "os"

func main() {
  allargs := os.Args
  args := os.Args[1:]
  fmt.Println("All args:", allargs)
  fmt.Println("Actual args:", args)
}