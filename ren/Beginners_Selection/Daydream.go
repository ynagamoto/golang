package main

import (
	"fmt"
	"strings"
)

var assets []string = []string{
	"eraser",
	"erase",
	"dreamer",
	"dream",
}

func main() {
	var s string
	fmt.Scan(&s)

	for _, v := range assets {
		s = strings.ReplaceAll(s, v, "")
	}

	if len(s) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
