package main

import (
	"fmt"
)

func getMax(d_arr *[]int) int {
	max := 0
	for _, v := range *d_arr {
		if max < v {
			max = v
		}
	}
	if max != 0 {
		for i, v := range *d_arr {
			if v == max { 
				(*d_arr)[i] = 0
			}
		}
	}
	return max
}

func main() {
	var n int
	fmt.Scan(&n)
	d_arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&d_arr[i])
	}

	ans := 0
	for ;; {
		if mochi := getMax(&d_arr); mochi == 0 {
			break
		}	
		ans++
	}

	fmt.Printf("%d\n", ans)
}