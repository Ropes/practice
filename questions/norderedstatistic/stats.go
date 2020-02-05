// Package norderedstatistic wat
package norderedstatistic

import (
	"fmt"
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

// Find Nth smallest number with traversal and splitting efficiency.
// Similar to insertion sort, traverse first number to position where it is smaller than numbers to right.
// Reduce input to have the first N values sorted in the list.
// []int{10,2,5,3,6,11,15,-1,-2} wE: 9
// []int{2,5,3,6,10,11,15,-1,-2}
// []int{2,5,3,6,10,11,-1,15,-2}
// []int{2,5,3,6,10,11,-1,-2,15} wE: 8
// []int{2,3,5,6,10,11,-1,-2,15} wE: 8
// []int{2,3,5,6,10,-1,11,-2,15} wE: 8
// []int{2,3,5,6,10,-1,-2,11,15} wE: 7
// []int{2,3,5,6,-1,10,-2,11,15} wE: 7
// []int{2,3,5,6,-1,-2,10,11,15} wE: 6
func findNthSmallest(input []int, n int) int {
	windowStart, windowEnd := 0, len(input)-1
	for (windowEnd - windowStart) > n {
		swpPt := 0
		fmt.Printf("input: %v\n", input)
		for i := windowStart + 1; i <= windowEnd; i++ {
			if input[i-1] > input[i] { // 10 > [i]
				input[i], input[i-1] = input[i-1], input[i] //swap the values
				swpPt = i                                   // move to next comparison value
			}
		}
		if swpPt < windowEnd {
			windowEnd = swpPt
		}
	}

	// Pick the Nth smallest(largest value from size N)
	smallest := input[0]
	for i := 0; i < windowEnd; i++ {
		if smallest < input[i] {
			smallest = input[i]
		}
	}
	return smallest
}

func sortSmallest(input []int, N, newNum int) []int {
	input = append(input, newNum)
	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })
	return input[:N]
}
