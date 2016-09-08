package main 

import (
  "fmt"
  "encoding/json"
)

type Employee struct {
    Name  string  `json:"name"`
    EmpId int       `json:"id"`
}

func main() {
  // intB, other := json.Marshal(1)
  i, _ := json.Marshal(1)
  fmt.Println("intB is:", string(i))
  // fmt.Println("Other is:", string(other))

  emp := Employee{"Ranjith", 99}
  empj, _ := json.Marshal(emp)
  fmt.Println("The Employee json object is:", string(empj))

  subs := map[string]int{"MATHS" : 196, "PHYSICS" : 197, "CHEMISTRY": 196}
  subsj, _ := json.Marshal(subs)
  fmt.Println("The subs json object is:", string(subsj))

  str := string(empj)
  e := Employee{}

  json.Unmarshal([]byte(str), &e)
  fmt.Println("Unmarshalled Employee is:", e)

}