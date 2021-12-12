package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"sort"
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

	basins := make([]int, 0)
	seen := make(map[Point]bool)
	for y := range grid {
		for x := range grid[y] {

			var size int

			queue := list.New()
			queue.PushBack(Point{Y: y, X: x})

			for queue.Len() > 0 {

				elem := queue.Front()
				queue.Remove(elem)

				point := elem.Value.(Point)

				if grid[point.Y][point.X] == 9 {
					continue
				}

				if seen[point] {
					continue
				}
				seen[point] = true

				size++

				for _, d := range dir {
					isValidPoint := d.Y+point.Y >= 0 &&
						d.X+point.X < len(grid[y]) &&
						d.Y+point.Y < len(grid) &&
						d.X+point.X >= 0
					if isValidPoint {
						queue.PushBack(Point{Y: d.Y + point.Y, X: d.X + point.X})
					}
				}
			}

			basins = append(basins, size)
		}
	}

	sort.Slice(basins, func(i, j int) bool { return basins[i] > basins[j] })

	fmt.Println(basins[0] * basins[1] * basins[2])
}
