package lru

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	tests := []struct {
		cap       int
		valsAdded [][]int
		finalKeys []int
	}{
		{
			cap:       3,
			valsAdded: [][]int{{1, 12}, {2, 12}, {3, 19}, {4, 6}, {5, 9}},
			finalKeys: []int{3, 4, 5},
		},
		{
			cap:       2,
			valsAdded: [][]int{{4, 6}, {5, 9}},
			finalKeys: []int{4, 5},
		},
		{
			cap:       200,
			valsAdded: [][]int{{4, 6}, {5, 9}},
			finalKeys: []int{4, 5},
		},
		{
			cap:       1,
			valsAdded: [][]int{{1, 12}, {2, 12}, {3, 19}, {4, 6}, {5, 9}},
			finalKeys: []int{5},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.finalKeys), func(t *testing.T) {
			lru := NewLRUCache(test.cap)
			//Insert keys
			for _, tup := range test.valsAdded {
				k, v := tup[0], tup[1]
				lru.Put(k, v)
			}

			// Assert contents of LRU
			for _, fk := range test.finalKeys {
				v := lru.Get(fk)
				if v == -1 {
					t.Errorf("expected key [%d] not found", fk)
					t.Logf("%#v", lru.stack)
				}
			}

		})
	}
}
