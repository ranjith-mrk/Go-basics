package main 

import(
  "fmt"
  "time"
  "math/rand"
)

func main() {
  var p = fmt.Println
  p("Random int:", rand.Intn(1000))

  p("Random float:", rand.Float64())
  p("Random no with range:", rand.Float64() * 1000)

  source := rand.NewSource(time.Now().UnixNano())
  r := rand.New(source)

  p("Random with the new source:", r.Intn(1000))
}