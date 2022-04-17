package main

import (
  "fmt"
)

func main() {
  var l, r int
  var s string
  fmt.Scan(&l, &r, &s)
  l--
  r--

  s_r := []rune(s)
  s_f := s_r[:l]
  s_l := s_r[r+1:]
  before := s_r[l:r+1]
  b_len := len(before)
  rev := make([]rune, b_len)
  for i := 0; i < b_len; i++ {
    rev[i] = before[b_len-1-i]
  }

  ans := append(s_f, rev...)
  ans = append(ans, s_l...)
  fmt.Println(string(ans))
}
