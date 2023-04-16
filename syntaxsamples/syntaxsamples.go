package syntaxsamples

import "fmt"

func init() {
	fmt.Println("Hello from the syntax package")
	privateFunction()
	mul, div := functionWithNamedReturnArguments(105, 7)
	fmt.Println("MUL", mul, "DIV", div)
}

func privateFunction() {
	fmt.Println("This function is private to the main package only")
}

func functionWithNamedReturnArguments(a, b int) (mul, div int) {
	//https://www.geeksforgeeks.org/named-return-parameters-in-golang/
	mul = a * b
	div = a / b
	return
}

func Hello() {
	fmt.Println("Hello from the syntaxsamples package")
}
