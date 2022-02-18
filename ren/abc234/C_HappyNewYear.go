package main

import (
  "fmt"
)

func main() {
  var k int
  fmt.Scan(&k)
  var b_runes []rune
  for ;; {
    quo := k/2
    rem := k%2

    if rem == 1 {
      rem = 2
    }

    if !(quo == 0 && rem == 0) {
      b_runes = append([]rune{'0'+rune(rem)}, b_runes...)
    }

    if quo == 0 {
      break
    } else {
      k = k/2
    }
  }
  ans := string(b_runes)
  fmt.Println(ans)
}
