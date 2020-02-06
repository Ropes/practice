package stringing

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestBalancedRunes(t *testing.T) {
	tests := []struct {
		runes    string
		balanced bool
	}{
		{
			runes:    "hi(hi)neh",
			balanced: true,
		},
		{

			runes:    "h(i(hi)neh)",
			balanced: true,
		},
		{
			runes:    "hi(()neh",
			balanced: false,
		},
		{
			runes:    "h()) wat",
			balanced: false,
		},
	}
	matches := map[rune]rune{
		'(': ')',
		')': 0,
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-runes-%d", i), func(t *testing.T) {
			input := make([]rune, 0)
			str := test.runes
			for len(str) > 0 {
				r, size := utf8.DecodeRuneInString(str)
				input = append(input, r)
				str = str[size:]
			}

			result := RuneBalanced(input, matches)
			if result != test.balanced {
				t.Errorf("rune balance test failed: %v", test.runes)
			}
		})
	}

}
