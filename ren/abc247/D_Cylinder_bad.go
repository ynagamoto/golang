package main

import (
  "fmt"
)

var queue []int

func main() {
  var n int
  fmt.Scan(&n)

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
  tmp := make([]int, c)
  for i, _ := range tmp {
    tmp[i] = x
  }
  queue = append(queue, tmp...)
}

func pop(c int) int {
  sum := 0
  for _, value := range queue[:c] {
    sum += value
  }
  queue = queue[c:]
  return sum
}
