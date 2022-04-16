package main

import (
  "fmt"
)

func main() {
  var s string
  fmt.Scan(&s)
  s_runes := []rune(s)
  check_arr := make([]int, 10)

  for _, r := range s_runes {
    index := int(r) - int('0')
    check_arr[index]++ 
  }

  for i, check := range check_arr {
    if check == 0 {
      fmt.Println(i)
      break
    }
  }
}
