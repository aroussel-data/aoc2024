package main

import (
	"fmt"
	"math"
	// "math"
)

func allNegative(diff_array []int) bool {
	for _, elem := range diff_array {
		if elem > 0 {
			return false
		}
	}
	return true
}

func allPositive(diff_array []int) bool {
	for _, elem := range diff_array {
		if elem < 0 {
			return false
		}
	}
	return true
}

func isOrdered(array []int) bool {
	if allNegative(array) || allPositive(array) {
		return true
	} else {
		return false
	}
}

func isStepSmallEnough(array []int) bool {
	var smallEnough bool = true
	for _, elem := range array {
		if math.Abs(float64(elem)) < 1 || math.Abs(float64(elem)) > 3 {
			smallEnough = false
		}
	}
	return smallEnough
}

func calculateDiff(array []int) []int {
	var diff_array []int
	for i := 0; i < len(array)-1; i++ {
		signed_diff := array[i] - array[i+1]
		diff_array = append(diff_array, signed_diff)
	}
	return diff_array
}

func main() {
	// 7 6 4 2 1  <-- safe
	// 1 2 7 8 9
	// 9 7 6 2 1
	// 1 3 2 4 5
	// 8 6 4 4 1
	// 1 3 6 7 9  <-- safe
	// Row is safe if two conditions are met:
	// - if all values are either decreasing or increasing
	// - if two adjacent values differ by between 1 and  3
	// Return the count of rows that are safe

	// First thing to check is if each row is either increasing or decreasing
	// bc if not then we short circuit and we know it's not safe

	// If first condition met, then we need distance between values of row
	// if any over 3 then short circuit as we know it's not safe

	// array_list := [][]int{
	// 	{7, 6, 4, 2, 1},
	// 	{1, 2, 7, 8, 9},
	// 	{9, 7, 6, 2, 1},
	// 	{1, 3, 2, 4, 5},
	// 	{8, 6, 4, 4, 1},
	// 	{1, 3, 6, 7, 9},
	// }

	var safe_rows int = 0
	for _, sub_array := range rows[50:55] {
		diff_array := calculateDiff(sub_array)
		// fmt.Println(diff_array)
		// fmt.Println(isOrdered(diff_array), isStepSmallEnough(diff_array))
		if isOrdered(diff_array) && isStepSmallEnough(diff_array) {
			safe_rows++
		}
	}
	fmt.Println(safe_rows)
}
