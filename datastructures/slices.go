package datastructures

import (
	"fmt"
	"strings"
)

/** ARRAYS */

func ArrayVar() {
	var scores [10]int
	scores[0] = 339
}

func ArrayInitialized() {
	scores := [4]int{1, 2, 3, 4}
	for index, value := range scores {
		fmt.Printf("scores[%d] = %d, ", index, value)
	}
	fmt.Println()
}

/** SLICES */

func SliceInitialized() {
	scores := []int{1, 5, 6, 234, 64353}
	fmt.Println(scores)
}

func SliceInitializedUsingMake() {
	scores := make([]int, 10)
	fmt.Println(scores)
}

func SliceInitializedUsingMakeWithoutResizeAndCrashing() {
	scores := make([]int, 0, 10)
	scores[7] = 1239
	fmt.Println(scores)
}

func SliceInitializedAndElementAppended() {
	scores := make([]int, 0, 10)
	scores = append(scores, 5)
	fmt.Println(scores)
}

func SliceInitializedAndResized() {
	scores := make([]int, 0, 10)
	scores = scores[0:7]
	fmt.Println(scores)
}

func SliceEnlargedWithDoubleLengthViaAppend() {
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println("Initial size of a Slice ", c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)
		newCap := cap(scores)
		if newCap != c {
			c = newCap
			fmt.Println("Slice has been resized to", c)
		}
	}
}

func SliceInitializationMethods() {
	names := []string{"leto", "jessica", "paul"}
	fmt.Println("NAMES:", names)
	checks := make([]bool, 10)
	fmt.Println("CHECKS:", checks)
	var elements []string
	fmt.Println("ELEMENTS:", elements)
	scores := make([]int, 0, 20)
	fmt.Println("SCORES:", scores)
}

func SliceModifyingSourceArray() {
	scores := []int{1, 2, 3, 4, 5}
	fmt.Println("Scores before:", scores)
	slice := scores[2:4]
	fmt.Println("Slice:", slice)
	slice[0] = 999
	fmt.Println("Scores after:", scores)
}

func StringSliceManipulation() {
	haystack := "the spice must flow"
	fmt.Printf("First space after 5th character (%s): %d\n", haystack[5:], strings.Index(haystack[5:], " "))
}

// SlicingSlices /** Slicing both slices and arrays shares results in creating slice or array sharing memory with the original slice or array
func SlicingSlices() {
	fmt.Println("====== Slicing Slices =======")
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}
