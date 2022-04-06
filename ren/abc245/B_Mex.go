package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)

  max_len := 2000 + 1
  check_arr := make([]int, max_len)
  for i := 0; i < n ; i++ {
    var tmp int
    fmt.Scan(&tmp)
    check_arr[tmp]++
  }

  for i, check := range check_arr {
    if check == 0 {
      fmt.Println(i)
      break
    }
  }
}
