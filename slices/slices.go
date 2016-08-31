package main

import "fmt"

func main() {
	
	var a = make([] int, 3)
	a[0] = 0
	a[1] = 1
	a[2] = 2

	fmt.Println("The slices value is", a)

	fmt.Println("The length is", len(a))

	a = append(a, 3, 4, 5)

	fmt.Println("The new slices is", a)
	fmt.Println("The new slices length is", len(a))

	b := [3]int{6,7,8}

	fmt.Println("b is", b)

	c := a[1:4]
	fmt.Println("1:4 is", c)
	fmt.Println("1:", a[1:])
	fmt.Println(":4", a[:4])
}