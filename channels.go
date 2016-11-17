package main

import "fmt"

func main() {
	message := make(chan string, 2)

	//go func() {
	message <- "ping"

	//message <- "pong"
	//}()

	fmt.Println(<-message)
	//fmt.Println(<-message)
}
