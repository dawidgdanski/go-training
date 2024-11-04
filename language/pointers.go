package language

import "fmt"

func init() {
	println("Pointers are about addresses")
	x := 10
	pointerToX := &x
	println(pointerToX)
	println(*pointerToX)
	z := 5 + *pointerToX
	println(z)
	accessingNilPointerRaisesError()
	var p = MakePersonPointer("Alice", "Bot")
	fmt.Println("A Person created: " + p.FirstName)

	firstSlice := []string{"Hello", "There"}
	fmt.Println("Before UpdateSlice: ", firstSlice)
	UpdateSlice(firstSlice, "What's up?")
	fmt.Println("After UpdateSlice: ", firstSlice)

	secondSlice := []string{"Welcome", "Goodbye"}
	fmt.Println("Before GrowSlice: ", secondSlice)
	GrowSlice(secondSlice, "")
	fmt.Println("After GrowSlice: ", secondSlice)
}

func accessingNilPointerRaisesError() {
	defer func() {
		if e := recover(); e != nil {
			println("Caught error: ", e)
		}
	}()

	println("Nil Pointer example")
	var x *int
	println(x == nil)
	println(*x)
}

func assigningPointerDirectlyDuringInstantiation() {
	p := Person{
		FirstName:  "Alice",
		MiddleName: makePointerTo("Bot"),
		LastName:   "AliceLastName",
	}

	fmt.Println(p)
}

func makePointerTo[T any](t T) *T {
	return &t
}

type Person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func MakePersonPointer(firstName string, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
	}
}

func UpdateSlice(slice []string, element string) {
	slice[len(slice)-1] = element
	fmt.Println("UpdateSlice: ", slice)
}

func GrowSlice(slice []string, element string) {
	grown := append(slice, element)
	fmt.Println("GrowSlice: ", grown)
}

func CreatePeople(size int) []Person {
	slice := make([]Person, size)

	for index, _ := range slice {
		suffix := string(rune(index + 1))
		slice[index] = Person{
			FirstName:  "First Name " + suffix,
			MiddleName: makePointerTo("Middle Name " + suffix),
			LastName:   "Last Name " + suffix,
		}
		return slice
	}

	return slice
}
