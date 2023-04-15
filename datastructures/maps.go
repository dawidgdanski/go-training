package datastructures

import "fmt"

type Node struct {
	Value    string
	Children map[string]*Node
}

func MapSimpleCreationExample() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]
	fmt.Println("Does the map have 'vegeta' key?", power, exists)
}

func MapAnotherCreationAndInterationExample() {
	lookup := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}

	fmt.Println("Lookup:", lookup)

	fmt.Println("Iterating over lookup:")
	for key, value := range lookup {
		fmt.Printf("lookup[%s] = %d\n", key, value)
	}
}
