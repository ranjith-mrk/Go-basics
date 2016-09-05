package main 

import "fmt"
import "time"

func worker(id int, result chan<- int, jobs <-chan int) {
  for j := range jobs{
    fmt.Println("Worker:",id, " job:", j)
    time.Sleep(time.Second)
    result <- j * 2
  }
}

func main() {
  jobs := make(chan int, 10)
  result := make(chan int, 10)

  for i := 1; i <= 3; i = i + 1 {
    go worker(i, result, jobs)    
  }

  for j := 1; j <= 10; j = j + 1 {
    jobs <- j
  } 
  close(jobs)
  for k := 1; k <= 10; k = k + 1 {
    <- result
  }

}