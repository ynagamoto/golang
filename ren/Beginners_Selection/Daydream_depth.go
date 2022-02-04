package main

import (
	"fmt"
	"strings"
)

var assets []string = []string{
	"dream",
	"dreamer",
	"erase",
	"eraser",
}

func daydream(s string) bool {
	for _, asset := range assets {
		// fmt.Printf("%s -> ", residue)
		if strings.HasPrefix(s, asset) {
			residue := strings.Replace(s, asset, "", 1)
			// fmt.Printf("%s (%s)\n", residue, asset)
			if len(residue) == 0 {
				return true
			}
			if daydream(residue) {
				return true
			}
		}
		// fmt.Printf("%s (%s)\n", residue, asset)
	}
	return false
}

func main() {

	var s string
	fmt.Scan(&s)

	if daydream(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
