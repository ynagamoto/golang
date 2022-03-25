package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
  arr_a := make([]int, n)
  arr_b := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&arr_a[i])
  }
  for i := 0; i < n; i++ {
    fmt.Scan(&arr_b[i])
  }

  ans1 := 0
  ans2 := 0
  check_map := make(map[int]int)
  for i := 0; i < n; i++ {
    // fmt.Println(arr_a[i], arr_b[i])
    if arr_a[i] == arr_b[i] {
      ans1++
    }
    check_map[arr_a[i]]++
    check_map[arr_b[i]]++
  }

  for _, check := range check_map {
    // fmt.Println(i, check)
    if check >= 2 {
      ans2++
    }
  }

  fmt.Println(ans1)
  fmt.Println(ans2-ans1)
}
