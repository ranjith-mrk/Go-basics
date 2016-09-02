package main 

import "fmt"

type rectange struct {
	length, breadth int
}

func (r rectange) area()  int {
	return r.length * r.breadth
}

func (r rectange) perimeter() int {
	return 2*(r.length + r.breadth)
}

func main() {
	r1 := rectange{10, 20}

	fmt.Println("Area of r1", r1.area())
	fmt.Println("Perimeter of r1", r1.perimeter())
}