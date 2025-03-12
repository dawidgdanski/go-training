package concurrency

import "fmt"

func deadLockExample() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		inGoroutine := 1
		ch1 <- inGoroutine // the goroutine cannot make a progress because ch1 cannot be written to until it is read
		fromMain := <-ch2
		fmt.Println("goroutine:", inGoroutine, fromMain)
	}()
	inMain := 2
	ch2 <- inMain // the main goroutine cannot make a progress because ch1 cannot be written to until it is read
	fromGoroutine := <-ch1
	fmt.Println("main:", inMain, fromGoroutine)
}

func DeadLockEliminated() {
	//deadLockExample()
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		inGoroutine := 1
		ch1 <- inGoroutine
		fromMain := <-ch2
		fmt.Println("goroutine:", inGoroutine, fromMain) // still not perfect though, since this log is not executed.
	}()
	inMain := 2
	var fromGoroutine int
	select {
	case ch2 <- inMain:
	case fromGoroutine = <-ch1:
	}
	fmt.Println("main:", inMain, fromGoroutine)

	// Anytime a closure uses a variable whose value might change, use a parameter to pass a copy of the variable’s current value into the closure.
	//for _, v := range a {
	//	go func(val int) {
	//		ch <- val * 2
	//	}(v)
	//}
}
