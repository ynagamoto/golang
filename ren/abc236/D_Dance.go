package main

import (
  "fmt"
)

var n, ans int
var comp [][]int

func check(flags []bool) int {
  for i, flag := range flags {
    if !flag {
      return i
    }
  }
  return -1
}

func search(flags []bool, point int) {
  index := check(flags)
  if index == -1 {
    if point > ans {
        ans = point
    }
    return
  }
  flags[index] = true
  for i := index+1; i <= 2*n-1; i++ {
    if flags[i] {
      continue
    }
    flags_cp := make([]bool, 2*n)
    copy(flags_cp, flags)
    flags_cp[i] = true
    // fmt.Println("  ", index, i)
    search(flags_cp, point^comp[index][i-1-index])
  }
}

func main() {
  // input
  // var n int <- move global
  fmt.Scan(&n)
  arr := make([][]int, 2*n-1)
  for i := 0; i < 2*n-1; i++ {
    tmp := make([]int, 2*n-1-i)
    for j := 0; j < 2*n-1-i; j++ {
      fmt.Scan(&tmp[j])
    }
    arr[i] = tmp
  }
  comp = arr

  /* fmt.Println(n)
  for _, nums := range comp {
    for _, num := range nums {
      fmt.Print(num, " ")
    }
    fmt.Println()
  } */

  ans = 0
  flags := make([]bool, 2*n)
  search(flags, 0)

  // output
  fmt.Println(ans)
}
