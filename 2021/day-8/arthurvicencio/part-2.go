package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var segments = []string{
	"abcefg",
	"cf",
	"acdeg",
	"acdfg",
	"bcdf",
	"abdfg",
	"abdefg",
	"acf",
	"abcdefg",
	"abcdfg",
}

func main() {

	inputFile, _ := ioutil.ReadFile("input.txt")

	parsedInput := strings.Split(string(inputFile), "\n")

	/*
		Deductions for left side of input
		if a letter occur 6 times it maps to "b"
		if a letter occur 4 times it maps to "e"
		if a letter occur 9 times it maps to "f"

		if the length of the segment is 2 the unknown letter is "c"
		if the length of the segment is 4 the unknown letter is "d"
		if the length of the segment is 3 the unknown letter is "a"
		the last unkown letter is "g"
	*/

	var sum int

	for _, v := range parsedInput {
		line := strings.Split(v, " | ")

		left := strings.Split(line[0], " ")

		charMap := make(map[rune]int)
		for _, segment := range left {
			for _, char := range segment {
				charMap[char]++
			}
		}

		unknownChars := make(map[rune]bool)
		decode := make(map[rune]rune)
		for char, count := range charMap {
			switch count {
			case 6:
				decode[char] = 'b'
				break
			case 4:
				decode[char] = 'e'
				break
			case 9:
				decode[char] = 'f'
				break
			default:
				unknownChars[char] = true
				break
			}
		}

		for i := 0; i < 2; i++ {
			for _, segment := range left {
				var decodedChar rune
				switch len(segment) {
				case 2:
					decodedChar = 'c'
					break
				case 4:
					decodedChar = 'd'
					break
				case 3:
					decodedChar = 'a'
					break
				}

				if decodedChar != 0 {
					unknown := getNotIn(decode, segment)
					if len(unknown) == 1 {
						decode[unknown[0]] = decodedChar
						delete(unknownChars, unknown[0])
					}
				}
			}
		}
		for char := range unknownChars {
			decode[char] = 'g'
		}

		var decodedValue string
		right := strings.Split(line[1], " ")
		for _, segment := range right {
			digit := getSegmentNumber(segment, decode)
			decodedValue += strconv.Itoa(digit)
		}

		n, _ := strconv.Atoi(decodedValue)
		sum += n
	}

	fmt.Println(sum)
}

func getNotIn(m map[rune]rune, s string) []rune {
	unknown := make([]rune, 0)
	for _, char := range s {
		if _, ok := m[char]; !ok {
			unknown = append(unknown, char)
		}
	}
	return unknown
}

func getSegmentNumber(str string, decode map[rune]rune) int {
	for i, segment := range segments {
		if len(segment) != len(str) {
			continue
		}
		inSegment := 0
		for _, checkChar := range segment {
			for _, char := range str {
				if decode[char] == checkChar {
					inSegment++
					break
				}
			}
		}
		if inSegment == len(segment) {
			return i
		}
	}
	return -1
}
