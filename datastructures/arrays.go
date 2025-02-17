package datastructures

import (
	"fmt"
	"go-training/person"
	"sort"
)

func init() {
	people := []person.Person{
		{FirstName: "Bob", LastName: "Bob LastName", Age: 13},
		{FirstName: "Alice", LastName: "Alice LastName", Age: 42},
	}
	fmt.Println("People: ", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].FirstName < people[j].FirstName
	})
	fmt.Println("Sorted People", people)
}
