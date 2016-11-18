package main

import "fmt"

func main() {
	//first batch
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//load the channel values
	for j := range requests {
		fmt.Println(j)
	}

}
