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

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			fmt.Print(arr_a[j][i])
			if j != h-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
