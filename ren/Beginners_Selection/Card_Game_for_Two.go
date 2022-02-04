package main

import (
	"fmt"
)

func getMax(a_arr *[]int) int {
	max := 0
	max_p := 0
	for i, v := range *a_arr {
		if max < v {
			max = v
			max_p = i
		}
	}
	(*a_arr)[max_p] = 0
	return max
}

func main() {
	var n int
	fmt.Scan(&n)
	a_arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a_arr[i])
	}

	alice := 0
	bob := 0
	for i := 0; i < n; i++ {
		// alice
		alice += getMax(&a_arr)
		if i++; i >= n {
			break
		}
		// bob
		bob += getMax(&a_arr)
	}
	fmt.Printf("%d\n", alice-bob)
}