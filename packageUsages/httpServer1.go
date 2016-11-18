package main

import (
	"fmt"
	"log"
	"net/http"
)

// method 1
// type indexHandler struct {
// }
//
// func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to Go Web Development")
// }

//method 2
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development")
}

func main() {
	//mux := http.NewServeMux()
	//method 1
	//index := &indexHandler{}

	//method 2
	//index := http.HandlerFunc(indexHandler)
	//mux.Handle("/", index)

	log.Println("Listening...")
	//http.ListenAndServe(":8008", mux)

	//method 3
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8008", nil)
}
