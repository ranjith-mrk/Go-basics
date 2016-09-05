package main 

import (
  "fmt"
  "sort"
)

type ByLength []string

func (s ByLength) Len() int {
  return len(s)
}

func (s ByLength) Swap(i,j int) {
  s[j], s[i] = s[i], s[j]
}

func (s ByLength) Less(i,j int) bool{
  return (len(s[i]) < len(s[j]) )
}

func main() {
  name := []string{"ranjith", "m", "kumar"}
  sort.Sort(ByLength(name))
  fmt.Println("The name is:", name)
}