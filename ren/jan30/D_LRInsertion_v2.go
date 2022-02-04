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
	ans := []rune{rune(n)}
	for i := n - 1; i >= 0; i-- {
		char := s_runes[i]
		if char == 'R' { // 左に付ける
			ans = append([]rune{rune(i)}, ans...)
		} else { // 右に付ける
			ans = append(ans, rune(i))
		}
	}

	// output
	for i, char := range ans {
		fmt.Print(char)
		if i != len(ans)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
