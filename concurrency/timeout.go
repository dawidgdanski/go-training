package concurrency

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// TimeLimit - see https://go.dev/play/p/Npsa4Zvu2ym
func TimeLimit() {
	result, err := timeLimit(doSomeWork, 2*time.Second)
	fmt.Println(result, err)
}

func timeLimit[T any](worker func() T, limit time.Duration) (T, error) {
	out := make(chan T, 1)
	ctx, cancel := context.WithTimeout(context.Background(), limit)
	defer cancel()
	go func() {
		out <- worker()
	}()
	select {
	case result := <-out:
		return result, nil
	case <-ctx.Done():
		var zero T
		return zero, errors.New("work timed out")
	}
}

func doSomeWork() int {
	if x := rand.Int(); x%2 == 0 {
		return x
	} else {
		time.Sleep(5 * time.Second)
		return 100
	}
}
