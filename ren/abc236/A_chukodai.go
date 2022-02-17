package main

import (
  "fmt"
)

func main() {
  var s string
  var a,b int
  fmt.Scan(&s)
  fmt.Scan(&a)
  fmt.Scan(&b)
  s_runes := []rune(s)

  tmp := s_runes[a-1]
  s_runes[a-1] = s_runes[b-1]
  s_runes[b-1] = tmp

  fmt.Println(string(s_runes))
}
