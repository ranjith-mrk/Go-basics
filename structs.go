package main 

import "fmt"

type employee struct {
	name string
	empid int
}

func  main() {
	fmt.Println(employee{"ranjith", 1})
	fmt.Println(employee{empid: 2, name: "Kumar"})
	emp3 := employee{"Dass", 3}
	fmt.Println("Emp3 name is", emp3.name)
	fmt.Println("Emp3 id is", emp3.empid)
}