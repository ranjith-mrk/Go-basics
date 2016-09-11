package main

import(
  "fmt"
  "bufio"
  "io/ioutil"
  "os"
)

func check(e error){
  if e != nil {
    panic(e)
  }
}

func main() {
  data, err := ioutil.ReadFile("/tmp/dat")
  check(err)
  fmt.Println("Raw data is", data)
  fmt.Println(string(data))

  f, err := os.Open("/tmp/dat")
  check(err)

  b5 := make([]byte, 5)
  no, err := f.Read(b5)
  check(err)
  fmt.Printf("%d bytes are read for data: %s\n", no, b5)

  b2 := make([]byte, 2)

  o, err := f.Seek(6, 0)
  check(err)

  f2, err := f.Read(b2)

  fmt.Printf("The characters read from position is %d, with bytes %d with returning value: %s\n", o, f2, b2)

  buf := bufio.NewReader(f)

  _, e := f.Seek(0, 0)
  check(e)

  bbuf, e := buf.Peek(8)
  check(e)
  fmt.Printf("Data via buffer is: %s\n", bbuf)


}