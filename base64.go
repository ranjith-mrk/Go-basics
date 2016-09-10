package main 

import "fmt"
import b64 "encoding/base64"

func main() {
  str := "alskdfklafa;sfj;ajsfq09adfasdf080as"

  sE := b64.StdEncoding.EncodeToString([]byte(str))
  fmt.Println("Standard encoded:", sE)

  sD, _ := b64.StdEncoding.DecodeString(sE)
  fmt.Println(string(sD))

}