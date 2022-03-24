package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
  arr := make([]int, 2*n+1)

  for i, j := 0, 0; i < 2*n+1; i++ {
    if i%2 == 0 {
      k := j
      for ; k < 2*n+1; k++ {
        if arr[k] != 1 {
          fmt.Println(k+1)
          break
        }
      }
      j = k+1
    } else {
      var num int
      fmt.Scan(&num)
      arr[num-1] = 1
    }
  }
  // fmt.Println("You Lose.")
}
