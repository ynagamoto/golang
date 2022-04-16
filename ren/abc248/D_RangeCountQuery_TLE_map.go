package main

import (
  "fmt"
)

func main() {
  var n, q int
  fmt.Scan(&n)
  count_maps := make([]map[int]int, n)
  for i := 0; i < n; i++ {
    count_maps[i] = make(map[int]int, n)
  }
  for i := 0; i < n; i++ {
    var value int
    fmt.Scan(&value)
    add_map(i, value-1, n, count_maps)
  }

  fmt.Scan(&q)
  for i := 0; i < q; i++ {
    var l, r, x int
    fmt.Scan(&l, &r, &x)
    l--
    r--
    x--
    if l == 0 {
      fmt.Println(count_maps[r][x])
    } else {
      fmt.Println(count_maps[r][x] - count_maps[l][x])
    }
  }
}

func add_map(s int, x int, n int, count_maps []map[int]int) {
  for i := s; i < n; i++ {
   count_maps[i][x]++
  }
}
