package main

import "fmt"
import "time"

func main() {
  fmt.Println(time.Now(), "start")
  c1 := make(chan string)

  go func() {
    time.Sleep(time.Second * 4)
    c1 <- "one"
  }()

  select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }

  c2 := make(chan string)
  go func()  {
    time.Sleep(time.Second * 3)
    c2 <- "result two"
  }()

  select {
  case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 6):
        fmt.Println("timeout 2")
    }
}
