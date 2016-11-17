package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
)

func main(){
  proxy := os.Getenv("HTTP_PROXY")
  fmt.Println(proxy)
  
  resp, err := http.Get("http://example.com")

  if (err != nil) {
    fmt.Println(err)
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))
}
