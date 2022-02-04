package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)
	ans := 0
	for _, grid := range input {
		if grid== '1' {
			ans++
		}
	}
	fmt.Printf("%d\n", ans)
}