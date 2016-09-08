package main 

import (
  "fmt"
  "regexp"
  "bytes"
)

func main() {
  r, _ := regexp.Compile("p([a-z]+)ch")
  fmt.Println(r.MatchString("peach"))
  fmt.Println(r.MatchString("Peach"))

  fmt.Println(r.FindString("peach beach"))
  fmt.Println(r.FindStringIndex("peach beach"))
  fmt.Println(r.FindAllString("peach beach pinch", 1))
  fmt.Println(r.FindAllString("peach beach pinch", -1))

  fmt.Println(r.FindAllString("peach beach pinch peach peach", 2))

  in := []byte("a peach")
  out := r.ReplaceAllFunc(in, bytes.ToUpper)
  fmt.Println(string(out))
}