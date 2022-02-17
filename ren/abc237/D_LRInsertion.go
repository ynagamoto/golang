package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	s_runes := []rune(s)
	ans := []rune{rune(0)}
	for i, curr := 0, 0; i < len(s); i++ {
		char := s_runes[i]
		if char == 'R' {
			curr += 1
		}

		ans_l := make([]rune, len(ans[:curr]))
		ans_r := make([]rune, len(ans[curr:]))
		copy(ans_l, ans[:curr])
		copy(ans_r, ans[curr:])
		// fmt.Print(ans, " -> ")
		tmp := append(ans_l, rune(i+1))
		// fmt.Print(tmp, " + ", ans_r)
		ans = append(tmp, ans_r...)
		// fmt.Println(" = ", ans)
	}

	for i, char := range ans {
		fmt.Print(char)
		if i != len(ans)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
