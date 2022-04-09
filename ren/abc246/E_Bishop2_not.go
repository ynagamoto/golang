package main

import (
  "fmt"
)

var field [][]rune

func main() {
  var n int
  p_a := make([]int, 2)
  p_b := make([]int, 2)
  fmt.Scan(&n)
  fmt.Scan(&p_a[0], &p_a[1])
  fmt.Scan(&p_b[0], &p_b[1])
  field_r := make([][]rune, n)
  for i := 0; i < n; i++ {
    var tmp string
    fmt.Scan(&tmp)
    field_r[i] = []rune(tmp)
  }
  field = field_r

  fmt.Println(field)
}
