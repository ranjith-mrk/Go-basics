package main 

import "fmt"

func addFive() func(int) int {
	no := 5
	return func(a int) int{		
		return a + no
	}
	
}

func  main() {
	plusFive := addFive()
	fmt.Println("5 added is:",plusFive(5))
	fmt.Println("10 added is:",plusFive(10))
	fmt.Println("15 added is:",plusFive(15))
	fmt.Println("22 added is:",plusFive(22))
}