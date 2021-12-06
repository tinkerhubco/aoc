package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Coord struct {
	Y int
	X int
}

func main() {

	inputFile, _ := ioutil.ReadFile("input.txt")

	raw := strings.Split(string(inputFile), "\n")

	seen := make(map[Coord]int)
	intersect := 0
	for _, v := range raw {
		var x1, y1, x2, y2 int

		fmt.Sscanf(v, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		isVertical := y1 == y2
		isHorizontal := x1 == x2

		switch {
		case isVertical:
			start := x1
			end := x2

			if x1 > x2 {
				start = x2
				end = x1
			}

			for ; start <= end; start++ {
				coord := Coord{
					Y: y1,
					X: start,
				}
				if seen[coord] == 1 {
					intersect++
				}
				seen[coord]++
			}
			break
		case isHorizontal:
			start := y1
			end := y2

			if y1 > y2 {
				start = y2
				end = y1
			}
			for ; start <= end; start++ {
				coord := Coord{
					Y: start,
					X: x1,
				}

				seen[coord]++
				if seen[coord] == 2 {
					intersect++
				}
			}
			break
		default:
			startY := y1
			endY := y2

			startX := x1
			endX := x2
			if y1 > y2 {
				startY = y2
				endY = y1

				startX = x2
				endX = x1
			}

			for ; startY <= endY; startY++ {
				coord := Coord{
					Y: startY,
					X: startX,
				}
				seen[coord]++
				if seen[coord] == 2 {
					intersect++
				}
				if startX > endX {
					startX--
				}
				if startX < endX {
					startX++
				}
			}
			break
		}
	}

	fmt.Println(intersect)
}
