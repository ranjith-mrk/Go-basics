package main 

import "fmt"
import "time"

func main() {
  ticker := time.NewTicker(time.Second)
  go func () {
    for t := range ticker.C {
      fmt.Println("Printing with ticker at:", t)
    }
  }()

  time.Sleep(time.Second * 5)
  ticker.Stop()
  fmt.Println("Completed ticker")
}