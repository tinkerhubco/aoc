package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	inputFile, _ := ioutil.ReadFile("input.txt")

	parsedInput := strings.Split(string(inputFile), "\n")

	var horizontal int
	var depth int

	for _, instruction := range parsedInput {

		var command string
		var x int
		fmt.Sscanf(instruction, "%s %d", &command, &x)

		switch command {
		case "forward":
			horizontal += x
			break
		case "down":
			depth += x
			break
		case "up":
			depth -= x
			break
		}
	}

	fmt.Println(horizontal * depth)
}
