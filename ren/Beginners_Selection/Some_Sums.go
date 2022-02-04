package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	ans := 0
	for i := 1; i <= n; i++ {
		num_str := strconv.Itoa(i)
		sum := 0
		for _, char := range num_str {
			num, _ := strconv.Atoi(string(char))
			sum += num
		}
		if sum >= a && sum <= b {
			ans += i
		}
	}
	fmt.Printf("%d\n", ans)
}