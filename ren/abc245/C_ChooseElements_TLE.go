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
  arrX := make([][]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&arrA[i])
  }
  for i := 0; i < n; i++ {
    fmt.Scan(&arrB[i])
    tmp := make([]int, 2)
    tmp[0] = arrA[i]
    tmp[1] = arrB[i]
    arrX[i] = tmp
  }

  flag := true
  for i := 0; i < n-1; i++ {
    var tmpAA, tmpAB int = k+1, k+1
    var tmpBA, tmpBB int = k+1, k+1
    flagA := false
    flagB := false

    // A
    if arrX[i][0] != -1 {
      tmpAA = int(math.Abs(float64(arrA[i]-arrA[i+1])))
      tmpAB = int(math.Abs(float64(arrA[i]-arrB[i+1])))
    }
    // B
    if arrX[i][1] != -1 {
      tmpBA = int(math.Abs(float64(arrB[i]-arrA[i+1])))
      tmpBB = int(math.Abs(float64(arrB[i]-arrB[i+1])))
    }

    if !(tmpAA <= k || tmpBA <= k) {
      arrX[i+1][0] = -1
      flagA = true
    }
    if !(tmpAB <= k || tmpBB <= k) {
      arrX[i+1][1] = -1
      flagB = true
    }

    if flagA && flagB {
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
