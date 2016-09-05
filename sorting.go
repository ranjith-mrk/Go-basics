package main 

import (
  "fmt"
  "sort"
)

func main() {
  s := []string{"c", "b", "a"}
  i := []int{400,20,30}

  sort.Strings(s)
  sort.Ints(i)

  fmt.Println("The sorted string are:", s)
  fmt.Println("The sorted ints are:", i)

  fmt.Println("The ints are sorted?:", sort.IntsAreSorted(i))
}