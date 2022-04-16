package main

import (
  "fmt"
)

func main() {
  var n, q int
  fmt.Scan(&n)
  count_arr := make([][]int, n)
  for i := 0; i < n; i++ {
    var value int
    fmt.Scan(&value)
    tmp := make([]int, n)
    if i == 0 {
      tmp[value-1]++
    } else {
      copy(tmp, count_arr[i-1])
      tmp[value-1] = count_arr[i-1][value-1] + 1
    }
    count_arr[i] = tmp
  }
  /* for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      fmt.Print(count_arr[j][i], " ")
    }
    fmt.Println()
  } */

  fmt.Scan(&q)
  for i := 0; i < q; i++ {
    var l, r, x int
    fmt.Scan(&l, &r, &x)
    l--
    r--
    x--
    if l == 0 {
      fmt.Println(count_arr[r][x])
    } else {
      fmt.Println(count_arr[r][x] - count_arr[l][x])
    }
  }
}
