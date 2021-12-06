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

	current, _ := strconv.Atoi(parsedInput[0])

	for _, val := range parsedInput[1:] {

		n, _ := strconv.Atoi(val)
		if n > current {
			increase++
		}

		current = n
	}

	fmt.Println(increase)
}
