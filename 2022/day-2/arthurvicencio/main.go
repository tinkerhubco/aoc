package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const (
	LOSE = 0
	DRAW = 3
	WIN  = 6
)

var item = map[string]int{
	"ROCK":     1,
	"PAPER":    2,
	"SCISSORS": 3,
}

var itemStats = map[string]map[string]int{
	"ROCK": map[string]int{
		"ROCK":     DRAW,
		"PAPER":    LOSE,
		"SCISSORS": WIN,
	},
	"PAPER": map[string]int{
		"ROCK":     WIN,
		"PAPER":    DRAW,
		"SCISSORS": LOSE,
	},
	"SCISSORS": map[string]int{
		"ROCK":     LOSE,
		"PAPER":    WIN,
		"SCISSORS": DRAW,
	},
}

var leftMap = map[string]string{
	"A": "ROCK",
	"B": "PAPER",
	"C": "SCISSORS",
}

var rightMap = map[string]string{
	"X": "ROCK",
	"Y": "PAPER",
	"Z": "SCISSORS",
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	var outcome int
	for _, line := range strings.Split(input, "\n") {
		var first, second string
		fmt.Sscanf(line, "%s %s", &first, &second)
		outcome += play(first, second)
	}
	return outcome
}

func part2() int {
	var outcome int
	for _, line := range strings.Split(input, "\n") {
		var first, second string
		fmt.Sscanf(line, "%s %s", &first, &second)
		outcome += playForPart2(first, second)
	}
	return outcome
}

func play(first, second string) int {
	left := leftMap[first]
	right := rightMap[second]
	return item[right] + itemStats[right][left]
}

func playForPart2(first, second string) int {
	left := leftMap[first]
	right := getDesiredItem(first, second)
	return item[right] + itemStats[right][left]
}

func getDesiredItem(first, second string) string {
	left := leftMap[first]
	switch second {
	case "X":
		return getItemWithResultVs(WIN, itemStats[left])
	case "Y":
		return getItemWithResultVs(DRAW, itemStats[left])
	case "Z":
		return getItemWithResultVs(LOSE, itemStats[left])
	}
	return ""
}

func getItemWithResultVs(result int, m map[string]int) string {
	for item, r := range m {
		if r == result {
			return item
		}
	}
	return ""
}
