package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input string

var versionSum int

func main() {

	var buffer bytes.Buffer
	for _, r := range input {
		n, _ := strconv.ParseUint(string(r), 16, 64)
		buffer.WriteString(fmt.Sprintf("%04b", n))
	}

	packets := buffer.String()

	result, _ := parsePackets(packets, 0)

	fmt.Printf("Part 1: %d\n", versionSum)
	fmt.Printf("Part 2: %d\n", result)
}

func parsePackets(packets string, i int) (int, int) {

	version, _ := strconv.ParseUint(packets[i:i+3], 2, 64)
	i += 3

	versionSum += int(version)

	packetType, _ := strconv.ParseUint(packets[i:i+3], 2, 64)
	i += 3

	nums := make([]int, 0)

	if packetType != 4 {
		lengthTypeId, _ := strconv.ParseUint(packets[i:i+1], 2, 64)
		i += 1

		if lengthTypeId == 0 {
			length, _ := strconv.ParseUint(packets[i:i+15], 2, 64)
			i += 15
			end := i + int(length)
			for i < end {
				var newNum int
				newNum, i = parsePackets(packets, i)
				nums = append(nums, newNum)
			}
		} else {
			length, _ := strconv.ParseUint(packets[i:i+11], 2, 64)
			i += 11
			for x := 0; x < int(length); x++ {
				var newNum int
				newNum, i = parsePackets(packets, i)
				nums = append(nums, newNum)
			}
		}
	}

	switch packetType {
	case 0:
		sum := nums[0]
		for _, n := range nums[1:] {
			sum += n
		}
		return sum, i
	case 1:
		product := nums[0]
		for _, n := range nums[1:] {
			product *= n
		}
		return product, i
	case 2:
		min := nums[0]
		for _, n := range nums[1:] {
			if n < min {
				min = n
			}
		}
		return min, i
	case 3:
		max := nums[0]
		for _, n := range nums[1:] {
			if n > max {
				max = n
			}
		}
		return max, i
	case 4:
		var bin string
		for {
			bin += packets[i+1 : i+5]
			i += 5
			if packets[i-5] == '0' {
				break
			}
		}
		n, _ := strconv.ParseUint(bin, 2, 64)
		return int(n), i
	case 5:
		if nums[0] > nums[1] {
			return 1, i
		} else {
			return 0, i
		}
		break
	case 6:
		if nums[0] < nums[1] {
			return 1, i
		} else {
			return 0, i
		}
		break
	case 7:
		if nums[0] == nums[1] {
			return 1, i
		} else {
			return 0, i
		}
		break
	}

	return 0, -1
}
