package main

import (
  "fmt"
  "math"
)

func main() {
  var n int
  fmt.Scan(&n)

  var s, ans int
  max := int(math.Pow10(6))
  for i := 0; ; i++ {
    ah3 := int(math.Pow(float64(i), 3))
    tmp := 4*ah3
    if tmp >= n {
      ans = tmp
      s = i
      break
    }
  }
  // fmt.Println(s, max)

  outSide: for i, k := s, s; i <= max; i++ {
    for j := k; j >= 0; j-- {
      tmp := int(math.Pow(float64(i), 3)) + int(math.Pow(float64(j), 3)) + i*j*(i+j)
      // fmt.Println(i, j, tmp)
      if n <= tmp && ans >= tmp {
        ans = tmp
      }
      if n >= tmp {
        k = j
        continue outSide
      }
    }
  }

  fmt.Println(ans)
}
