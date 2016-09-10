package main 

import "fmt"
import "time"

var p = fmt.Println

func main() {
  t := time.Now()
  p(t.Format(time.RFC3339))  

  p("Formatting to time:", t.Format("10:00PM"))
  p("Another test:", t.Format("Sun Jan 1 20:00:00 2007"))

  tp, _ := time.Parse(time.RFC3339, "2015-06-11T06:20:00+00:00")
  p(tp)

}