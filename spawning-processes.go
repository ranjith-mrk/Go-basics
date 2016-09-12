package main 

import "fmt"
import "os/exec"
import "io/ioutil"

func main() {
  dateCmd := exec.Command("date")

  dateOut, _ := dateCmd.Output()

  fmt.Println(">date")
  fmt.Println(string(dateOut))

  grepCmd := exec.Command("grep", "goodbye")

  grepIn, _  := grepCmd.StdinPipe()
  grepOut, _ := grepCmd.StdoutPipe()

  grepCmd.Start()

  grepIn.Write([]byte("hello grep\ngoodbye grep"))

  grepIn.Close()

  grepBytes, _ := ioutil.ReadAll(grepOut)
  grepCmd.Wait()

  fmt.Println("> grep goodbye")
  fmt.Println(string(grepBytes))
}