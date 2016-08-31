package main 

import "fmt"

func  main() {
	var a [5]int

	fmt.Println("Empty array is", a)

	a[2] = 100
	fmt.Println("Array with value is:", a[2])

	fmt.Println("The array elements are", a)

	b := [5]int{1,2,3,4,5}

	fmt.Println("The length of b is", len(b))
	fmt.Println("The elements of b are", b)

	var c [3][3] int
	fmt.Println("Two d array is", c)
}