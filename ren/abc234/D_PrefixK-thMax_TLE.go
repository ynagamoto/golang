package main

import (
  "fmt"
  "sort"
)

func main() {
  var n, k int
  fmt.Scan(&n, &k)
  arr := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&arr[i])
  }

  sorted_arr := make([]int, k+1)
  copy(sorted_arr, arr[:k-1])
  for i := k-1; i < n; i++ {
    if sorted_arr[k-1] < arr[i] {
      sorted_arr[k] = arr[i]
      sort.Sort(sort.Reverse(sort.IntSlice(sorted_arr)))
    }
    fmt.Println(sorted_arr[k-1])
  }
}
