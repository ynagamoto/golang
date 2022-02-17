package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	// 反転させたものと比較
	s_runes := []rune(s)
	rev_runes := make([]rune, len(s))
	for i, char := range s_runes {
		rev_runes[len(s_runes)-1*(i+1)] = char
	}
	rev_s := string(rev_runes)
	s_trim := strings.Trim(s, "a")
	rev_trim := strings.Trim(rev_s, "a")
	s_trim_index := strings.Index(s, s_trim)
	rev_trim_index := strings.Index(rev_s, rev_trim)

	// fmt.Printf("s: %s, rev_s: %s\n", s, rev_s)
	// fmt.Printf("s_trim: %s, rev_trim: %s\n", s_trim, rev_trim)
	// fmt.Printf("s: %d, rev_s: %d\n", s_trim_index, rev_trim_index)

	if s_trim_index <= rev_trim_index {
		if s_trim == rev_trim {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else {
		fmt.Println("No")
	}
}
