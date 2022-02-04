package main

import (
  "fmt"
  "math"
)

func main() {
  var n float64
  fmt.Scan(&n)

  min := -1 * math.Pow(2, 31)
  max := math.Pow(2, 31)

  if n >= min && n < max {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}
