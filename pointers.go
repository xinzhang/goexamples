package main

import "fmt"

type Student struct {
    Id int
    Name string
    Address string
}

func (s Student) Register() (int, error) {
  if s.Id <= 0 {
      return -1, fmt.Errorf("the value can't be less than 0")
  } else {
    return 10000 + s.Id, nil
  }
}

func main() {
  //var a Student
  s := Student{1, "John", "20 Will St"}
  p := &s
  fmt.Println(*p)
  s.Id = 2
  s.Name = "Hello"
  fmt.Println(*p)

  regId, err := s.Register()
  if (err != nil) {
    fmt.Println("error")
  }
  fmt.Println("register id ", regId)

  q := *p //value copy to a new
  (*p).Id = -34
  p.Name = "Smith"

  regId, err = p.Register()
  if (err != nil) {
    fmt.Println(err.Error())
  }
  fmt.Println("register id ", regId)

  fmt.Println(p)
  fmt.Println(&s)
  fmt.Println(&q)

  var a Student
  a = Student{3, "ff", "fffff"}

  fmt.Println(a)
}
