package person

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *Person) WhoAmI() {
	fmt.Println("I am person of name ", p.FirstName)
}
