package interpolate

import "testing"

func TestInterpolation(t *testing.T) {
	output := interpolate(inputEg)
	for i, v := range output {
		t.Logf("%d: %#v", i, v)
	}
}
