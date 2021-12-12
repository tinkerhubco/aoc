package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Point struct {
	Y int
	X int
}

type Card struct {
	NumbersCalled map[Point]bool
	Grid          [][]int
}

func main() {
	inputFile, _ := ioutil.ReadFile("input.txt")

	parsedInput := strings.Split(string(inputFile), "\n\n")

	numbersCalled := make([]int, 0)
	for _, number := range strings.Split(parsedInput[0], ",") {
		n, _ := strconv.Atoi(number)
		numbersCalled = append(numbersCalled, n)
	}

	cards := make([]*Card, 0)
	for _, cardInput := range parsedInput[1:] {
		card := &Card{
			Grid:          make([][]int, 0),
			NumbersCalled: make(map[Point]bool),
		}
		for _, row := range strings.Split(cardInput, "\n") {
			gridRow := make([]int, 0)
			for _, cell := range strings.Split(row, " ") {
				if cell == "" {
					continue
				}
				n, _ := strconv.Atoi(cell)
				gridRow = append(gridRow, n)
			}
			card.Grid = append(card.Grid, gridRow)
		}
		cards = append(cards, card)
	}

	var winningCard Card
	var lastNumberCalled int
winner:
	for _, number := range numbersCalled {
		for _, card := range cards {

			markNumber(*card, number)
			if isWinningCard(*card) {
				winningCard = *card
				lastNumberCalled = number
				break winner
			}
		}
	}

	score := cardScore(winningCard)
	fmt.Println(score * lastNumberCalled)
}

func markNumber(card Card, number int) {
mark:
	for y := 0; y < len(card.Grid); y++ {
		for x := 0; x < len(card.Grid[0]); x++ {
			if card.Grid[y][x] == number {
				card.NumbersCalled[Point{Y: y, X: x}] = true
				break mark
			}
		}
	}
}

func isWinningCard(card Card) bool {
	for y := 0; y < len(card.Grid); y++ {
		rowCount := 0
		for x := 0; x < len(card.Grid[0]); x++ {
			if card.NumbersCalled[Point{Y: y, X: x}] {
				rowCount++
			}
		}
		if rowCount == len(card.Grid[0]) {
			return true
		}
	}

	for x := 0; x < len(card.Grid[0]); x++ {
		colCount := 0
		for y := 0; y < len(card.Grid); y++ {
			if card.NumbersCalled[Point{Y: y, X: x}] {
				colCount++
			}
		}
		if colCount == len(card.Grid) {
			return true
		}
	}

	return false
}

func cardScore(card Card) int {
	sum := 0
	for y := 0; y < len(card.Grid); y++ {
		for x := 0; x < len(card.Grid[0]); x++ {
			if !card.NumbersCalled[Point{Y: y, X: x}] {
				sum += card.Grid[y][x]
			}
		}
	}
	return sum
}
