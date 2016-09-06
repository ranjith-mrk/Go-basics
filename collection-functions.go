package main 

import (
  "strings"
  "fmt"
)


func Index(list []string, t string) int {
  for i, v := range list {
    if v == t {
      return i
    }
  }
  return -1
}

func Include(list []string, t string) bool {
  return Index(list,t) >= 0
}

func Any(list []string, f func(string) bool) bool{
  for _, v := range list {
    if f(v){
      return true
    }
  }
  return false
}

func All(list []string, f func(string) bool) bool{
  for _, v := range list {
    if !f(v){
      return false
    }
  }
  return true
}

func main() {
  names := []string{"Ranjith", "Kumar", "Goku", "Natsu", "Kamina"}
  fmt.Println(Index(names, "Goku"))
  fmt.Println(Include(names, "Kumar"))
  fmt.Println(Any(names, func(name string) bool{
    return strings.HasPrefix(name, "Na")
  }))
  fmt.Println(All(names, func(name string) bool{
    return len(name) > 5
  }))
}