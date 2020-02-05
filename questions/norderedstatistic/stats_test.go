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

func TestSmallestN(t *testing.T) {
	tests := []struct {
		input    []int
		N        int
		smallest []int
	}{
		{
			input:    []int{2, 1, 6, 3, 9, 7, 5},
			N:        7,
			smallest: []int{1, 2, 3, 5, 6, 7, 9},
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
		{
			input:    []int{2, -1, -1, 3, 9, 7, 5},
			N:        2,
			smallest: []int{-1, -1},
		},
		{
			input:    []int{2, -1, -1, 3, 9, 7, 5},
			N:        1,
			smallest: []int{-1},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("testFindSmallest-%d", i), func(t *testing.T) {
			mins := findNMins(test.input, test.N)
			if !reflect.DeepEqual(mins, test.smallest) {
				t.Errorf("findSmallest %v != %v", mins, test.smallest)
			}
		})
	}
}

func TestInPlaceSmallestN(t *testing.T) {
	tests := []struct {
		input []int
		N     int
		exp   int
	}{
		{
			input: []int{10, 2, 5, 3, 6, 11, 15, -1, -2},
			N:     3,
			exp:   2,
		},
		{
			input: []int{10, 2, 5, 3, 6, 11, 15, -1, -2},
			N:     5,
			exp:   5,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-in-place-%d", i), func(t *testing.T) {
			v := findNthSmallest(test.input, test.N)
			if v != test.exp {
				t.Errorf("wrong value returned: %d != %d", v, test.exp)
			}
		})
	}

}
