package syntaxsamples

import "fmt"

func init() {
	fmt.Println("Hello from calculator.go")
}

type Calculator interface {
	Sum(left float64, right float64) float64

	Subtract(left float64, right float64) float64

	Divide(left float64, right float64) float64

	Multiply(left float64, right float64) float64
}

func NewCalculator() Calculator {
	return mapBasedCalculator{
		operationsRegistry: map[string]func(float64, float64) float64{
			"+": func(a, b float64) float64 { return a + b },
			"-": func(a, b float64) float64 { return a - b },
			"/": func(a, b float64) float64 { return a / b },
			"*": func(a, b float64) float64 { return a * b },
		},
	}
}

type mapBasedCalculator struct {
	operationsRegistry map[string]func(float64, float64) float64
}

func (calc mapBasedCalculator) Sum(left float64, right float64) float64 {
	return calc.operationsRegistry["+"](left, right)
}

func (calc mapBasedCalculator) Subtract(left float64, right float64) float64 {
	return calc.operationsRegistry["-"](left, right)
}

func (calc mapBasedCalculator) Divide(left float64, right float64) float64 {
	return calc.operationsRegistry["/"](left, right)
}

func (calc mapBasedCalculator) Multiply(left float64, right float64) float64 {
	return calc.operationsRegistry["*"](left, right)
}
