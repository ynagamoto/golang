package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
  var n int
  fmt.Scan(&n)
  sc.Scan()
  line := sc.Text()

  nums := strings.Split(line, " ")
  counts := make([]int, n)
  for _, s := range nums {
    num, _ := strconv.Atoi(string(s))
    /* num, err := strconv.Atoi(string(s))
    if err != nil {
      break
    } */
    counts[num-1]++
  }

  for i, c := range counts {
    if c == 3 {
      fmt.Println(i+1)
      break
    }
  }
}
