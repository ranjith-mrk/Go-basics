package main 

import "fmt"

func multipleSum(nums ...int) int{
	fmt.Println("The numbers are", nums)
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}	
	return sum
}

func main() {
	two 	:= multipleSum(1,2)
	fmt.Println("2 args sum is:", two)
	five := multipleSum(1,2,3,4,5)
	fmt.Println("5 args sum is:", five)
}