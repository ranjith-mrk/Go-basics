package main 

import "fmt"

func main() {
	message := make(chan string, 2)
	message <- "one"
	message <- "two"

	fmt.Println(<-message)
	fmt.Println(<-message)
}