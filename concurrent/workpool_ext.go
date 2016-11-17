package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type DataResult struct {
	WorkerId  int
	FromValue int
	Result    int
	Received  time.Time
}

func (dr DataResult) print() {
	b, err := json.Marshal(dr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func worker(id int, jobs <-chan int, results chan<- *DataResult) {
	for j := range jobs {
		//fmt.Println("worker ", id, " processing job ", j)
		time.Sleep(time.Second * 3)

		d := &DataResult{WorkerId: id, FromValue: j, Result: j * 2, Received: time.Now()}
		results <- d
	}

}

func main() {
	fmt.Println(time.Now(), "start")

	jobs := make(chan int, 100)
	results := make(chan *DataResult, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 15; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 15; a++ {
		r := <-results
		r.print()
	}

}
