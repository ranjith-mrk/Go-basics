package main 

import "fmt"

func main() {
	if 3 % 2 == 0 {
		fmt.Println("It is even")
	} else {
		fmt.Println("It's odd")
	}

	if 4 % 2 == 0 {
		fmt.Println("It's even")
	}

	if a := 4; a == 4 {
		fmt.Println("Its 4")
	}
}