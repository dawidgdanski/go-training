package generate

import "fmt"

type Direction int

const (
	_ Direction = iota
	North
	South
	East
	West
)

//go:generate stringer -type=Direction
func DirectionWithStringGenerated() {
	// Customizations: https://arjunmahishi.com/posts/golang-stringer
	fmt.Println(North.String())
}
