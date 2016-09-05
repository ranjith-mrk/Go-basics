package main 

import "fmt"
import "time"

func main() {
  messages := make(chan int, 5)
  for i := 1; i <=5; i = i + 1 {
    messages <- i
  }

  close(messages)

  limiter := time.Tick(time.Millisecond * 300)

  for m :=  range messages {
    <- limiter
    fmt.Println("Message:", m, time.Now())
  }
}