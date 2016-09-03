package main 

import "fmt"

func main() {
	message := make(chan string)

	go func() {
		message <- "message"
	}()

	msg := <- message
	fmt.Println(msg)
}