package main

import (
	"fmt"
)

func main() {
	var h, w int
	fmt.Scan(&h)
	fmt.Scan(&w)

	arr_a := make([][]int, h)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			var tmp int
			fmt.Scan(&tmp)
			arr_a[i] = append(arr_a[i], tmp)
		}
	}

	arr_b := make([][]int, w)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			arr_b[i] = append(arr_b[i], arr_a[j][i])
		}
	}

	for i := 0; i < len(arr_b); i++ {
		for j := 0; j < len(arr_b[0]); j++ {
			fmt.Print(arr_b[i][j], " ")
		}
		fmt.Printf("\b\b\n")
	}
}
