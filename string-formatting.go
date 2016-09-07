package main 

import "fmt"

var pr = fmt.Printf

type point struct {
  x,y int
}

func main() {
  p := point{1,2}
  pr("%v\n", p)
  pr("%+v\n", p)
  pr("%#v\n", p)

  pr("%T\n", p)
  pr("%t\n", true)
  pr("%b\n", 7)
  pr("%x\n", 11)
  pr("%c\n", 65)
  pr("%e\n", 1230000000.0)
  pr("|%-6.2f|%-6.2f|\n", 1.2, 1.234)
}