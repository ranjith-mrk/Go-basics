package main 

import (
  "fmt"
  "time"
  "math/rand"
  "runtime"
  "sync"
  "sync/atomic"
)


func main() {
  state := make(map[int]int)

  mutex := &sync.Mutex{}

  var ops int64 = 0
  var t int = 0
  for i := 0; i < 100; i = i + 1 {
    go func () {
      total := 0
      for {
        key := rand.Intn(5)
        mutex.Lock()
        total += state[key]
        t += state[key]
        mutex.Unlock()
        atomic.AddInt64(&ops, 1)
        runtime.Gosched()
      }
    }()
  }

  for i := 0; i < 10; i = i + 1 {
    go func () {
      key   := rand.Intn(5)
      value := rand.Intn(1000)
      mutex.Lock()
      state[key] = value
      mutex.Unlock()
      atomic.AddInt64(&ops, 1)
      runtime.Gosched()
    }()
  }

  time.Sleep(time.Second)

  final := atomic.LoadInt64(&ops)
  fmt.Println("The final value is:",final)

  mutex.Lock()
  fmt.Println("The value of state is:", state)
  fmt.Println("The value of total is:", t)
  mutex.Unlock()

}