package main

import (
  "fmt"
)

func main1() {
  var m map[string]int

  m = make(map[string]int)

  m["route"] = 3
  m["abc"] = 5
  m["birdy"] = -1
  m["eagle"] = -2

  fmt.Println(m)

  i:=m["birdy111"]
  fmt.Println(i)

  fmt.Println("map length", len(m))
  delete(m, "abc")

  fmt.Println(m)
}

func main() {
  type Person struct {
       Name  string
       Likes []string
   }
   var people []*Person

   likes := make(map[string][]*Person)
   for _, p := range people {
       for _, l := range p.Likes {
           likes[l] = append(likes[l], p)
       }
   }
   for _, p := range likes["cheese"] {
        fmt.Println(p.Name, "likes cheese.")
    }
}
