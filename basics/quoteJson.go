package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var m map[string]interface{}

func main() {

	contents, err := ioutil.ReadFile("quote.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(contents, &m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success", m["success"].(map[string]interface{})["total"])

	c := m["contents"].(map[string]interface{})
	fmt.Println(c["copyright"])

	arr := c["quotes"].([]interface{})
	fmt.Println(arr[0].(map[string]interface{})["quote"])

}
