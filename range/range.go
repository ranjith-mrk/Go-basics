package main 

import "fmt"

func main() {
	a := [7]int{1,2,3,4,5,6,7}
	sum := 0
	for i, num := range a {
		fmt.Println("Index:",i, ", Number:", num)
		sum = sum + num
	}

	fmt.Println("The sum is:", sum)

	m := map[string]int{"A": 65, "B": 66, "C": 67}
	for key, value := range m{
		fmt.Println("key:",key, ", Number:", value)	
	}

}