package fastfactorial

import (
	"fmt"
	"testing"
)

func TestFastFib(t *testing.T) {
	tests := []struct {
		N   int
		Exp int
	}{
		{
			N:   0,
			Exp: 0,
		},
		{
			N:   1,
			Exp: 1,
		},
		{
			N:   2,
			Exp: 1,
		},
		{
			N:   3,
			Exp: 2,
		},
		{
			N:   4,
			Exp: 3,
		},
		{
			N:   5,
			Exp: 5,
		},
		{
			N:   6,
			Exp: 8,
		},
		{
			N:   10,
			Exp: 55,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]n:%d", i, test.N), func(t *testing.T) {
			x := FastFib(test.N)
			if x != test.Exp {
				t.Errorf("incorrect fib calculated: %d != %d", x, test.Exp)
			}
		})
	}

}
