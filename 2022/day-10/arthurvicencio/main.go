package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: \n%s", part2())
}

func part1() int {
	cycles := parseCommands(input)
	var result int
	for _, index := range []int{20, 60, 100, 140, 180, 220} {
		result += cycles[index-1] * index
	}
	return result
}

func part2() string {
	cycles := parseCommands(input)
	var result string
	for i := 0; i < 240; i++ {
		result += draw(i, cycles[i])
	}
	return result
}

func parseCommands(input string) []int {
	var cycles []int
	x := 1
	for _, line := range strings.Split(input, "\n") {
		var cmd string
		var amount int

		fmt.Sscanf(line, "%s %d", &cmd, &amount)

		cycles = append(cycles, x)
		if cmd == "addx" {
			cycles = append(cycles, x)
			x += amount
		}
	}
	return cycles
}

func draw(cycle int, x int) string {
	var result string
	rowCycle := cycle % 40
	if rowCycle <= x+1 && rowCycle >= x-1 {
		result = "#"
	} else {
		result = "."
	}
	if rowCycle == 39 {
		result += "\n"
	}
	return result
}
