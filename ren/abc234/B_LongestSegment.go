package main

import (
  "fmt"
  "math"
)

func main() {
  var n int
  fmt.Scan(&n)
  points := make([][]int , n)
  for i, _ := range points {
    var x, y int
    fmt.Scan(&x)
    fmt.Scan(&y)
    points[i] = []int{x, y}
  }

  ans := 0.0
  for i := 0; i < n; i++ {
    for j := i+1; j < n; j++ {
      x_diff2 := math.Pow(float64(points[j][0]-points[i][0]), 2)
      y_diff2 := math.Pow(float64(points[j][1]-points[i][1]), 2)
      dis := math.Sqrt(x_diff2 + y_diff2)
      if dis > ans {
        ans = dis
      }
    }
  }

  fmt.Println(ans)
}
