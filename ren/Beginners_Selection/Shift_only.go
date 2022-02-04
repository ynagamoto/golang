package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	input := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
	}
	
	count := 0
	outside: for ;; {
		for i := 0; i < n; i++ {
			if input[i]%2 == 1 {
				break outside
			} else {
				input[i] = input[i]/2
			}
		}
		count++
	}

	fmt.Printf("%d\n", count)
}