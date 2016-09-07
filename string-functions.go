package main 

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
  p("Contains: ", s.Contains("abc", "a"))  
  p("Count: ", s.Contains("abcabcaaa", "a"))
  p("HasPrefix: ", s.HasPrefix("abc", "a"))
  p("HasSuffix: ", s.HasSuffix("abc", "c"))
  p("Index:", s.Index("abc", "b"))
  p("Repeat:", s.Repeat("*", 50))
  p("ToLower:", s.ToLower("Hello World"))
}