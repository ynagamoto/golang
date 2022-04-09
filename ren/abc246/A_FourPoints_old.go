package main

import (
  "fmt"
)

func main() {
  points := make([][]int, 3)
  max_x, max_y, min_x, min_y := -1000, -1000, 1000, 1000
  for i := 0; i < 3; i++ {
    tmp := make([]int, 2)
    fmt.Scan(&tmp[0], &tmp[1])
    // x
    if max_x < tmp[0] {
      max_x = tmp[0]
    }
    if min_x > tmp[0] {
      min_x = tmp[0]
    }
    // y
    if max_y < tmp[1] {
      max_y = tmp[1]
    }
    if min_y > tmp[1] {
      min_y = tmp[1]
    }
    points[i] = tmp
  }
  // fmt.Println(max_x, min_x)
  // fmt.Println(max_y, min_y)

  ans := make([]int, 2)
  checks := make([]int, 4)
  for _, point := range points {
    if point[0] == max_x {
      if point[1] == max_y {
        checks[3]++
      } else {
        checks[1]++
      }
    } else {
      if point[1] == max_y {
        checks[2]++
      } else {
        checks[0]++
      }
    }
  }


  for i, check := range checks {
    if check == 0 {
      switch i {
        case 0:
          ans[0] = min_x
          ans[1] = min_y
        case 1:
          ans[0] = max_x
          ans[1] = min_y
        case 2:
          ans[0] = min_x
          ans[1] = max_y
        case 3:
          ans[0] = max_x
          ans[1] = max_y
      }
    }
  }

  fmt.Println(ans[0], ans[1])
}
