package main

import (
  "fmt"
)

var n, x int
var a_arr [][]int

func main() {
  fmt.Scan(&n, &x)
  l_arr := make([]int, n)
  a_arr = make([][]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&l_arr[i])
    tmp := make([]int, l_arr[i])
    for j := 0; j < l_arr[i]; j++ {
      fmt.Scan(&tmp[j])
    }
    a_arr[i] = tmp
  }

  count := search(0, 1)
  fmt.Println(count)
}

func search(i int, mul int) int {
  count := 0
  for k := 0; k < len(a_arr[i]); k++ {
    if i == n-1 {
      if x == mul*a_arr[i][k] {
        count++
      }
    } else {
      if x < mul*a_arr[i][k] {
        continue
      } else {
        count += search(i+1, mul*a_arr[i][k])
      }
    }
  }
  return count
}
