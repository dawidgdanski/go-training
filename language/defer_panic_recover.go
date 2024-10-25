package language

import "fmt"

func init() {
	fmt.Println("Hello from defer_panic_recover")
	deferFirst()
	deferStack()
	panicAndRecover()
	deferExample()
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

func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first: ", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second: ", val)
	}(a)
	a = 30
	defer func(val int) int {
		fmt.Println("There is no way to read the value returned ", val)
		return val
	}(a)
	a = 40
	fmt.Println("exiting: ", a)
	return a
}
