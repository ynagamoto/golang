package main

import (
  "fmt"
  "math"
)

func main() {
  var n int
  fmt.Scan(&n)
  points := make([][]int, n)
  max_y := 0
  for i := 0; i < n; i++ {
    tmp := make([]int, 2)
    fmt.Scan(&tmp[0]) // x
    fmt.Scan(&tmp[1]) // y
    points[i] = tmp
    if max_y < tmp[1] {
      max_y = tmp[1]
    }
  }
  var s string
  fmt.Scan(&s)

  // move
  s_runes := []rune(s)
  check_arr := make([][]int, max_y+1) // [0] is <- min, [1] is -> max, [2] is moving right, [3] is moving left
  flag := false
  for i, s_rune := range s_runes {
    p_x := points[i][0]
    p_y := points[i][1]

    if check_arr[p_y] == nil {
      tmp := make([]int, 4)
      tmp[0] = int(math.Pow10(9))
      tmp[1] = 0
      tmp[2] = 0
      tmp[3] = 0
      check_arr[p_y] = tmp
    }

    update := false
    if s_rune == 'R' {
      check_arr[p_y][2]++
      if p_x < check_arr[p_y][0] {
        check_arr[p_y][0] = p_x
        update = true
      }
    } else {
      check_arr[p_y][3]++
      if p_x > check_arr[p_y][1] {
        check_arr[p_y][1] = p_x
        update = true
      }
    }

    if update {
      if check_arr[p_y][2] > 0 && check_arr[p_y][3] > 0 {
        if check_arr[p_y][0] <= check_arr[p_y][1] {
          flag = true
        }
      }
    }
  }

  // output
  if flag {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}
