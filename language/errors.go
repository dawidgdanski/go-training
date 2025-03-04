package language

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"slices"

	"github.com/google/uuid"
)

const (
	ErrorFoo = Sentinel("foo error")
	ErrorBar = Sentinel("bar error")
)

func init() {
	openZipFile()
	data, err := LoginAndGetData(uuid.New().String(), "pass123", "secret.txt")
	if err != nil {
		fmt.Println("Error: " + err.Error())
	} else {
		fmt.Println("Data: " + string(data))
	}

	err = GenerateErrorBroken(true)
	fmt.Println("GenerateErrorBroken(true) returns non-nil error:", err != nil)
	err = GenerateErrorBroken(false)
	fmt.Println("GenerateErrorBroken(false) returns non-nil error:", err != nil)

	err = GenerateErrorOk(false)
	fmt.Println("GenerateErrorOk(false) returns non-nil error:", err != nil)
	err = GenerateErrorOk(true)
	fmt.Println("GenerateErrorOk(true) returns non-nil error:", err != nil)

	var myErr StatusErr
	errors.As(err, &myErr)
	fmt.Println("myErr: " + myErr.Error())
	if isStatusError := errors.Is(err, myErr); isStatusError {
		fmt.Println("This is StatusErr: " + myErr.Error())
	} else {
		fmt.Println("This is not StatusError: " + myErr.Error())
	}

	checkerError := fileChecker("not_here.txt")
	if checkerError != nil {
		fmt.Println(checkerError)
		if wrappedError := errors.Unwrap(checkerError); wrappedError != nil {
			fmt.Println(wrappedError)
		}
	}
	if errors.Is(checkerError, os.ErrNotExist) {
		fmt.Println("That file does not exist")
	}

	errorFromFunction := errorFunction()
	fmt.Println(errorFromFunction)

	validationError := validatePerson(Person{})
	if validationError != nil {
		fmt.Println(validationError)
	}

	mergeError := fmt.Errorf("first: %w, second: %w, third: %w", errors.New("first error"), errors.New("second error"), errors.New("third error"))
	fmt.Println(mergeError)

	detectInnerErrors()
	anError := createMyError()
	var myError MyError
	if errors.As(anError, &myError) {
		fmt.Println("MyError:", myError)
	}

	var myErrorAnother interface {
		Unwrap() []error
	}
	if errors.As(anError, &myErrorAnother) {
		fmt.Println("Errors: ", myErrorAnother.Unwrap())
	}

	for _, val := range []int{1, 2, 0, 6} {
		fmt.Printf("div %d/%d\n", 60, val)
		div60(val)
	}
}

func openZipFile() {
	data := []byte("This is not a zip file")
	fileReader := bytes.NewReader(data)
	_, err := zip.NewReader(fileReader, fileReader.Size())
	if errors.Is(err, zip.ErrFormat) {
		fmt.Println("Told you so")
	}
}

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func LoginAndGetData(uid, password, file string) ([]byte, error) {
	token, err := login(uid, password)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user: %s", uid),
		}
	}
	data, err := getData(token, file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("Data not found for file %s: %s", uid, file),
		}
	}
	fmt.Println(data)
	return []byte(""), nil
}

func getData(token string, file string) (string, error) {
	random := newRandom()
	if random.Int31()%2 == 0 {
		return fmt.Sprintf("Data %s", file), nil
	} else {
		return "", fmt.Errorf("could not find the data for file: %s", file)
	}
}

func newRandom() *rand.Rand {
	return rand.New(rand.NewSource(23))
}

func login(uid string, password string) (string, error) {
	random := newRandom()
	if random.Int31()%2 == 0 {
		return uid, nil
	} else {
		return "", errors.New("Unsuccessful login")
	}
}

func GenerateErrorBroken(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status:  InvalidLogin,
			Message: "Test Message",
		}
	}

	return genErr
}

func GenerateErrorOk(flag bool) error {
	if flag {
		return StatusErr{
			InvalidLogin,
			"Test Message",
		}
	} else {
		return nil
	}

}

// WRAPPING ERRORS

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("fileChecker - error while opening file: %w", err) // error wrapping
	}
	err = f.Close()
	if err != nil {
		return fmt.Errorf("fileChecker - error while closing file: %w", err) // error wrapping
	}
	return nil
}

func errorFunction() error {
	internalFunction := func() error {
		return errors.New("Internal Error")
	}

	err := internalFunction()

	return fmt.Errorf("errorFunction: %v", err) // err is not wrapped, only its message is copied
}

func validatePerson(person Person) error {
	var errs []error
	if len(person.FirstName) == 0 {
		errs = append(errs, fmt.Errorf("first name is empty"))
	}
	if person.MiddleName == nil {
		errs = append(errs, fmt.Errorf("middle name is empty"))
	}
	if len(person.LastName) == 0 {
		errs = append(errs, fmt.Errorf("last name is empty"))
	}

	return errors.Join(errs...)
}

type MyError struct {
	Codes  []int
	Errors []error
}

func (m MyError) Error() string {
	return errors.Join(m.Errors...).Error()
}

func (m MyError) Is(err error) bool {
	if m2, ok := err.(MyError); ok {
		return slices.Equal(m.Codes, m2.Codes)
	}
	return false
}

func (m MyError) Unwrap() []error {
	return m.Errors
}

func detectInnerErrors() {
	err := errorFunction()
	switch err := err.(type) {
	case interface{ Unwrap() error }:
		innerErr := err.Unwrap()
		fmt.Println(innerErr.Error())
	case interface{ Unwrap() []error }:
		innerErrs := err.Unwrap()
		for _, innerErr := range innerErrs {
			fmt.Println(innerErr.Error())
		}
	default:
		fmt.Println(err)
	}
}

func createMyError() error {
	return MyError{
		Errors: []error{errors.New("Hello")},
		Codes:  []int{1, 2, 3},
	}
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Recovered from panic: %+v\n", v)
		}

	}()

	fmt.Printf("%d\n", 60/i)
}
