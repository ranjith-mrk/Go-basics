package main 

import "fmt"
import "os"

func main() {
  defer fmt.Println("Test")
  os.Exit(3)
}