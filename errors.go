package main 

import "fmt"

type argError struct {
  arg  int
  prob string
}

func test(i int) (int, error) {
	return i, &argError{i, "Will always throw error"}
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func main() {
	_, e := test(1)
	if ae, ok := e.(*argError); ok {
		fmt.Println("ok is:", ok)
		fmt.Println(ae)
    fmt.Println(ae.arg)
    fmt.Println(ae.prob)
  }
}