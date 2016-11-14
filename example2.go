package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

func main()  {
    var files []os.FileInfo
    var err error
    files, err = ioutil.ReadDir(path)

    if err != nil {
      fmt.Println(err)
      return
    }

    for _,info := range files {
      fmt.Println(info.Name())
    }
    
}
