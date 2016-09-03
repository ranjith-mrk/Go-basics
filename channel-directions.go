package main 

import "fmt"

func ping(ping chan<- string, message string) {
	fmt.Println("Inside ping")
	ping <- message
}

func pong(ping <-chan string, pong chan<- string) {
	fmt.Println("Inside pong")
	message := <- ping
	pong <- message
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Hello world")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}