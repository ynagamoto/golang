package main

import (
  "fmt"
)

var que_f, que_r int
var que [][]int

func main() {
  var n int
  fmt.Scan(&n)

  que_f, que_r = 0, 0
  que = make([][]int, n)
  for i := 0; i < n; i++ {
    var q, c int
    fmt.Scan(&q)
    switch q {
      case 1:
        var x int
        fmt.Scan(&x, &c)
        push(x, c)
      case 2:
        fmt.Scan(&c)
        fmt.Println(pop(c))
    }
  }
}

func push(x int, c int) {
  tmp := []int{x, c}
  que[que_r] = tmp
  que_r++
}

func pop(c int) int {
  sum := 0
  for ; c > 0; {
    if que[que_f][1] < c {
      sum += que[que_f][0] * que[que_f][1]
      c -= que[que_f][1]
      que_f++
    } else {
      sum += que[que_f][0] * c
      que[que_f][1] -= c
      c = 0
    }
  }
  return sum
}
