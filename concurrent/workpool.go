package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, " processing job ", j)
		time.Sleep(time.Second * 3)
		results <- j * 2
	}
}

func main() {
	fmt.Println(time.Now(), "start")

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 15; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 15; a++ {
		r := <-results
		fmt.Println("result: ", r)
	}

}
