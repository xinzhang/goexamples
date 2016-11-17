package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "log"
  //"strings"
  "net/http"
  "encoding/json"
  "gopkg.in/mgo.v2"
  //"gopkg.in/mgo.v2/bson"
)

type Person struct {
  Name string
  Phone string
}

func Handler(response http.ResponseWriter, request *http.Request) {
  response.Header().Set("Content-type", "text/html")
  webpage, err := ioutil.ReadFile("index.html")
  if err != nil {
    http.Error(response, fmt.Sprintf("home.html file error %v", err), 500)
  }
  
  fmt.Fprint(response, string(webpage))
  //fmt.Println(response, string(webpage))
}

func APIHandler(response http.ResponseWriter, request *http.Request) {

  //set mimetype to json
  response.Header().Set("Content-type", "application/json")
  err := request.ParseForm()
  if (err != nil) {
    http.Error(response, fmt.Sprintln("error parsing url %v",err), 500)
  }

  //get mongodb conn session
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }

  defer session.Close()

  session.SetMode(mgo.Monotonic, true)

  //get the people collection
  //c := session.DB("test").C("people")

  var result = make([]string, 1000)
  switch request.Method {
    case "GET":

      p1 := Person{Name:"Smith", Phone:"12312"}
      b, err := json.Marshal(p1)
      if (err != nil) {
        fmt.Println(err)
        return
      }
      result[0] = fmt.Sprintf("%s", string(b))

    case "POST":
      name := request.PostFormValue("name")
      fmt.Println(name)

      result[0] = "{Id:4, Result:true}"

    default:
  }

  json,err := json.Marshal(result)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Fprintf(response, "%v", string(json))
}


func main() {

  port := 1234
  var err string
  portstring := strconv.Itoa(port)

  mux := http.NewServeMux()
  mux.Handle("/api/", http.HandlerFunc(APIHandler))
  mux.Handle("/", http.HandlerFunc(Handler))

  log.Println("Listening on port " + portstring + " ... ")
  errs := http.ListenAndServe(":" + portstring, mux)
  if errs != nil {
      log.Fatal("ListenAndServerError: ", err)
  }

}
