package main 

import(
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    up := strings.ToUpper(scanner.Text())
    fmt.Println(up)
  }

}