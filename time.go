package main 

import "fmt"
import "time"

var p = fmt.Println

func main() {
  now := time.Now()
  p(now)

  p("Year:", now.Year())
  p("Month:", now.Month())
  p("Day:", now.Day())
  p("Minute:", now.Minute())

  then := time.Date(2015, 5, 10, 22, 5, 10, 0, time.UTC)
  p(then)

  p(then.Before(now)) 

  diff := now.Sub(then)

  p("Diff between diff:", diff)

  p("Diff hours is:", diff.Hours())

  p(then.Add(diff))
}