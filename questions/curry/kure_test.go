package curry

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		N      int
		ExpFib int
	}{
		{
			N:      1,
			ExpFib: 0,
		},
		{
			N:      2,
			ExpFib: 1,
		},
		{
			N:      3,
			ExpFib: 1,
		},
		{
			N:      9,
			ExpFib: 21,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.N), func(t *testing.T) {
			v := calculateFib(test.N)
			if test.ExpFib != v {
				t.Errorf("fibonnaci values don't match: %d != %d", test.ExpFib, v)
			}
		})
	}

}
