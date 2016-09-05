package main 

import "fmt"

func main() {
  mess := make(chan int)
  done := make(chan bool)

  go func () {
    for{
      m, more := <- mess
      if more {
        fmt.Println("Message received:",m)
      }else{
        fmt.Println("Mess completed")
        done <- true
        return
      }
    }
  }()

  for i := 0; i < 10; i = i + 1 {
    mess <- i 
    fmt.Println("Sending message:",i)    
  }
  close(mess)
  <- done
}