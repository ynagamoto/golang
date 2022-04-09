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
    points[i] = tmp
  }

  var max_x, max_y, min_x, min_y int
  for i, point := range points {
    if i == 0 {
      max_x, max_y = point[0], point[1]
    } else {
      if point[0] > max_x {
        min_x = max_x
        max_x = point[0]
      } else {
        min_x = point[0]
      }
      if point[1] > max_y {
        min_y = max_y
        max_y = point[1]
      } else {
        min_y = point[1]
      }
    }
  }

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
