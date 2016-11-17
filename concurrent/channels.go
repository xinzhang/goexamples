package main

import "fmt"

func main() {
  messages := make (chan string, 2)

  //go func() { messages <- "ping"} ()
  messages<-"buffer"
  messages<-"channel"

  fmt.Println(<-messages)
  fmt.Println(<-messages)
}
