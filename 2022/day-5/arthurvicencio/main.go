package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %s\n", part1())
	fmt.Printf("Part 2: %s\n", part2())
}

func part1() string {
	rawInput := strings.Split(input, "\n\n")

	stacks := createStacks(strings.Split(rawInput[0], "\n"))

	for _, cmd := range strings.Split(rawInput[1], "\n") {
		var amount, from, to int
		fmt.Sscanf(cmd, "move %d from %d to %d", &amount, &from, &to)

		from, to = from-1, to-1

		stacks[to] = append(stacks[to], reverseRuneSlice(stacks[from][len(stacks[from])-amount:])...)
		stacks[from] = stacks[from][:len(stacks[from])-amount]

	}
	return getTopsOfStackToString(stacks)
}

func part2() string {
	rawInput := strings.Split(input, "\n\n")

	stacks := createStacks(strings.Split(rawInput[0], "\n"))

	for _, cmd := range strings.Split(rawInput[1], "\n") {
		var amount, from, to int
		fmt.Sscanf(cmd, "move %d from %d to %d", &amount, &from, &to)

		from, to = from-1, to-1

		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-amount:]...)
		stacks[from] = stacks[from][:len(stacks[from])-amount]

	}
	return getTopsOfStackToString(stacks)
}

func createStacks(inputStacks []string) [][]rune {
	stacks := make([][]rune, (len(inputStacks[0])/4)+1)
	for _, elms := range inputStacks {
		var stackIndex int
		for i := 0; i < len(elms); i += 4 {
			if elms[i] == ' ' {
				stackIndex++
				continue
			}
			stacks[stackIndex] = append([]rune{rune(elms[i+1])}, stacks[stackIndex]...)
			stackIndex++
		}

	}
	return stacks
}

func getTopsOfStackToString(stacks [][]rune) string {
	var str string
	for _, stack := range stacks {
		str += string(stack[len(stack)-1])
	}
	return str
}

func reverseRuneSlice(r []rune) []rune {
	result := make([]rune, 0)
	for i := len(r) - 1; i >= 0; i-- {
		result = append(result, r[i])
	}
	return result
}
