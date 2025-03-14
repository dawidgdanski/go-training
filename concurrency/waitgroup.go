package concurrency

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

func ProcessAndGather() {
	inputChannel := make(chan int, 10)
	go func() {
		for i := 0; i < 100; i++ {
			inputChannel <- rand.Int()
		}
		close(inputChannel)
	}()

	processor := func(value int) string {
		return strconv.Itoa(value)
	}
	result := processAndGather(inputChannel, processor, 5)
	fmt.Println("Process and gather - elements processed", len(result))
}

func processAndGather[T, R any](in <-chan T, processor func(T) R, num int) []R {
	out := make(chan R, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
		fmt.Println("Out channel closed")
	}()
	var result []R
	for v := range out {
		result = append(result, v)
	}
	return result
}
