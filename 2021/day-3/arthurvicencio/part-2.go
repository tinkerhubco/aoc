package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	inputFile, _ := ioutil.ReadFile("input.txt")

	parsedInput := strings.Split(string(inputFile), "\n")

	oxygenSet := make(map[string]bool)
	co2Set := make(map[string]bool)
	for _, bin := range parsedInput {
		oxygenSet[bin] = true
		co2Set[bin] = true
	}

	for i := 0; len(oxygenSet) > 1; i++ {
		most := findMostCommonFromPosition(oxygenSet, i)
		removeNotContainsFromPosition(oxygenSet, i, most)
	}

	for i := 0; len(co2Set) > 1; i++ {
		most := findMostCommonFromPosition(co2Set, i)
		least := byte('1')
		if most == '1' {
			least = '0'
		}
		removeNotContainsFromPosition(co2Set, i, least)
	}

	var oxygen string
	for bin := range oxygenSet {
		oxygen = bin
	}

	var co2 string
	for bin := range co2Set {
		co2 = bin
	}

	oxygenValue, _ := strconv.ParseInt(oxygen, 2, 64)
	co2Value, _ := strconv.ParseInt(co2, 2, 64)

	fmt.Println(oxygenValue * co2Value)
}

func findMostCommonFromPosition(set map[string]bool, pos int) byte {
	charMap := make(map[byte]int)
	for bin := range set {
		charMap[bin[pos]]++
	}

	if charMap['0'] > charMap['1'] {
		return '0'
	}
	return '1'
}

func removeNotContainsFromPosition(set map[string]bool, pos int, find byte) {
	for key := range set {
		if key[pos] != find {
			delete(set, key)
		}
	}
}
