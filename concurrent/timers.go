package main

import "fmt"
import "time"

func main() {
	fmt.Println(time.Now(), "start")
	// timer1:= time.NewTimer(time.Second * 2)
	//
	// <- timer1.C
	// fmt.Println(time.Now(), "timer 1 expired")
	//
	// timer2 := time.NewTimer(time.Second*3)
	// go func() {
	//   <- timer2.C
	//   fmt.Println(time.Now(), "timer 2 expired")
	// }()
	//
	// //time.Sleep(time.Second * 5)
	//
	// stop2 := timer2.Stop()
	// if stop2 {
	//   fmt.Println(time.Now(), "timer 2 stopped")
	// }

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println(time.Now(), "tikcer at", t)
		}
	}()

	time.Sleep(time.Second * 5)
	ticker.Stop()
	fmt.Println(time.Now(), "tikcer stopped")

}
