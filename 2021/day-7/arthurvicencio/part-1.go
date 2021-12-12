package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func main() {

	inputFile, _ := ioutil.ReadFile("input.txt")

	parsedInput := strings.Split(string(inputFile), ",")

	nums := make([]int, 0)
	for _, v := range parsedInput {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}

	least := MaxInt
	for i := range nums {
		fuel := 0
		for _, v := range nums {
			dist := abs(v - i)
			fuel += dist
		}
		if fuel < least {
			least = fuel
		}
	}

	fmt.Println(least)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
