package sorted

func MatchingInts(a, b []int) []int {
	out := make([]int, 0)
	if len(a) != len(b) {
		return []int{}
	}

	j := 0
	for i := 0; i < len(a); i++ {
		for ; j <= i; j++ {
			if b[j] >= a[i] {
				break
			}
		}
		if a[i] == b[j] {
			out = append(out, a[i])
		}
	}
	return out
}
