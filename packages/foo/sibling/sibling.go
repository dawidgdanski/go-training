package sibling

import "go-training/packages/foo/internal"

func SiblingFoo(a int) int {
	return internal.Doubler(a)
}
