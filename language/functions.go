package language

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

func init() {
	fmt.Println("Hello from the function.go", printFunctionName(myFunc))
	fmt.Println(
		"1 + 2 =", sum(func(a int, b int) int { return a + b }),
	)
	// structs can be used as named parameters
	_ = myFunc(MyFuncOpts{
		FirstName: "John",
		LastName:  "Doe",
	})

	fmt.Println("variadic parameters example: ", addTo(1, 23, 45))
	fmt.Println("variadic anonymous array example: ", addTo(3, []int{2, 3, 4, 5}...))
	arr := []int{2, 4, 6, 8}
	fmt.Println("variadic array example: ", addTo(3, arr...))

	result, _, _ := divAndRemainder(1, 2)
	_, _, err := divAndRemainder(1, 0)
	fmt.Println("Division", result, err)

	closureCanModifyOuterVariable()

}

type Add func(a int, b int) int

func sum(sumFunction Add) int {
	return sumFunction(1, 2)
}

func printFunctionName(fn interface{}) string {
	value := reflect.ValueOf(fn)
	ptr := value.Pointer()
	ffp := runtime.FuncForPC(ptr)
	return ffp.Name()
}

func myFunc(opts MyFuncOpts) error {
	fmt.Println("myFunc with MyFuncOpts", opts)
	return errors.New("an Error")
}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func addTo(base int, vals ...int) []int {
	out := make([]int, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divAndRemainder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	} else {
		return num / denom, num % denom, nil
	}
}

func closureCanModifyOuterVariable() {
	a := 10
	println("Before closure: ", a)
	func() {
		a += 20
		println("In closure: ", a)
	}()
	println("After closure: ", a)
}
