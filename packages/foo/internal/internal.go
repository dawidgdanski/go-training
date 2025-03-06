package internal

// Doubler function is accessible to both the parent package and the child packages of the internal package
func Doubler(a int) int {
	return a * 2
}
