package main

import (
  "fmt"
)

func main() {
  var n int
  var t string
  fmt.Scan(&n)
  fmt.Scan(&t)

  x, y := 0, 0
  dire := 0 // 0: 右，1: 下，2: 左，3:上
  t_runes := []rune(t)
  for _, t_rune := range t_runes {
    if t_rune == 'S' {
      switch dire {
        case 0:
          x++
        case 1:
          y--
        case 2:
          x--
        case 3:
          y++
      }
    } else {
      dire = (dire+1) % 4
    }
  }
  fmt.Println(x, y)
}
