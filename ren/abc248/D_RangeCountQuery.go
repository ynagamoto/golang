package main

import (
  "fmt"
  "sort"
)

func main() {
  var n, q int
  fmt.Scan(&n)
  count_arr := make([][]int, n)
  for i := 0; i < n; i++ {
    var value int
    fmt.Scan(&value)
    count_arr[value-1] = append(count_arr[value-1], i)
  }

  fmt.Scan(&q)
  for i := 0; i < q; i++ {
    var l, r, x int
    fmt.Scan(&l, &r, &x)
    l--
    r--
    x--
    low, high := 0, 0
    if count_arr[x] != nil {
      low = sort.Search(
        len(count_arr[x]),
        func(i int) bool {
          return count_arr[x][i] >= l
        },
      )
      high = sort.Search(
        len(count_arr[x]),
        func(i int) bool {
          return count_arr[x][i] >= r+1
        },
      )
    }
    fmt.Println(high-low)
  }
}
