package concurrency

import (
	"context"
	"fmt"
	"time"
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

func ProcessChannel(ch chan int) []int {
	length := len(ch)
	results := make(chan int, length)
	for i := 0; i < length; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}
	var out []int
	for i := 0; i < length; i++ {
		out = append(out, <-results)
	}
	fmt.Println("Finished processing channel", out)
	return out
}

func process(v int) int {
	const seconds = 1
	time.Sleep(time.Duration(seconds) * time.Second)
	return v * 2
}
