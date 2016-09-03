package main 

import "fmt"

func l(from string) {
	for i := 0; i < 3; i = i + 1 {
		fmt.Println("from", from, ":",i)
	}
}

func main() {
	l("normal")
	go l("routine")

	go func (s string) {
		fmt.Println("Inline function string",s)
	}("inline routine")

	var input string
  fmt.Scanln(&input)

  fmt.Println("The input is:", input)

}