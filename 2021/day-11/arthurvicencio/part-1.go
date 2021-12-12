package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Point struct {
	Y int
	X int
}

// all directions
var dirs = []Point{
	{Y: -1, X: -1},
	{Y: -1, X: 0},
	{Y: -1, X: 1},
	{Y: 0, X: -1},
	{Y: 0, X: 1},
	{Y: 1, X: -1},
	{Y: 1, X: 0},
	{Y: 1, X: 1},
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

	steps := 100

	var flashes int
	for i := 0; i < steps; i++ {
		flashed := make(map[Point]bool)
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				simulate(grid, flashed, y, x)
			}
		}
		flashes += len(flashed)
	}

	fmt.Println(flashes)
}

func simulate(grid [][]int, flashed map[Point]bool, y, x int) {
	queue := list.New()
	queue.PushBack(Point{Y: y, X: x})

	for queue.Len() > 0 {

		elem := queue.Front()
		queue.Remove(elem)

		point := elem.Value.(Point)

		if flashed[point] {
			continue
		}

		grid[point.Y][point.X]++
		if grid[point.Y][point.X] > 9 {
			grid[point.Y][point.X] = 0
			flashed[point] = true
			adjacentPoints := getAdjacentPoints(grid, point.Y, point.X)
			for _, point := range adjacentPoints {
				queue.PushBack(point)
			}
		}
	}
}

func getAdjacentPoints(grid [][]int, y, x int) []Point {
	points := make([]Point, 0)
	for _, d := range dirs {
		isValidPoint := d.Y+y >= 0 &&
			d.X+x < len(grid[y]) &&
			d.Y+y < len(grid) &&
			d.X+x >= 0
		if isValidPoint {
			points = append(points, Point{Y: d.Y + y, X: d.X + x})
		}
	}
	return points
}
