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

	charMaps := make([]map[byte]int, 0)
	for i := 0; i < len(parsedInput[0]); i++ {
		charMap := make(map[byte]int)
		for _, bin := range parsedInput {
			charMap[bin[i]]++
		}
		charMaps = append(charMaps, charMap)
	}

	var gamma string
	var epsilon string
	for i := 0; i < len(charMaps); i++ {
		if charMaps[i]['1'] >= charMaps[i]['0'] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaValue, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonValue, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(gammaValue * epsilonValue)
}
