package main 

import "fmt"

func main() {
	
	mess1 := make(chan string)

  select {
    case msg := <- mess1:
      fmt.Println("The message1 is ", msg)
    default:
      fmt.Println("Default block1 of message 1")
  }

  msg := "hi"
  select {
    case mess1 <- msg :
      fmt.Println("The message1 is ", mess1)
    default:
      fmt.Println("Default block2 of message 1")
  }



}
