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
