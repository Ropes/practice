package curry

import (
	"context"
	"fmt"
)

func fibonacci(ctx context.Context, c chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-ctx.Done():
			fmt.Println("quit")
			return
		}
	}
}

func calculateFib(n int) int {
	var v int
	c := make(chan int)
	ctx, can := context.WithCancel(context.Background())
	go func() {
		for i := 0; i < n; i++ {
			v = <-c
			fmt.Println(v)
		}
		can()
	}()
	fibonacci(ctx, c)
	return v
}
