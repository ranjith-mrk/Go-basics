package main 

import "fmt"

func main() {
  two   := addTwo(1,1)
  three := addThree(1,1,1)

  fmt.Println("sum two is", two)
  fmt.Println("sum two is", three)
}

func addTwo(n1 int, n2 int) int{
  return n1 + n2
}

func  addThree(n1, n2, n3 int) int{
  return n1 + n2 + n3
}
