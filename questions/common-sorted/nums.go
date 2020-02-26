package sorted

func MatchingInts(a, b []int) []int {
	out := make([]int, 0)
	if len(a) != len(b) {
		return []int{}
	}

	j := 0
	for i := 0; i < len(a); i++ {
		if a[i] < b[j] { // Skip forward over smaller values
			continue
		}
		for j < len(b)-1 { // Increment j to skip small b values
			if b[j] >= a[i] {
				break
			} else if j < len(b)-1 {
				j++
			}
		}
		if a[i] == b[j] {
			out = append(out, a[i])
		}
	}
	return out
}
