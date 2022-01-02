package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	Y int
	X int
}

var dirs = []Point{
	{Y: -1, X: 0},
	{Y: 0, X: 1},
	{Y: 1, X: 0},
	{Y: 0, X: -1},
}

//go:embed input.txt
var input string

// Implementation of dijkstra’s algorithm
func main() {

	rawInput := strings.Split(input, "\n")

	grid := buildGrid(rawInput)

	h := &StateHeap{}
	heap.Init(h)
	heap.Push(h, State{Point{0, 0}, 0})

	best := make(map[Point]int)

	for h.Len() > 0 {

		current := heap.Remove(h, 0).(State)

		if current.Position == (Point{len(grid) - 1, len(grid[0]) - 1}) {
			fmt.Println(current.Cost)
			return
		}

		for _, next := range getAdjacentPoints(current.Position, grid) {

			nextCost := best[current.Position] + grid[next.Y][next.X]

			prevCost, exists := best[next]
			if !exists || prevCost > nextCost {
				best[next] = nextCost
				heap.Push(h, State{next, nextCost})
			}
		}
	}
}

func buildGrid(rawInput []string) [][]int {
	grid := make([][]int, 0)
	for _, line := range rawInput {
		row := make([]int, 0)
		for _, cell := range line {
			n, _ := strconv.Atoi(string(cell))
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	original := make([][]int, len(grid))
	for i := range grid {
		original[i] = make([]int, len(grid[i]))
		copy(original[i], grid[i])
	}

	for n := 0; n < 4; n++ {
		for i := range original {
			for j := range original[i] {
				original[i][j]++
				if original[i][j] > 9 {
					original[i][j] = 1
				}
				grid[i] = append(grid[i], original[i][j])
			}
		}
	}

	originalRow := make([][]int, len(grid))
	for i := range grid {
		originalRow[i] = make([]int, len(grid[i]))
		copy(originalRow[i], grid[i])
	}

	for n := 0; n < 4; n++ {
		for i := range originalRow {
			newRow := make([]int, 0)
			for j := range originalRow[i] {
				originalRow[i][j]++
				if originalRow[i][j] > 9 {
					originalRow[i][j] = 1
				}
				newRow = append(newRow, originalRow[i][j])
			}
			grid = append(grid, newRow)
		}
	}

	return grid
}

func getAdjacentPoints(p Point, grid [][]int) []Point {
	points := make([]Point, 0)
	for _, d := range dirs {
		nextPoint := Point{
			Y: p.Y + d.Y,
			X: p.X + d.X,
		}
		isValid := nextPoint.Y < len(grid) &&
			nextPoint.Y >= 0 &&
			nextPoint.X < len(grid[0]) &&
			nextPoint.X >= 0
		if isValid {
			points = append(points, nextPoint)
		}
	}
	return points
}

type State struct {
	Position Point
	Cost     int
}

type StateHeap []State

func (h StateHeap) Len() int           { return len(h) }
func (h StateHeap) Less(i, j int) bool { return h[i].Cost < h[j].Cost }
func (h StateHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StateHeap) Push(x interface{}) {
	*h = append(*h, x.(State))
}

func (h *StateHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
