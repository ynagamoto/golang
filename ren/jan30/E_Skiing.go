package main

import (
	"fmt"
)

var lowest_h int = 100000000
var lowest_i int
var h_arr []int
var area [][]int

func search_depth(curr int, score int) int {
	better := -100000000
	if curr == lowest_i {
		return score // 一番下
	} else {
		for _, slope := range area {
			if slope[0] == curr {
				var result int
				if h_arr[slope[0]] >= h_arr[slope[1]] {
					result = search_depth(slope[1], score+h_arr[curr]-h_arr[slope[1]])
				} else {
					result = search_depth(slope[1], score-2*(h_arr[slope[1]]-h_arr[curr]))
				}
				if result > score {
					better = result
				}
			}
		}
		return better
	}
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	h_inp := make([]int, n)
	slopes := make([][]int, m)
	for i, _ := range h_inp {
		fmt.Scan(&h_inp[i])
		if h_inp[i] < lowest_h {
			lowest_h = h_inp[i]
			lowest_i = i
		}
	}
	h_arr = h_inp
	for i, _ := range slopes {
		var u, v int
		fmt.Scan(&u)
		fmt.Scan(&v)
		slopes[i] = []int{u - 1, v - 1}
	}
	area = slopes

	// とりあえず一番低いところまで深さ優先で
	score := search_depth(0, 0)
	if score < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(score)
	}
}
