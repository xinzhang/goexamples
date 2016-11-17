package main

import (
  "fmt"
)

func main() {
  var numbers []int = []int{3,4,6,7,3,4,1,2, 9, 3,2}
  fmt.Println("initial value", numbers)

  bubbleSort(numbers)
  fmt.Println("after sort", numbers)
}

func bubbleSort(numbers []int ){
  var N int = len(numbers)
  var i int
  for i=0;i<N;i++ {
    sweep(numbers)
  }
}

func sweep(numbers []int) {
  var N int = len(numbers)
  var firstIndex int = 0
  var secondIndex int = 1

  for secondIndex < N {
    var firstNumber int = numbers[firstIndex]
    var secondNumber int = numbers[secondIndex]

    if firstNumber > secondNumber {
      numbers[firstIndex] = secondNumber
      numbers[secondIndex] = firstNumber
    }

    firstIndex++
    secondIndex++
  }
}
