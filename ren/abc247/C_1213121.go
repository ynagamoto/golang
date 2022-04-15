package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
  ans := make([][]int, n)
  ans[0] = []int{1}

  for i := 1; i <= n-1; i++ {
    tmp := append(ans[i-1], i+1)
    ans[i] = append(tmp, ans[i-1]...)
  }

  for _, value := range ans[n-1] {
    fmt.Print(value, " ")
  }
  fmt.Println()
}
