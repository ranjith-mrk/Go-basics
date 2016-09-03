package main 

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Println("Worker started")
	time.Sleep(time.Second)
	fmt.Println("Completed Worker")
	done <- true
}


func main() {
	fmt.Println("Inside main")
	done := make(chan bool)
	go worker(done)
	<- done
	// fmt.Println(done)
	fmt.Println("Completed main")
}