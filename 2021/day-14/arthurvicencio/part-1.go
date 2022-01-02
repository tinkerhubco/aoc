package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	parsedInput := strings.Split(input, "\n\n")

	state := parsedInput[0]

	rules := make(map[string]string)
	for _, rule := range strings.Split(parsedInput[1], "\n") {
		var from string
		var to string
		fmt.Sscanf(rule, "%s -> %s", &from, &to)
		rules[from] = to
	}

	steps := 10

	for i := 0; i < steps; i++ {
		for j := 0; j < len(state)-1; j++ {
			if rules[state[j:j+2]] != "" {
				state = state[:j+1] + rules[state[j:j+2]] + state[j+1:]
				j++
			}

		}
	}

	charCount := make(map[rune]int)
	for _, char := range state {
		charCount[char]++
	}

	max, min := 0, int(1e18)
	for _, count := range charCount {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}

	fmt.Println(max - min)
}
