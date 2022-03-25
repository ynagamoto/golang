package main

import (
  "fmt"
)

func main() {
  var v int
  arr := make([]int, 3)
  fmt.Scan(&v)
  for i := 0; i < 3; i++ {
    fmt.Scan(&arr[i])
  }

  for i := 0; v >= 0; i = (i+1)%3 {
    v = v - arr[i]
    if v < 0 {
      switch i {
        case 0:
          fmt.Println("F")
        case 1:
          fmt.Println("M")
        case 2:
          fmt.Println("T")
      }
      break
    }
  }
}
