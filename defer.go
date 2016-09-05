package main 

import "fmt"

func one(){
  fmt.Println("Function one")
}

func two(){
  fmt.Println("Function two")
}

func three(){
  fmt.Println("Function three")
}

func main() {
  one()
  defer three()
  two()
}