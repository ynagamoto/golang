package main

import (
	"fmt"
)

func main() {
	// input
	var a, b int
	fmt.Scan(&a, &b)

  // mul & check
	if a*b % 2 == 0 {
		fmt.Printf("Even\n")
	} else {
		fmt.Printf("Odd\n")
	}
}