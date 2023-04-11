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
	fmt.Printf("%s: Increasing power %d\n", s.Name, power)
	s.Power += power
}

func (s *Saiyan) WhoAmI() {
	fmt.Printf("I am %s (%s)\n", s.Name, s.Person.Name)
}

func NewSaiyanPointer(name string, power int) *Saiyan {
	return &Saiyan{
		Person: &person.Person{Name: name},
		Power:  power,
		Ancestor: &Saiyan{
			Person:   &person.Person{Name: "Goku"},
			Power:    1001,
			Ancestor: nil,
		},
	}
}

func NewSaiyanPointerWithNew(name string, power int) *Saiyan {
	goku := new(Saiyan)
	goku.Name = name
	goku.Power = power
	return goku
}

func NewSaiyanCopy(name string, power int) Saiyan {
	return Saiyan{
		Person: &person.Person{Name: name},
		Power:  power,
	}
}
