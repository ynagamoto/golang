package main

import (
  "fmt"
)

func main() {
  var a, b, k int
  fmt.Scan(&a, &b, &k)

  for i := 0; ; i++ {
    if a >= b {
      fmt.Println(i)
      break
    }
    a *= k
  }
}
