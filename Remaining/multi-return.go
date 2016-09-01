package main 

import "fmt"

func swap(a, b int) (int, int){
  return b, a
}

func main() {
  a,b := swap(5,6)
  fmt.Println(a,b)
}