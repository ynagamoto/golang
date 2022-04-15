package main

import (
  "fmt"
)

func main() {
  var n, x, y int
  fmt.Scan(&n, &x, &y)
  arr := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&arr[i])
  }

  ans := 0
  l_count := 0
  for i, l := range arr {
    if l < y || l > x {
      l_count = 0
      continue
    }

    l_count++
    xflag, yflag := false, false
    if l == x {
      xflag = true
    }
    if l == y {
      yflag = true
    }

    if xflag || yflag {
      r_count := 0
      for _, r := range arr[i:] {
        if r < y || x < r {
          break
        }

        if r == x {
          xflag = true
        }
        if r == y {
          yflag = true
        }

        if xflag && yflag {
          r_count++
        }
      }
      if xflag && yflag {
        ans += l_count * r_count
        l_count = 0
      }
    }
  }

  fmt.Println(ans)
}
