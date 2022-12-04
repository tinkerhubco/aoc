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
	characterMap := generateCharaterMap()

	var score int
	for _, line := range strings.Split(input, "\n") {

		first := line[:len(line)/2]
		second := line[len(line)/2:]

		firstMap := generateMapFromString(first)
		secondMap := generateMapFromString(second)

		for char := range firstMap {
			if secondMap[char] {
				score += characterMap[char]
			}
		}
	}
	return score
}

func part2() int {
	characterMap := generateCharaterMap()

	rawInput := strings.Split(input, "\n")
	groups := make([][]string, 0)
	for i := 0; i < len(rawInput); i += 3 {
		groups = append(groups, []string{rawInput[i], rawInput[i+1], rawInput[i+2]})
	}

	var score int
	for _, group := range groups {

		fisrt := generateMapFromString(group[0])
		second := generateMapFromString(group[1])
		third := generateMapFromString(group[2])

		for char := range inMap(fisrt, second, third) {
			score += characterMap[char]
		}
	}
	return score
}

func generateCharaterMap() map[rune]int {
	charMap := make(map[rune]int)
	val := 1
	char := 'a'
	for char <= 'z' {
		charMap[char] = val
		val++
		char++
	}
	char = 'A'
	for char <= 'Z' {
		charMap[char] = val
		val++
		char++
	}
	return charMap
}

func generateMapFromString(s string) map[rune]bool {
	m := make(map[rune]bool)
	for _, ch := range s {
		m[ch] = true
	}
	return m
}

func inMap(m map[rune]bool, c ...map[rune]bool) map[rune]bool {
	r := make(map[rune]bool)
	for v := range m {
		var found int
		for _, cm := range c {
			if cm[v] {
				found++
			}
		}
		if found == len(c) {
			r[v] = true
		}
	}
	return r
}
