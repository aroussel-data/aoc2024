package main

import (
	"fmt"
	"slices"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var left_slice = left[:]
	var right_slice = right[:]
	slices.Sort(left_slice)
	slices.Sort(right_slice)
	var total int
	for i := 0; i < len(left_slice); i++ {
		total += Abs(left_slice[i] - right_slice[i])
	}
	fmt.Println(total)
}
