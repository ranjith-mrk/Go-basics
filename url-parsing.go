package main 

import (
  "fmt"
  "net"
  "net/url"
)

func main() {
  s1 := "postgres://user:pass@host.com:5432/path?k=v#f"
  // s2 := "http://www.facebook.com/home?q=hello"
  u, err := url.Parse(s1)

  if err != nil {
    panic(err)
  }

  fmt.Println("Scheme is:", u.Scheme)

  fmt.Println("User name test:", u.User)
  fmt.Println("User exact name:", u.User.Username())
  pass, _ := u.User.Password()
  fmt.Println("Passowrd is:", pass)

  fmt.Println("Host:", u.Host)
  host, port, test := net.SplitHostPort(u.Host)
  fmt.Println("Host again:", host)
  fmt.Println("port :", port)
  fmt.Println("testing the third value:", test)

  params, _ := url.ParseQuery(u.RawQuery)

  fmt.Println(params)
}