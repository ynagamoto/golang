package main

import (
  "fmt"
)

func main() {
  var a, b, c, d int
  // fmt.Scan(&a)
  fmt.Scan(&a, &b, &c, &d)

  taka := 60*60*a + 60*b
  ao := 60*60*c + 60*d + 1

  if taka < ao {
    fmt.Println("Takahashi")
  } else {
    fmt.Println("Aoki")
  }
}
