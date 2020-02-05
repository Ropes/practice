// Package norderedstatistic wat
package norderedstatistic

import (
	"math"
	"sort"
)

// O(n), no faster way if list unsorted.
func findSmallestNumber(input []int) int {
	smallest := int(math.MaxInt64)
	for _, v := range input {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}

func findNMins(input []int, minN int) []int {
	mins := make([]int, minN)
	//initialize array with max numbers
	for i := 0; i < minN; i++ {
		mins[i] = int(math.MaxInt64)
	}

	for _, v := range input {
		mins = sortSmallest(mins, minN, v)
	}

	return mins
}

func sortSmallest(input []int, N, newNum int) []int {
	input = append(input, newNum)
	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })
	return input[:N]
}
