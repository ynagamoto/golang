package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)

  nums := make([]int, 4*n-1)
  for i, _ := range nums {
    fmt.Scan(&nums[i])
  }

  // count
  counts := make([]int, n)
  for _, num := range nums {
    counts[num-1]++
  }

  // output
  for i, count := range counts {
    if count == 3 {
      fmt.Println(i+1)
      break
    }
  }
}
