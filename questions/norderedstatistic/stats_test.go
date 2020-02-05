package norderedstatistic

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	x := int(math.MaxInt64)
	if x < 0 {
		t.Errorf("MaxInt64 does not coerce nicely: %v", x)
	}
	t.Logf("%v %v", reflect.TypeOf(x), x)
}

func TestSmallest(t *testing.T) {
	tests := []struct {
		input    []int
		smallest int
	}{
		{
			input:    []int{2, 1, 6, 3, 9, 7, 5},
			smallest: 1,
		},
		{
			input:    []int{2, -1, 6, 3, 9, 7, 5},
			smallest: -1,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			smallest := findSmallestNumber(test.input)
			if smallest != test.smallest {
				t.Errorf("wrong value returned: %v", smallest)
			}
		})
	}
}

func TestSortSmallestN(t *testing.T) {
	tests := []struct {
		input    []int
		N        int
		smallest []int
	}{
		{
			input:    []int{2, 1, 6, 3, 9, 7, 5},
			N:        1,
			smallest: []int{1},
		},
		{
			input:    []int{2, 1, 6, 3, 9, 7, 5},
			N:        2,
			smallest: []int{1, 2},
		},
		{
			input:    []int{2, -1, 6, 3, 9, 7, 5},
			N:        3,
			smallest: []int{-1, 2, 3},
		},
		{
			input:    []int{2, -1, -1, 3, 9, 7, 5},
			N:        3,
			smallest: []int{-1, -1, 2},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("testSmallest-%d", i), func(t *testing.T) {
			mins := sortSmallest(test.input, test.N, int(math.MaxInt64))
			if !reflect.DeepEqual(mins, test.smallest) {
				t.Errorf("expected smallest %v != calculated: %v", test.smallest, mins)
			}
		})
	}
}