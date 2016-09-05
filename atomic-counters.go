package main 

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {
  var o uint64 = 0
  for i := 0; i < 50; i = i + 1 {
    go func () {
      for {
        atomic.AddUint64(&o, 1)
        runtime.Gosched()
      }
    }()
  }

  time.Sleep(time.Second)
  final := atomic.LoadUint64(&o)
  fmt.Println("The final value is:", final)
}