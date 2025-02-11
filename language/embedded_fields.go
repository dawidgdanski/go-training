package language

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee // embedded field. Wrapping type inherits all embedded field type's attributes
	Reports  []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	return m.Reports
}

type Inner struct {
	X int
}

type Outer struct {
	Inner
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.X * 2)
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func dynamicDispatchExample() {
	var o Outer
	fmt.Println("o.Inner.X =", o.Inner.X)

	fmt.Println("o.Double() =", o.Double())
}

func init() {
	m := Manager{
		Employee: Employee{
			Name: "John",
			ID:   "12345",
		},
		Reports: []Employee{
			{
				Name: "Alice",
				ID:   "12345",
			},
		},
	}
	fmt.Println("Manager: ", m.Name, m.ID)
	fmt.Println("Find new employees: ", m.FindNewEmployees())
	dynamicDispatchExample()
}
