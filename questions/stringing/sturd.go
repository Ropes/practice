package stringing

// hi(hi)neh 	true
// h(i(hi)neh) 	true
// hi(()neh 	false
// h()) wat 	false
func RuneBalanced(str []rune, chkBalance map[rune]rune) bool {
	for i := 0; i < len(str); i++ {
		v, ok := chkBalance[str[i]]

		if ok {
			if v == 0 {
				return false // matched a bad rune
			}
			return isBalanced(str[i+1:], chkBalance, v)
		}
	}
	return true //default return is false
}

// isBalanced asserts that the sub list of runes has the balanced character
func isBalanced(s []rune, chkBalance map[rune]rune, r rune) bool {
	if len(s) < 1 {
		return false // expecting rune to balance
	}
	// Search for balanced rune 'r'
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == r { // matching rune found
			return RuneBalanced(s[:i], chkBalance)
		}
	}
	return false
}
