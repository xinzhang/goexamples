package main

import "fmt"

func add(a int, b int) int {
  return a + b;
}

func vals() (int, int, int){
  return 3, 4, 5
}

func namedvals() (a int, b int) {
  a = 1
  b = 2
  return
}

func sum (nums ...int) int {
  total := 0
  for _, num := range nums {
    total += num
  }
  return total
}

func main(){
  fmt.Println( add(3, 5) )

  a, b, c := vals()
  fmt.Println(a,b,c)

  i,j := namedvals()
  fmt.Println(i, j)

  nums := []int{1, 2, 3, 4, 5, 6}
  fmt.Println( sum( 1, 2, 3, 4) )
  fmt.Println( sum(nums...))
  
}
