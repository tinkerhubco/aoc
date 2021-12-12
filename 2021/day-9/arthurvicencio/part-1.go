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

// up, right, down and left
var dir = []Point{
	{Y: -1, X: 0},
	{Y: 0, X: 1},
	{Y: 1, X: 0},
	{Y: 0, X: -1},
}

func main() {

	inputFile, _ := ioutil.ReadFile("input.txt")

	parsedInput := strings.Split(string(inputFile), "\n")

	grid := make([][]int, 0)
	for _, rowInput := range parsedInput {
		row := make([]int, 0)
		line := strings.Split(rowInput, "")
		for _, cell := range line {
			n, _ := strconv.Atoi(cell)
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	var sum int
	for y := range grid {
		for x, height := range grid[y] {
			if isLowPoint(grid, y, x) {
				sum += height + 1
			}
		}
	}

	fmt.Println(sum)
}

func isLowPoint(grid [][]int, y, x int) bool {
	for _, d := range dir {
		isValidPoint := d.Y+y >= 0 &&
			d.X+x < len(grid[y]) &&
			d.Y+y < len(grid) &&
			d.X+x >= 0
		if !isValidPoint {
			continue
		}
		if grid[y][x] >= grid[y+d.Y][x+d.X] {
			return false
		}
	}
	return true
}
