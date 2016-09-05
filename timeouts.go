package main 

import "fmt"
import "time"

func main() {
	c1 := make(chan string)
	go func (){
		time.Sleep(time.Second * 2)
		c1 <- "channel 1"
	}()

	select {
		case <- c1:
			fmt.Println("Message for ch1")
		case <- time.After(time.Second):
			fmt.Println("Timeout for message c1")
	}

	c2 := make(chan string)
	go func (){
		time.Sleep(time.Second * 2)
		c2 <- "channel 2"
	}()

	select {
		case <- c2:
			fmt.Println("Message for ch2")
		case <- time.After(time.Second * 3):
			fmt.Println("Timeout for message c2")
	}


}