package main

import "fmt"

type person struct {
  name string
  age int
}

type rect struct {
  width,height int
}

func (r *rect) area() int {
  return r.width * r.height
}

func (r rect) perim() int {
  return r.width * 2 + r.height * 2
}

func main() {
    fmt.Println(person{"Ian", 10})
    fmt.Println(person{name:"John", age:5})

    r := rect {width:5, height:5}
    fmt.Println(r.area(), r.perim())

    rp := &r
    fmt.Println(rp.area(), rp.perim())
}
