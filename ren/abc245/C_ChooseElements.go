package main

import (
  "fmt"
  "math"
)

func main() {
  var n, k int
  fmt.Scan(&n, &k)
  arrA := make([]int, n)
  arrB := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&arrA[i])
  }
  for i := 0; i < n; i++ {
    fmt.Scan(&arrB[i])
  }

  flag := true
  tmp := make([]int, 2)
  tmp[0] = arrA[0]
  tmp[1] = arrB[0]
  for i := 1; i <= n-1; i++ {
    tmpAA, tmpAB := k+1, k+1
    tmpBA, tmpBB := k+1, k+1

    if tmp[0] != -1 {
      tmpAA = int(math.Abs(float64(tmp[0]-arrA[i])))
      tmpAB = int(math.Abs(float64(tmp[0]-arrB[i])))
    }
    if tmp[1] != -1 {
      tmpBA = int(math.Abs(float64(tmp[1]-arrA[i])))
      tmpBB = int(math.Abs(float64(tmp[1]-arrB[i])))
    }

    if tmpAA <= k || tmpBA <= k {
      tmp[0] = arrA[i]
    } else {
      tmp[0] = -1
    }
    if tmpAB <= k || tmpBB <= k {
      tmp[1] = arrB[i]
    } else {
      tmp[1] = -1
    }

    if tmp[0] == -1 && tmp[1] == -1 {
      flag = false
      break
    }
  }

  if flag {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}
