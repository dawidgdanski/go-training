package errorhandling

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

func init() {
	fmt.Println("HELLO from errorhandling package!")
}

func ParseAndPrintNumberOrLogError(input string) {
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Not a valid number:", err)
	} else {
		fmt.Println(n)
	}
}

func Process(count int) error {
	if count < 1 {
		return errors.New("Input number below 1: " + strconv.Itoa(count))
	} else {
		return nil
	}
}

func ScanInput() {
	var input int
	_, err := fmt.Scan(&input)
	if err == io.EOF {
		fmt.Println("No more input!")
	}
}

func privateFunction() {
	fmt.Println("This function is private to the 'stdin' package only")
}
