package main

import (
  "fmt"
  "log"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Person struct {
  Name string
  Phone string
}

func main() {
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }

  defer session.Close()


  session.SetMode(mgo.Monotonic, true)

  c:= session.DB("test").C("people")

  //Insert
  //err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
  //              &Person{"Cla", "+55 53 8402 8510"})

  // if err != nil {
  //    log.Fatal(err)
  // }

  //find one
       resultOne := Person{}
       err = c.Find(bson.M{"name": "Cla"}).One(&resultOne)
       if err != nil {
               log.Fatal(err)
       }

     fmt.Println("Phone:", resultOne.Phone)

  //get all
  //var result *[]Person
  result := []Person{}
  iter := c.Find(nil).Limit(100).Iter()
  err = iter.All(&result)

  if err != nil {
          log.Fatal(err)
          return
  }
  fmt.Println(result)
  for i:=0;i<len(result);i++ {
    fmt.Println(result[i])
  }


}
