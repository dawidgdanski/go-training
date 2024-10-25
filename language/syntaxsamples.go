package language

import "fmt"

func init() {
	fmt.Println("Hello from the syntax package")
	privateFunction()
	mul, div := functionWithNamedReturnArguments(105, 7)
	fmt.Println("MUL", mul, "DIV", div)
	InitializedIfStatement()
}

func privateFunction() {
	fmt.Println("This function is private to the main package only")
}

func InitializedIfStatement() {
	count := 2
	if x := 10; count > x {
		fmt.Println("THIS", x)
	} else {
		fmt.Println("THAT", x)
	}
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
