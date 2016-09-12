package main 

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
  sign := make(chan os.Signal, 1)
  done := make(chan bool, 1)

  signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

  go func() {
    received := <-sign
    fmt.Println("The received signal:", received)
    done <- true
  }()

  fmt.Println("Waiting for the process to end")
  <-done 
  fmt.Println("Completed")
}