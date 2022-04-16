package main

import (
  "fmt"
)

var n, m, k int

func main() {
  fmt.Scan(&n, &m, &k)

  arr := make([]int, n)
  count := 0
  for tmp := 0; tmp != 1; tmp = add(0, arr) {
    if sum_arr(arr) <= k {
      count++
    }
  }
  fmt.Println(count%998244353)
}

func add(i int, arr []int) int {
  value := arr[i]+1 // init value is 0
  if value+1 <= m {
    arr[i]++
  } else {
    if i != n-1 {
      arr[i] = 0
      return add(i+1, arr)
    } else {
      return 1
    }
  }
  return 0
}

func sum_arr(arr []int) int {
  sum := 0
  for _, value := range arr {
    sum += value+1
  }
  return sum
}
