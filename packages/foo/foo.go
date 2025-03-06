package foo

import "go-training/packages/foo/internal"

func FooDoubler(a int) int {
	return internal.Doubler(a)
}
