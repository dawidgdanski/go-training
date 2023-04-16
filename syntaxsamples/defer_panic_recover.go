package syntaxsamples

import "fmt"

func init() {
	fmt.Println("Hello from defer_panic_recover")
	deferFirst()
	deferStack()
}

// https://go.dev/blog/defer-panic-and-recover

func deferFirst() {
	i := 0
	defer fmt.Println("deferFirst: defer", i)
	i++
	fmt.Println("deferFirst: finishing")
	return
}

func deferStack() {
	for i := 0; i < 4; i++ {
		defer fmt.Printf("deferStack: defer %d\n", i)
	}
	fmt.Println("deferStack: finishing")
}
