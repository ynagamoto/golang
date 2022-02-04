package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	t := make([]int, n)
	sche := make([][]int, n)
	for i := 0; i < n; i++ {
		hoge := make([]int, 2)
		fmt.Scan(&t[i])
		fmt.Scan(&hoge[0], &hoge[1])
		sche[i] = []int{hoge[0], hoge[1]}
	}

	flag := false
	now := 0
	current := []int{0, 0}
	for i := 0; i < n; i++ {
		flag = false
		x_diff := int(math.Abs(float64(sche[i][0]) - float64(current[0])))
		y_diff := int(math.Abs(float64(sche[i][1]) - float64(current[1])))
		diff := x_diff + y_diff
		limit := t[i] - now
		// fmt.Printf("limit -> %d, diff -> %d\n", limit, diff)
		if limit > diff {
			if (limit-diff)%2 == 0 {

				flag = true
			}
		} else if limit == diff {
			flag = true
		} else {
			break
		}
		now = t[i]
		current[0] = sche[i][0]
		current[1] = sche[i][1]
	}

	if flag {
		fmt.Print("Yes\n")
	} else {
		fmt.Print("No\n")
	}
}
