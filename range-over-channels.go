package main 

import "fmt"

func main() {
  mess := make(chan string, 2)
  mess <- "mess 1"
  mess <- "mess 2"

  close(mess)

  for m := range mess {
    fmt.Println(m)
  }
}