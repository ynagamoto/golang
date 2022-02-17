package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
  // input
  var n, m int
  fmt.Scan(&n)
  fmt.Scan(&m)
  /* s := make([]string, n)
  t := make([]string, m)
  for i := 0; i < n; i++ {
    fmt.Scan(&s[i])
  }
  for i := 0; i < m; i++ {
    fmt.Scan(&t[i])
  } */
  sc.Scan()
  s_line := sc.Text()
  sc.Scan()
  t_line := sc.Text()
  s := strings.Split(s_line, " ")
  t := strings.Split(t_line, " ")

  // count
  for i,j := 0, 0; i < n && j < m; i++ {
    if s[i] == t[j] {
      fmt.Println("Yes")
      j++
    } else {
      fmt.Println("No")
    }
  }
}
