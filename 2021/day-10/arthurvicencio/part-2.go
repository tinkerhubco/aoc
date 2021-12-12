package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {

	inputFile, _ := ioutil.ReadFile("input.txt")

	parsedInput := strings.Split(string(inputFile), "\n")

	scores := make([]int, 0)

lineCheck:
	for _, line := range parsedInput {

		stack := list.New()
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
					continue lineCheck
				case lastChar != '[' && char == ']':
					continue lineCheck
				case lastChar != '{' && char == '}':
					continue lineCheck
				case lastChar != '<' && char == '>':
					continue lineCheck
				}
			}
		}

		var score int
		for stack.Len() > 0 {
			elem := stack.Back()
			stack.Remove(elem)

			lastChar := elem.Value.(rune)
			switch lastChar {
			case '(':
				score = (score * 5) + 1
				break
			case '[':
				score = (score * 5) + 2
				break
			case '{':
				score = (score * 5) + 3
				break
			case '<':
				score = (score * 5) + 4
				break
			}
		}

		scores = append(scores, score)
	}

	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
	fmt.Println(scores[len(scores)/2])
}
