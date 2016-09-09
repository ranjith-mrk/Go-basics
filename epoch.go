package main 

import "fmt"
import "time"

func main() {
  now := time.Now()
  fmt.Println(now)

  secs := now.Unix()
  fmt.Println("Seconds is:", secs)

  fmt.Println(time.Unix(secs, 0))
}