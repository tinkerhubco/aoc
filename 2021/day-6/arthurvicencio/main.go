package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	file, _ := ioutil.ReadFile("input.txt")

	raw := strings.Split(string(file), ",")

	lanternfishes := make([]int, 0)
	for _, val := range raw {
		var n int
		fmt.Sscanf(val, "%d", &n)
		lanternfishes = append(lanternfishes, n)
	}

	// part 1: 80
	days := 256

	type Next struct {
		Value int
		Day   int
	}

	memo := make(map[Next]int)

	var compute func(num int, day int) int
	compute = func(num int, day int) int {
		sum := 1

		key := Next{Value: num, Day: day}
		if _, ok := memo[key]; ok {
			return memo[key]
		}

		for d := day; d < days; d++ {
			switch num {
			case 0:
				num = 6
				sum += compute(8, d+1)
				break
			default:
				num--
				break
			}
		}
		memo[key] = sum

		return memo[key]
	}

	var lanternfishCount int
	for _, lanternfish := range lanternfishes {
		lanternfishCount += compute(lanternfish, 0)
	}

	fmt.Println(lanternfishCount)
}
