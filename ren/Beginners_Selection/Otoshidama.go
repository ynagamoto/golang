package main

import (
	"fmt"
)

func main() {
	var n, y int
	fmt.Scan(&n, &y)

	flag := false
outside:
	for i := 0; i <= n; i++ {
		if 10000*i > y {
			break
		}
		for j := 0; j <= n-i; j++ {
			if 5000*j > y {
				break
			}
			for k := 0; k <= n-i-j; k++ {
				if 1000*k > y {
					break
				}
				if i+j+k != n {
					continue
				}
				sum := 10000*i + 5000*j + 1000*k
				if sum == y {
					flag = true
					fmt.Printf("%d %d %d\n", i, j, k)
					break outside
				}
			}
		}
	}
	if !flag {
		fmt.Printf("%d %d %d\n", -1, -1, -1)
	}
}
