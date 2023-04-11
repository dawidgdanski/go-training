package person

import "fmt"

type Person struct {
	Name string
}

func (p *Person) WhoAmI() {
	fmt.Println("I am person of name ", p.Name)
}
