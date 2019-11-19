package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// there is a json-to-go tool to create the type
type QuoteData struct {
	Success struct {
		Total int `json:"total"`
	} `json:"success"`
	Contents struct {
		Quotes []struct {
			Quote      string   `json:"quote"`
			Length     string   `json:"length"`
			Author     string   `json:"author"`
			Tags       []string `json:"tags"`
			Category   string   `json:"category"`
			Date       string   `json:"date"`
			Permalink  string   `json:"permalink"`
			Title      string   `json:"title"`
			Background string   `json:"background"`
			ID         string   `json:"id"`
		} `json:"quotes"`
		Copyright string `json:"copyright"`
	} `json:"contents"`
}

func main() {
	response, err := http.Get("http://api.theysaidso.com/qod")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var result QuoteData
	json.Unmarshal(contents, &result)
	//fmt.Println(string(contents))
	fmt.Println(result.Success.Total)
	for _, q := range result.Contents.Quotes {
		fmt.Println(q.Quote)
	}

}
