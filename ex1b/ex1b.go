package main

import (
	"fmt"
	"math"
)

func main() {
	var left_slice = left[:]
	var right_slice = right[:]
	// NOTE: don't use a set for adding frequencies. I don't think doing an exponent
	// calc on the unique occurence of each int is a good idea as it causes
	// the calc result to be too high, possibly resulting in an overflowed value.
	// Stick with doing the frequency for each number, even if duplicated.
	left_set := make(map[uint64]uint64)
	for _, element := range left_slice {
		left_set[element] = 0
	}

	for _, element := range right_slice {
		if val, ok := left_set[element]; ok {
			left_set[element] = val + 1
		}
	}

	var similarity_score uint64
	for k, v := range left_set {
		if v > 0 {
			raised_val := uint64(math.Pow(float64(k), float64(v)))
			similarity_score += raised_val
		}
	}
	fmt.Println(similarity_score)
}
