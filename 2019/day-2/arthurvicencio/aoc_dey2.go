package main

import (
	"fmt"
	"strings"
	"strconv"
)

//var input string = `1,1,1,4,99,5,6,0,99`
var input = `1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,9,1,19,1,19,5,23,1,23,5,27,2,27,10,31,1,31,9,35,1,35,5,39,1,6,39,43,2,9,43,47,1,5,47,51,2,6,51,55,1,5,55,59,2,10,59,63,1,63,6,67,2,67,6,71,2,10,71,75,1,6,75,79,2,79,9,83,1,83,5,87,1,87,9,91,1,91,9,95,1,10,95,99,1,99,13,103,2,6,103,107,1,107,5,111,1,6,111,115,1,9,115,119,1,119,9,123,2,123,10,127,1,6,127,131,2,131,13,135,1,13,135,139,1,9,139,143,1,9,143,147,1,147,13,151,1,151,9,155,1,155,13,159,1,6,159,163,1,13,163,167,1,2,167,171,1,171,13,0,99,2,0,14,0`
func main() {

	f, i, j := find()
	fmt.Println(f[0], i, j, 100 * i + j)
}

func find() ([]int64, int64, int64) {
	var ip []string = strings.Split(input, ",")
	var o []int64

	for _, v := range ip {
	  n, _ := strconv.ParseInt(v, 10, 64)
	  o = append(o, n)
	}
	
	//for k := 1; k < 99; k++ {
	for i := 1; i < 99 ; i++ {
		for j := 1 ; j < 99 ; j++ {
			//fmt.Println(int64(i), int64(j))
			d := append(o[:0:0], o...)
			d[1] = int64(i)
			d[2] = int64(j)
			n := get(d)
			//fmt.Println(100 * n[1] + n[2])
			if n[0] == 19690720 {
				return n, n[1], n[2]
			}
		}
	}
	//}
	//fmt.Println(o)
	return o, 0, 0
}

func get(o []int64) []int64 {
	var e int = len(o)
	for i := 0; i < e; i++ {

		if i % 4 == 0 {
			if o[i] == 99 {
			   //fmt.Println(o[0])
			   continue
			}
			if i + 1 >= e || i + 3 >= e || i + 2 >= e {
				return o
			}
			if o[i + 1] >= int64(e) || o[i + 3] >= int64(e) || o[i + 2] >= int64(e) {
				return o
			}

			if o[i] == 1 {
				o[o[i + 3]] = o[o[i + 1]] + o[o[i + 2]]
			}
			if o[i] == 2 {
				o[o[i + 3]] = o[o[i + 1]] * o[o[i + 2]]
			}
			i += 3
			
		}
	}
	return o
}


