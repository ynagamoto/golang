package main

import (
  "fmt"
)

func f(t int) int {
  return t*t+2*t+3
}

func main() {
  var t int
  fmt.Scan(&t)
  fmt.Println(f(f(f(t)+t) + f(f(t))))
}
