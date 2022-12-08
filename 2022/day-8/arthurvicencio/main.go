package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	X int
	Y int
}

// up, down, left and right
var dir = []Point{
	Point{0, -1},
	Point{0, 1},
	Point{-1, 0},
	Point{1, 0},
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	grid := parseGrid(input)
	var visible int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if isVisible(x, y, grid) {
				visible++
			}
		}
	}
	return visible
}

func part2() int {
	grid := parseGrid(input)
	var max int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if isVisible(x, y, grid) {
				sceniceScore := getScenicScoresFrom(x, y, grid)
				if sceniceScore > max {
					max = sceniceScore
				}
			}
		}
	}
	return max
}

func parseGrid(input string) [][]int {
	grid := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]int, 0)
		for i := 0; i < len(line); i++ {
			n, _ := strconv.Atoi(string(line[i]))
			row = append(row, n)
		}
		grid = append(grid, row)
	}
	return grid
}

func isVisible(x, y int, grid [][]int) bool {
look:
	for _, d := range dir {
		var dx, dy = x + d.X, y + d.Y
		for isPointValid(dx, dy, grid) {
			if grid[y][x] <= grid[dy][dx] {
				continue look
			}
			dx, dy = dx+d.X, dy+d.Y
		}
		return true
	}
	return false
}

func getScenicScoresFrom(x, y int, grid [][]int) int {
	trees := 1
	for _, d := range dir {
		var dx, dy = x + d.X, y + d.Y
		var score int
		for isPointValid(dx, dy, grid) {
			score++
			if grid[y][x] <= grid[dy][dx] {
				break
			}
			dx, dy = dx+d.X, dy+d.Y
		}
		trees = trees * score
	}
	return trees
}

func isPointValid(x, y int, grid [][]int) bool {
	return y >= 0 && y <= len(grid)-1 && x >= 0 && x <= len(grid[y])-1
}
