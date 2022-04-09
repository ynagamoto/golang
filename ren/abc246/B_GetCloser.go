package main

import (
  "fmt"
  "math"
)

func main() {
  var a, b float64 
  fmt.Scan(&a, &b)

  // x <- cos, y <- sin
  // r = 1

  // sin = y/r -> y = r*sin
  // cos = x/r -> x = r*cos

  r := math.Sqrt(a*a+b*b)
  sin := b/r
  cos := a/r
  x := cos
  y := sin

  fmt.Println(x, y)
}
