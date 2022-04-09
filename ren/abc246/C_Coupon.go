package main

import (
  "fmt"
  "sort"
)

func main() {
  var n, k, x int
  fmt.Scan(&n, &k, &x)

  values := make([]int, n)
  for i := 0; i < n; i++ {
    var tmp int
    fmt.Scan(&tmp)
    if k - tmp/x >= 0 {
      k = k - tmp/x
      tmp = tmp%x
    } else {
      tmp = tmp - k*x
      k = 0
    }
    values[i] = tmp
  }
  // fmt.Println(values, k)

  // sort
  sort.Ints(values)
  // fmt.Println(values)

  for i := n-1; i >= 0 && k > 0; i-- {
    values[i] = 0
    k--
  }

  // output
  ans := 0
  for _, value := range values {
    ans += value
  }
  fmt.Println(ans)
}
