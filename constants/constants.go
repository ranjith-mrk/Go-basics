package main 

import "math"
import "fmt"

const a int = 20000
func  main() {
	const m = 5000000
	fmt.Println("a is", a)
	fmt.Println("m is", m)
	fmt.Println("typecase m is", int64(m))
	fmt.Println("Sin is", math.Sin(m))
}