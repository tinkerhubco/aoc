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
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	var count int
	for _, line := range strings.Split(input, "\n") {
		var fmin, fmax, smin, smax int
		fmt.Sscanf(line, "%d-%d,%d-%d", &fmin, &fmax, &smin, &smax)

		if inRange(fmin, fmax, smin, smax) || inRange(smin, smax, fmin, fmax) {
			count++
		}
	}
	return count
}

func part2() int {
	var count int
	for _, line := range strings.Split(input, "\n") {
		var fmin, fmax, smin, smax int
		fmt.Sscanf(line, "%d-%d,%d-%d", &fmin, &fmax, &smin, &smax)

		if overlap(fmin, fmax, smin, smax) {
			count++
		}
	}
	return count
}

func inRange(fmin, fmax, smin, smax int) bool {
	return fmin <= smin && fmax >= smax
}

func overlap(fmin, fmax, smin, smax int) bool {
	return fmin <= smax && fmax >= smin
}
