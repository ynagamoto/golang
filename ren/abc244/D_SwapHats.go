package main

import (
  "fmt"
)

func main() {
  var s_f, s_s, s_t string
  var t_f, t_s, t_t string
  fmt.Scan(&s_f)
  fmt.Scan(&s_s)
  fmt.Scan(&s_t)
  fmt.Scan(&t_f)
  fmt.Scan(&t_s)
  fmt.Scan(&t_t)

  count := 0
  if t_f != s_f {
    if t_s == s_f {
      t_s = t_f
      t_f = s_f
    } else if t_t == s_f {
      t_t = t_f
      t_f = s_f
    }
    count++
  }

  if t_s != s_s {
    t_s = s_s
    t_t = s_t
    count++
  }

  if count%2 == 0 {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}
