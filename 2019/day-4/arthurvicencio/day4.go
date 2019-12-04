package main

import (
	"fmt"
	"strings"
	"strconv"
)

var input string = "109165-576723"
//var input string = "111220-111224"

func main() {
	var parsedInput []string = strings.Split(input, "-")
	min, _  := strconv.Atoi(parsedInput[0])
	max, _ := strconv.Atoi(parsedInput[1])
	
	c := 0
	for ; min <= max ; min++ {
		mmin := strconv.Itoa(min)
		double := false
		lr := 0
		lrc := 0
		has := false
		//fmt.Println(mmin)
		for i := 0; i < 5; i++ {
		        m1, _  := strconv.Atoi(string(mmin[i]))
		 	m2, _  := strconv.Atoi(string(mmin[i + 1]))
			//fmt.Println(m1, m2)
			if m1 > m2 {
				double = false
				break
			}
			if m1 == m2 {
				if m1 == lr {
					lrc++
				} else {
					lrc = 2
				}
				lr = m1
				double = true
				continue
			} else {
				if lrc == 2 {
					has = true
				}
				lr = 0
				lrc = 0
			}
		}
		if double == true {
			if lrc == 2 || has {
				c++
			}
		}
		
	}
	// 308
	// 1756
	// 1991
	fmt.Println(c)
}
