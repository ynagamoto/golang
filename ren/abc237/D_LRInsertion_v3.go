package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	ans_l := make([]int, n)
	count := 0
	for i, r := range s {
		if r == 'R' { // 左に付ける
			fmt.Print(i, " ")
		} else { // 右に付ける
			ans_l[count] = i
			count++
		}
	}
	fmt.Print(n)

	// output ans_l
	if count > 0 {
		fmt.Print(" ")
		for i := count - 1; i > 0; i-- {
			fmt.Print(ans_l[i], " ")
		}
		fmt.Println(ans_l[0])
	} else {
		fmt.Println()
	}
}
