package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	dest := "http://localhost:54560/api/eventtypes"

	v := url.Values{}
	v.Set("evt", "abcd")

	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}
	req, _ := http.NewRequest("POST", dest, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value")
	fmt.Printf("%+v\n", req)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
}
