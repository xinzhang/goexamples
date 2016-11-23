package main

import "fmt"
import "github.com/FlashBoys/go-finance"

func main() {
	// q, err := finance.AUDUSD
	// if err == nil {
	// 	fmt.Println(q)
	// }

	q, err := finance.GetQuote("AAPL")
	if err == nil {
		fmt.Println(q)
	}

}
