package syntaxsamples

import "fmt"

func init() {
	fmt.Println(
		"1 + 2 =", sum(func(a int, b int) int { return a + b }),
	)
}

type Add func(a int, b int) int

func sum(sumFunction Add) int {
	return sumFunction(1, 2)
}
