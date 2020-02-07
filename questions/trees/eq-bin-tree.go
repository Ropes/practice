package trees

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	cha, chb := make(chan int, 1), make(chan int, 1)
	ctx, can := context.WithTimeout(context.Background(), 5*time.Second)
	defer can()

	done := false
	for done {
		select {
		case a := <-cha:
			b := <-chb
			if a != b {
				return false
			}
		case <-ctx.Done():
			done = true
			break
		}
	}
	return true
}

func Main() {
	ch := make(chan int, 1)
	t := tree.New(1)
	go Walk(t, ch)

	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		}
	}
}
