package saiyan

import (
	"fmt"
	"go-training/person"
)

type Saiyan struct {
	*person.Person
	Power    int
	Ancestor *Saiyan
}

func (s *Saiyan) SuperPower() {
	power := 10000
	fmt.Printf("%s: Increasing power %d\n", s.FirstName, power)
	s.Power += power
}

func (s *Saiyan) WhoAmI() {
	fmt.Printf("I am %s (%s)\n", s.FirstName, s.Person.FirstName)
}

func NewSaiyanPointer(name string, power int) *Saiyan {
	return &Saiyan{
		Person: &person.Person{FirstName: name},
		Power:  power,
		Ancestor: &Saiyan{
			Person:   &person.Person{FirstName: "Goku"},
			Power:    1001,
			Ancestor: nil,
		},
	}
}

func NewSaiyanPointerWithNew(name string, power int) *Saiyan {
	goku := new(Saiyan)
	goku.FirstName = name
	goku.Power = power
	return goku
}

func NewSaiyanCopy(name string, power int) Saiyan {
	return Saiyan{
		Person: &person.Person{FirstName: name},
		Power:  power,
	}
}

func ExtractPowers(saiyans []*Saiyan) []int {
	powers := make([]int, len(saiyans))
	for index, saiyanElement := range saiyans {
		powers[index] = saiyanElement.Power
	}
	return powers
}
