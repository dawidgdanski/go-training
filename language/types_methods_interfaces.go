package language

import "fmt"

/*
*
https://news.ycombinator.com/item?id=35784598
*/
func init() {
	var o Outer
	fmt.Println("o.Inner.X =", o.Inner.X)

	fmt.Println("o.Double() =", o.Double())
}

type Inner struct {
	X int
}

type Outer struct {
	Inner
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.X * 2)
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}
