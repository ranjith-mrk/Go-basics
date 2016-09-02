package main 

import "fmt"
import "math"

type shapes interface {
	area() float64
	perimeter() float64
}

type rectange struct {
	length, breadth float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectange) area() float64 {
	return r.length * r.breadth
}

func (r rectange) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func measure(s shapes) {
	fmt.Println(s)
	fmt.Println("Area is:", s.area())
	fmt.Println("Perimeter is:", s.perimeter())
}

func main() {
	r := rectange{4,5}
	c := circle{3}

	measure(r)
	measure(c)
}