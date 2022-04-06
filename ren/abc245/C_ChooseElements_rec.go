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
  arrX := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&arrA[i])
  }
  for i := 0; i < n; i++ {
    fmt.Scan(&arrB[i])
  }

  flag := choose(n-1, n, k, arrA, arrB, arrX)

  if flag {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}

func choose(i int, n int, k int, arrA []int, arrB []int, arrX []int) bool {
  if i == -1 {
    return true
  }

  var flagA, flagB bool
  if i == n-1 {
    arrX[i] = arrA[i]
    flagA = choose(i-1, n, k, arrA, arrB, arrX)
    arrX[i] = arrB[i]
    flagB = choose(i-1, n, k, arrA, arrB, arrX)
  } else {
    before := arrX[i+1]
    tmpA := int(math.Abs(float64(arrA[i]-before)))
    tmpB := int(math.Abs(float64(arrB[i]-before)))

    if tmpA <= k {
      arrX[i] = arrA[i]
      flagA = choose(i-1, n, k, arrA, arrB, arrX)
    }
    if tmpB <= k {
      arrX[i] = arrB[i]
      flagB = choose(i-1, n, k, arrA, arrB, arrX)
    }
  }

  if !(flagA || flagB) {
    return false
  } else {
    return true
  }
}
