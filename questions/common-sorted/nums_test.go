package sorted

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortedNums(t *testing.T) {
	tests := []struct {
		A   []int
		B   []int
		exp []int
	}{
		{
			A:   []int{13, 27, 35, 40, 49, 55, 59},
			B:   []int{17, 35, 39, 40, 55, 58, 60},
			exp: []int{35, 40, 55},
		},
		{
			A:   []int{13, 27, 35, 40, 49},
			B:   []int{13, 39, 42, 58, 60},
			exp: []int{13},
		},
		{
			A:   []int{27, 35, 40, 49},
			B:   []int{13, 39, 42, 58, 60},
			exp: []int{},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			out := MatchingInts(test.A, test.B)
			if !reflect.DeepEqual(out, test.exp) {
				t.Errorf("incorrect results returned: %v expected: %v", out, test.exp)
			}
		})
	}

}
