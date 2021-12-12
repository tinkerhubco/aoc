package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	inputFile, _ := ioutil.ReadFile("input.txt")

	parsedInput := strings.Split(string(inputFile), "\n")

	var score int

	for _, line := range parsedInput {

		stack := list.New()

lineCheck:
		for _, char := range line {
			switch char {
			case '(', '[', '{', '<':
				stack.PushBack(char)
				break
			default:
				elem := stack.Back()
				stack.Remove(elem)

				lastChar := elem.Value.(rune)
				switch {
				case lastChar != '(' && char == ')':
					score += 3
					break lineCheck
				case lastChar != '[' && char == ']':
					score += 57
					break lineCheck
				case lastChar != '{' && char == '}':
					score += 1197
					break lineCheck
				case lastChar != '<' && char == '>':
					score += 25137
					break lineCheck
				}
			}
		}
	}

	fmt.Println(score)
}
