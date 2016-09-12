package main 

import "syscall"
import "os"
import "os/exec"

func main() {
  
  binary, _ := exec.LookPath("ls")

  args := []string{"ls", "-a", "-l"}
  env := os.Environ()

  err := syscall.Exec(binary, args, env)
  if err != nil {
    panic(err)
  }
}