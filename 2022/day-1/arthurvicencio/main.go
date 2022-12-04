package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	groups := make([][]string, 0)
	for _, line := range strings.Split(input, "\n\n") {
		groups = append(groups, strings.Split(line, "\n"))
	}

	var max int
	for _, grp := range groups {
		var score int
		for _, val := range grp {
			n, _ := strconv.Atoi(val)
			score += n
		}
		if score > max {
			max = score
		}
	}
	return max
}

func part2() int {
	groups := make([][]string, 0)
	for _, line := range strings.Split(input, "\n\n") {
		groups = append(groups, strings.Split(line, "\n"))
	}

	var maxGroups []int
	for _, grp := range groups {
		var score int
		for _, val := range grp {
			n, _ := strconv.Atoi(val)
			score += n
		}
		maxGroups = append(maxGroups, score)
	}
	sort.Slice(maxGroups, func(i, j int) bool { return maxGroups[i] > maxGroups[j] })
	return maxGroups[0] + maxGroups[1] + maxGroups[2]
}
