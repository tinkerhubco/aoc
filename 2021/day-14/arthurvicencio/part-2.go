package main

import (
	_ "embed"
	"fmt"
	"strings"
)

type State struct {
	Pair string
	Step int
}

//go:embed input.txt
var input string

func main() {
	parsedInput := strings.Split(input, "\n\n")

	rules := make(map[string]string)
	for _, rule := range strings.Split(parsedInput[1], "\n") {
		var from, to string
		fmt.Sscanf(rule, "%s -> %s", &from, &to)
		rules[from] = to
	}

	state := parsedInput[0]

	charCount := simulate(state, rules)

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

func simulate(state string, rules map[string]string) map[byte]int {
	charCount := make(map[byte]int)
	for i := range state {
		charCount[state[i]]++
	}

	memo := make(map[State]map[byte]int)

	var run func(pair string, step int) map[byte]int
	run = func(pair string, step int) map[byte]int {
		chars := make(map[byte]int)

		if step >= 40 {
			return chars
		}

		if _, ok := memo[State{pair, step}]; ok {
			return memo[State{pair, step}]
		}

		if rule, ok := rules[pair]; ok {
			chars[rule[0]]++
			addToMap(chars, run(string(pair[0])+rule, step+1))
			addToMap(chars, run(rule+string(pair[1]), step+1))
		}

		memo[State{pair, step}] = chars

		return chars
	}

	for i := 0; i < len(state)-1; i++ {
		addToMap(charCount, run(state[i:i+2], 0))
	}

	return charCount
}

func addToMap(m, n map[byte]int) {
	for k, v := range n {
		m[k] += v
	}
}
