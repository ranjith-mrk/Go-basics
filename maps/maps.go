package main

import "fmt"

func main() {
	var m = make(map[string]string)

  m["a"] = "Apple"
  m["b"] = "Ball"

  fmt.Println("The value of m is", m)
  fmt.Println("Specific key value", m["a"])

  fmt.Println("The length i.e no of keys", len(m))

  bval, presence := m["b"]
  fmt.Println("presence check of b", presence)
  fmt.Println("val retried is", bval)

  delete(m, "b")

  fmt.Println("map after delete of b", m)

  n := map[string]int{"hello": 1, "world": 2}

  fmt.Println("Directly initialized map", n)

}