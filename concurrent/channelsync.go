package main

import "fmt"
import "time"

func worker(done chan bool) {
  fmt.Print("working...")
  time.Sleep(time.Second * 2)
  fmt.Println("done")

  done<-true
}


func ping(pings chan<- string, msg string) {
    pings <-msg
}

func pong(pings <-chan string, pongs chan<- string) {
  msg := <-pings
  pongs <- msg
}

func main() {
  //channel sync
  // chan1 := make (chan bool, 2)
  // go worker(chan1)
  // go worker(chan1)
  // <-chan1
  // <-chan1

  pings := make(chan string, 1)
  pongs := make(chan string, 1)

  ping(pings, "hello world");
  pong(pings, pongs);

  fmt.Println( <- pongs)
}
