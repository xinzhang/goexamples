package main

import (
    "encoding/json"
    "fmt"
)

type Response1 struct {
  Page int
  Fruits []string
}

func main() {
  fltB, _ := json.Marshal(2.34)
  fmt.Println(string(fltB))

  resp1 := Response1 {
    Page:3,
    Fruits: []string{"Apple", "Orange", "Pear"} }

  obj1, _ := json.Marshal(resp1)
  fmt.Println(string(obj1))

  str := `{"page": 1, "fruits": ["apple", "peach"]}`
  res2 := Response1{}
  json.Unmarshal([]byte(str), &res2)
  fmt.Println(res2)

  fmt.Println("Page=", res2.Page)
  fmt.Println("Fruits=", len(res2.Fruits), res2.Fruits[0])

}
