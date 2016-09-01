package main 

import "fmt"

func changeval(i int){
	i = 0
}

func changepointer(j *int) {
	*j = 0
}

func main() {
	i := 100
	changeval(i)
	fmt.Println("Calling after changeval", i)
	changepointer(&i)
	fmt.Println("Calling after changepointer", i)
}