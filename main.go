// nolint
package main

import (
	"fmt"
	"go-training/concurrency"
	"go-training/datastructures"
	"go-training/errorhandling"
	"go-training/language"
	"go-training/logging"
	"go-training/mathsamples"
	"go-training/person"
	"go-training/saiyan"
	"os"
)

func main() {
	fmt.Println("Executable", os.Args[0])
	printArguments()
	variableFunction()
	power := getPower()
	fmt.Printf("Hello power %d!\n", power)
	gokuPower()
	datastructures.ArrayInitialized()
	newPower, exists := calculatePower("")
	if exists {
		fmt.Printf("New Power = %d\n", newPower)
	}
	//goku := Saiyan{}
	goku := saiyan.Saiyan{
		Person: &person.Person{FirstName: "Goku"},
		Power:  9000,
	}
	goku.SuperPower()
	goku.WhoAmI()
	fmt.Printf("Modifying %s's power (%d)\n", goku.FirstName, goku.Power)
	CopiedArgumentOperation(goku)
	fmt.Println("Nothing changes,", goku)
	PointerOperation(&goku)
	fmt.Println("Changed", goku)
	datastructures.SliceInitializedUsingMake()
	datastructures.SliceInitializedAndResized()
	datastructures.SliceEnlargedWithDoubleLengthViaAppend()
	datastructures.SliceInitializationMethods()
	fmt.Println("Powers extracted", saiyan.ExtractPowers([]*saiyan.Saiyan{&goku}))
	datastructures.SliceModifyingSourceArray()
	datastructures.StringSliceManipulation()
	datastructures.SlicingSlices()
	datastructures.SlicingSlicesWithSharedStorage()
	mathsamples.WorstScores()
	datastructures.MapSimpleCreationExample()
	datastructures.MapAnotherCreationAndInterationExample()
	logging.ProcessViaConsole()
	errorhandling.ParseAndPrintNumberOrLogError("234")
	errorhandling.ParseAndPrintNumberOrLogError("hello")
	fmt.Println("ERROR:", errorhandling.Process(0))
	language.Hello()
	concurrency.HelloFromConcurrency()
	fmt.Println("Sum of 2.0 and 3.0 is", language.NewCalculator().Sum(2.0, 3.0))
	var personPointer = language.MakePersonPointer("John", "Doe")
	fmt.Println("Another person: " + personPointer.FirstName)

	peopleSlice := language.CreatePeople(10_000_000)
	fmt.Println("People Slice length: ", len(peopleSlice))
	//errorhandling.ScanInput()
	//written, err := filesystem.CopyFile("Hello", "There")
	//fmt.Println("File copying result:", written, err)
}

func printArguments() {
	if len(os.Args) != 2 {
		fmt.Println("Hello, World!")
	} else {
		fmt.Println("It's over", os.Args[1])
	}
}

func variableFunction() {
	var power int
	power = 9000
	fmt.Printf("It's a power: %d units\n", power)
}

func getPower() int {
	return 9001
}

func gokuPower() {
	name, power := "Goku", 9001
	fmt.Printf("%s's power is over %d\n", name, power)
}

func gokuPower2() {
	power := 5002
	fmt.Printf("The default power is %d\n", power)

	name, power := "Goku", 9001
	fmt.Printf("%s's power is over %d\n", name, power)
}

func calculatePower(name string) (int, bool) {
	return len(name), true
}

func CopiedArgumentOperation(copy saiyan.Saiyan) {
	copy.Power += 10000
}

func PointerOperation(ref *saiyan.Saiyan) {
	ref.Power += 10000
}

func privateFunction() {
	fmt.Println("This function is private to the main package only")
}

func functionWithNamedReturnArguments(a, b int) (mul, div int) {
	//https://www.geeksforgeeks.org/named-return-parameters-in-golang/
	mul = a * b
	div = a / b
	return
}
