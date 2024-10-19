package datastructures

import (
	"fmt"
	. "go-training/person"
	"sort"
)

func init() {
	people := []Person{
		{"Bob", "Bob LastName", 13},
		{"Alice", "Alice LastName", 42},
	}
	fmt.Println("People: ", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].FirstName < people[j].FirstName
	})
	fmt.Println("Sorted People", people)
}
