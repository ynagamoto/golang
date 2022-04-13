package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
  names := make([][]string, n)
  nmap := make(map[string]*map[string]int)

  for i := 0; i < n; i++ {
    tmp := make([]string, 2)
    fmt.Scan(&tmp[0], &tmp[1])
    names[i] = tmp
    if nmap[tmp[1]] == nil {
      fname_map := make(map[string]int)
      nmap[tmp[1]] = &(fname_map)
    }
    (*nmap[tmp[1]])[tmp[0]]++
    // fmt.Println(nmap[tmp[1]])
  }
  // fmt.Println(nmap)

  flag := false
  for _, name := range names {
    fname, lname := name[0], name[1]
    if (*nmap[lname])[fname] >= 2 {
      flag = true
      break
    }
    if nmap[fname] != nil {
      fmt.Println(nmap[fname])
      if (*nmap[fname])[lname] >= 1 {
        flag = true
        break
      }
    }
  }

  if !flag {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}
