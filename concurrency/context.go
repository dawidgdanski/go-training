package concurrency

import (
	"context"
	"fmt"
)

func countTo(context context.Context, n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			select {
			case <-context.Done():
				return
			case ch <- i:
				fmt.Println("Counting...", i)
			}
		}
	}()
	return ch
}

func Count(n int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := countTo(ctx, n)

	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println("Read:", i)
	}
}
