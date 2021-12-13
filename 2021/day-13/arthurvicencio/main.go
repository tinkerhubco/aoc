package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	inputFile, _ := ioutil.ReadFile("input.txt")

	parsedInput := strings.Split(string(inputFile), "\n\n")

	var xLen int

	grid := make([][]byte, 0)
	for _, line := range strings.Split(parsedInput[0], "\n") {
		point := strings.Split(line, ",")
		x, _ := strconv.Atoi(point[0])
		y, _ := strconv.Atoi(point[1])

		if y > len(grid) {
			for i := len(grid); i <= y; i++ {
				grid = append(grid, make([]byte, xLen))
			}
		}
		if x > xLen {
			for i := range grid {
				grid[i] = append(grid[i], make([]byte, (x-len(grid[i])+1))...)
			}
			xLen = x
		}
		grid[y][x] = '#'
	}

	var part1 int
	for _, line := range strings.Split(parsedInput[1], "\n") {
		var foldCommand string
		fmt.Sscanf(line, "fold along %s", &foldCommand)

		fold := strings.Split(foldCommand, "=")
		n, _ := strconv.Atoi(fold[1])

		switch fold[0] {
		case "y":
			grid = foldY(grid, n)
			break
		case "x":
			grid = foldX(grid, n)
			break
		}
		if part1 == 0 {
			part1 = countDots(grid)
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Println("Part 2:")

	for y := range grid {
		for x := range grid[y] {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Println()
	}
}

func foldY(grid [][]byte, y int) [][]byte {
	newGrid := make([][]byte, 0)

	for i := 0; i < y; i++ {
		row := make([]byte, len(grid[0]))
		for j := range grid[i] {
			row[j] = grid[i][j]
			if dy := y + (y - i); dy < len(grid) && grid[dy][j] == '#' {
				row[j] = '#'
			}
		}
		newGrid = append(newGrid, row)
	}
	return newGrid
}

func foldX(grid [][]byte, x int) [][]byte {
	newGrid := make([][]byte, 0)

	for i := range grid {
		row := make([]byte, x)
		for j := 0; j < x; j++ {
			row[j] = grid[i][j]
			if dx := x + (x - j); dx < len(grid[i]) && grid[i][dx] == '#' {
				row[j] = '#'
			}
		}
		newGrid = append(newGrid, row)
	}
	return newGrid
}

func countDots(grid [][]byte) int {
	var dots int

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '#' {
				dots++
			}
		}
	}
	return dots
}
