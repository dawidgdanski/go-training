package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ConcurrentRunnable func()

func init() {
	executeConcurrently("unsafe increment operation", func() {
		incrUnsafe()
	})
	executeConcurrently("synchronized increment operation", func() {
		incrSynchronized()
	})
	simpleChannelExample()
	timeoutChannelExample()
}

func HelloFromConcurrency() {
	fmt.Println("Hello from concurrency package")
}

var counter = 0

func executeConcurrently(name string, function ConcurrentRunnable) {
	fmt.Println("Concurrent execution of", name)
	for i := 0; i < 20; i++ {
		go function()
	}
	fmt.Println()
	time.Sleep(time.Millisecond * 10)
}

func incrUnsafe() {
	counter++
	fmt.Printf("%d; ", counter)
}

var (
	counter2 = 0
	lock     sync.Mutex
)

func incrSynchronized() {
	lock.Lock()
	defer lock.Unlock()
	counter2++
	fmt.Printf("%d; ", counter2)
}

func simpleChannelExample() {
	c := make(chan int, 10)
	finish := make(chan bool)
	waitGroup := sync.WaitGroup{}
	workerCount := 5
	waitGroup.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		worker := &Worker{id: i, waitGroup: &waitGroup}
		go worker.process(c, finish)
	}

	for j := 0; j < 20; j++ {
		c <- rand.Int()
		time.Sleep(time.Millisecond * 50)
	}
	finish <- true
	close(c)
	close(finish)
	waitGroup.Wait()
}

func timeoutChannelExample() {
	fmt.Println("\n=========== Timeout Channel Example ===========")
	channel1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "result 1"
	}()
	select {
	case response1 := <-channel1:
		fmt.Println(response1)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	channel2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "result 2"
	}()

	select {
	case response2 := <-channel2:
		fmt.Println(response2)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
	fmt.Println("=========== Timeout Channel Example ===========")
}

type Worker struct {
	waitGroup *sync.WaitGroup
	id        int
}

func (w *Worker) process(c chan int, finish chan bool) {
	for {
		select {
		case <-finish:
			fmt.Printf("Worker %d finished.\n", w.id)
			w.waitGroup.Done()
			return
		default:
			data := <-c
			fmt.Printf("Worker %d received %d\n", w.id, data)
			time.Sleep(time.Millisecond * 100)
		}
	}
}
