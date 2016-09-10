package main 

import "fmt"
import "strconv"

var p = fmt.Println
func main() {
  f1, _ := strconv.ParseFloat("1.11111111111", 2)
  p(f1)
  i1, _ := strconv.ParseInt("200", 0, 64)
  p(i1)

  ih, _ := strconv.ParseInt("0x111", 0, 64)
  p(ih)

  a, _ := strconv.Atoi("9899")
  p(a)

}