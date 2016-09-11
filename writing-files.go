package main 

import(
  "fmt"
  "io/ioutil"
  "bufio"
  "os"
)


func check(e error) {
  if e != nil {
    panic(e)
  }
}


func main() {
  fdata := []byte("hello\nworld\n")

  e := ioutil.WriteFile("/tmp/dat1", fdata, 0644)
  check(e)

  f, e := os.Create("/tmp/dat2")
  check(e)

  str11 := []byte{115, 111, 109, 101, 10}
  os1, e := f.Write(str11)
  check(e)

  fmt.Printf("No of bytes:%d\n", os1)

  str12 := "goodness\n"
  os2, e := f.WriteString(str12)
  fmt.Printf("No of %d bytes\n", os2)
  check(e)

  w := bufio.NewWriter(f)
  os3, e := w.WriteString("within\n")
  check(e)
  fmt.Printf("No of %d bytes\n", os3)

  f.Sync()
  w.Flush()
}