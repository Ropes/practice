package interpolate

const gap = 5

var inputEg = [][]int{{0, 0}, {5, 5}, {15, 15}, {20, 25}, {40, 40}}

// b = y2/y1
// a = y-ax
// solving for:
// y = b-ax

func xVal(in []int) int {
	return in[0]
}
func yVal(in []int) int {
	return in[1]
}

func interpolate(input [][]int) [][]int {
	output := make([][]int, 0)

	for i := 0; i < len(input); i++ {
		if i == len(input)-1 {
			output = append(output, input[i])
			break
		}
		tup0, tup1 := input[i], input[i+1]

		// determine if interpolation necessary
		xDiff := xVal(tup1) - xVal(tup0)
		if xDiff > gap {
			// Interpolate points
			for j := 1; j < xDiff/gap; j++ {
				x1, y1 := xVal(tup1), yVal(tup1)
				x0, y0 := xVal(tup0), yVal(tup0)
				xNew := x0 + (j * 5)
				// Proper calculation function: https://en.wikipedia.org/wiki/Linear_interpolation#Linear_interpolation_between_two_known_points
				y := (y0*(x1-xNew) + y1*(xNew-x0)) / (x1 - x0)
				output = append(output, []int{xNew, y})
			}
		} else {
			output = append(output, input[i])
		}
	}
	return output
}
