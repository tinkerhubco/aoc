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

	var increase int

	first, _ := strconv.Atoi(parsedInput[0])
	second, _ := strconv.Atoi(parsedInput[1])
	third, _ := strconv.Atoi(parsedInput[2])

	currentSum := first + second + third

	for i := 1; i < len(parsedInput)-2; i++ {
		first, _ = strconv.Atoi(parsedInput[i])
		second, _ := strconv.Atoi(parsedInput[i+1])
		third, _ = strconv.Atoi(parsedInput[i+2])

		sum := first + second + third
		if sum > currentSum {
			increase++
		}

		currentSum = sum
	}

	fmt.Println(increase)
}
