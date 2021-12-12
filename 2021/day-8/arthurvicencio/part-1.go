package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	inputFile, _ := ioutil.ReadFile("input.txt")

	parsedInput := strings.Split(string(inputFile), "\n")

	var sum int

	for _, v := range parsedInput {
		line := strings.Split(v, " | ")

		for _, digit := range strings.Split(line[1], " ") {
			switch len(digit) {
			case 2, 4, 3, 7:
				sum++
				break
			}
		}
	}

	fmt.Println(sum)
}
