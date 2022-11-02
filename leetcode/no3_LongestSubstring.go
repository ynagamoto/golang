package main

import (
  "fmt"
)

func lengthOfLongestSubstring(s string) int {
  res := 0
  check_map := make(map[string]int)    
  for beg, end := 0, 0; end < len(s); end++ {
    curr := string(s[end])
    if beg == 0 && end == 0 { // init
      check_map[curr]++
      res++
    } else {
      if check_map[curr] == 0 {
        check_map[curr]++
        if end == len(s)-1 {
          if res < end-beg+1 {
            res = end-beg+1
          }
        }
      } else { // appeared the knew number
        // check len
        if res < end-beg {
          res = end-beg
        }
        // move beg
        for ; beg < end; beg++ {
          if string(s[beg]) == curr {
            beg++
            break
          }
          check_map[string(s[beg])] = 0
        }
      }
    }
  }
  return res
}

func main() {
  // s := "abcadc"
  // s := "abcabcbb"
  // s := "bbbbb"
  // s := "au"
  s := "tmmzuxt"
  fmt.Printf("%d\n", lengthOfLongestSubstring(s))
}
