package main

import (
  "fmt"
)

func main() {
  hoge := make(map[int]string)
  fuga := make(map[int][]string)

  if hoge[0] == "" {
    fmt.Println("hogehoge")
  }
}
