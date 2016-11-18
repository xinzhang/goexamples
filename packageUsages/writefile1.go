package main

import "io/ioutil"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("./tmp/dat1", d1, 0644)

	check(err)
}
