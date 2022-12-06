package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %d\n", countCharactersBeforeUnique(input, 4))
	fmt.Printf("Part 2: %d\n", countCharactersBeforeUnique(input, 14))
}

func countCharactersBeforeUnique(str string, length int) int {
	for i := 0; i < len(str)-length; i++ {
		seen := make(map[byte]bool)
		for j := i; j < i+length; j++ {
			seen[str[j]] = true
		}
		if len(seen) == length {
			return i + length
		}
	}
	return -1
}
