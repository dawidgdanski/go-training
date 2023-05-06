package syntaxsamples

import "fmt"

func init() {
	fmt.Println("Hello from defer_panic_recover")
	deferFirst()
	deferStack()
	panicAndRecover()
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

func panicAndRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panicAndRecover: Recovered in panicAndRecover()", r)
		}
	}()
	fmt.Println("panicAndRecover: Calling g.")
	g(0)

}

func g(i int) {
	if i > 3 {
		fmt.Println("g: Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("g: Defer in g", i)
	fmt.Println("g: Printing in g", i)
	g(i + 1)
}
