package child

import "go-training/packages/foo/internal"

func ChildDoubler(a int) int {
	return internal.Doubler(a)
}
