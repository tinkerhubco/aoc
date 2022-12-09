package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	X, Y int
}

// up, down, left, right and diagonals
var dirs = []Point{
	Point{-1, -1},
	Point{0, -1},
	Point{1, -1},

	Point{-1, 0},
	Point{1, 0},

	Point{-1, 1},
	Point{0, 1},
	Point{1, 1},
}

func main() {
	fmt.Printf("Part 1: %d\n", trackLastKnot(2, input))
	fmt.Printf("Part 2: %d\n", trackLastKnot(10, input))
}

func trackLastKnot(length int, input string) int {
	track := make(map[Point]bool)
	knots := make([]Point, length)
	for i := 0; i < length; i++ {
		knots[i] = Point{}
	}
	for _, line := range strings.Split(input, "\n") {
		var dir string
		var amount int
		fmt.Sscanf(line, "%s %d", &dir, &amount)

		for i := 0; i < amount; i++ {
			switch dir {
			case "U":
				knots[0].Y--
				break
			case "D":
				knots[0].Y++
				break
			case "L":
				knots[0].X--
				break
			case "R":
				knots[0].X++
				break
			}
			for i := 1; i < len(knots); i++ {
				if !isNearBy(knots[i], knots[i-1]) {
					knots[i] = moveTowards(knots[i], knots[i-1])
				}
			}
			track[knots[len(knots)-1]] = true
		}
	}
	return len(track)
}

func isNearBy(point1, point2 Point) bool {
	for _, d := range dirs {
		if (Point{point1.X + d.X, point1.Y + d.Y}) == point2 {
			return true
		}
	}
	return false
}

func moveTowards(point1, point2 Point) Point {
	dest := Point{point1.X, point1.Y}
	if point1.X < point2.X {
		dest.X++
	}
	if point1.X > point2.X {
		dest.X--
	}
	if point1.Y < point2.Y {
		dest.Y++
	}
	if point1.Y > point2.Y {
		dest.Y--
	}
	return dest
}
