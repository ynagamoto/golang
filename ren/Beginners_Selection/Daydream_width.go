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
	var residues []string
	for _, asset := range assets {
		// fmt.Printf("%s -> ", residue)
		if strings.HasPrefix(s, asset) {
			residues = append(residues, strings.Replace(s, asset, "", 1))
			// fmt.Printf("%s (%s)\n", residue, asset)
		}
		// fmt.Printf("%s (%s)\n", residue, asset)
	}

	for _, v := range residues {
		if len(v) == 0 { // search fin
			return true
		} else {
			if daydream(v) {
				return true
			}
		}
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
