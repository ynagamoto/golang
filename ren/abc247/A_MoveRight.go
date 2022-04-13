package main

import (
  "fmt"
)

func main() {
  var n_str string
  fmt.Scan(&n_str)
  // weight := []int{8, 4, 2, 1}
  n_runes := []rune(n_str)
  ans := append([]rune{'0'}, n_runes[:3]...)

  for _, r := range ans {
    fmt.Print(string(r))
  }
  fmt.Println()
}
