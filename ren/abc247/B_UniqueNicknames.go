package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
  names := make([][]string, n)

  for i := 0; i < n; i++ {
    tmp := make([]string, 2)
    fmt.Scan(&tmp[0], &tmp[1])
    names[i] = tmp
  }

  flag := false
  for i, full_name := range names {
    for j, name := range full_name {
      if name != "" {
        check := false
        for k, tmp := range names[i+1:] {
          curr := i+k+1
          ftmp, ltmp := tmp[0], tmp[1]
          if name == ftmp {
            check = true
            names[curr][0] = ""
          }
          if name == ltmp {
            check = true
            names[curr][1] = ""
          }
        }
        if check {
          names[i][j] = ""
        }
      }
    }
    if full_name[0] == "" && full_name[1] == "" {
      flag = true
      break
    }
  }

  if !flag {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}
