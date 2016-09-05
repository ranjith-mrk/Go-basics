package main 

import (
  "fmt"
  "time"
  "math/rand"
  "sync/atomic"
)

type readOp struct {
  key int
  resp chan int
}

type writeOp struct {
  key int
  value int
  resp chan bool
}

func main() {
  var ops int64 = 0

  reads := make(chan *readOp)
  writes := make(chan *writeOp)

  go func () {
    var state = make(map[int]int)
    for {
      select {
        case read := <-reads :
          read.resp <- state[read.key]
        case write := <-writes :
          state[write.key] = write.value 
          write.resp <- true
      }
    }
  }()

  for r := 0 ; r < 100; r = r + 1 {
    go func () {
      for {
        read := &readOp{
          key: rand.Intn(100),
          resp: make(chan int)}
        reads <- read
        <-read.resp
        atomic.AddInt64(&ops, 1)
      }
    }()
  }

  for w := 0; w < 10; w = w + 1 {
    go func () {
      for {
        write := &writeOp{
          key: rand.Intn(5),
          value: rand.Intn(100),
          resp: make(chan bool)}
        writes <- write 
        <-write.resp
        atomic.AddInt64(&ops, 1)
      }
    }()
  }

  time.Sleep(time.Second)

  final := atomic.LoadInt64(&ops)
  fmt.Println("The final value is:",final)
}