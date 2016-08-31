package main 

import "fmt"

func main() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 5; j = j + 1 {
		fmt.Println(j)
	}

	for {
		fmt.Println("Hello world")
		break
	}
}