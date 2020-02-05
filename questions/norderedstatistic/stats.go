// Package norderedstatistic wat
package norderedstatistic

import "math"

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
