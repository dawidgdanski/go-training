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

	tree := &IntTree{}
	tree.Insert(20).Insert(10).Insert(5).Insert(2).Insert(1)

	fmt.Println("The IntTree contains 0:", tree.Contains(0))
	fmt.Println("The IntTree contains 5", tree.Contains(5))
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

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	} else if val < it.val {
		it.left = it.left.Insert(val)
		return it
	} else {
		it.right = it.right.Insert(val)
		return it
	}
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}
