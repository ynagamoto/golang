package main

import (
  "fmt"
)

func main() {
  var x, y int
  fmt.Scan(&x, &y)

  ans := 0
  if x <= y {
    lack := y - x
    ans = lack / 10
    if lack%10 > 0 {
      ans++
    }
  }
  fmt.Println(ans)
}
