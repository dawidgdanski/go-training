package main

import (
	"fmt"
	"go-training/datastructures"
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
		Person: &person.Person{Name: "Goku"},
		Power:  9000,
	}
	goku.SuperPower()
	goku.WhoAmI()
	fmt.Printf("Modifying %s's power (%d)\n", goku.Name, goku.Power)
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
	mathsamples.WorstScores()
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
