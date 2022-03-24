package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
  var s string
  fmt.Scan(&s)
  s_num := len(s)
  s_runes := []rune(s)
  fmt.Println(string(s_runes[s_num-1]))
}
