package main 

import "fmt"
import "time"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func () {
		time.Sleep(time.Second)
		fmt.Println("Testing for value of time's second", time.Second)
		ch1 <- "one"
	}()

	go func () {
		time.Sleep(time.Second * 2)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
			case msg := <- ch1:
				fmt.Println("Ch1 Message is", msg)
			case msg := <- ch2:
				fmt.Println("Ch2 Message is", msg)
		}
	}

}