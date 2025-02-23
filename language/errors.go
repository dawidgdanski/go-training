package language

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

const (
	ErrorFoo = Sentinel("foo error")
	ErrorBar = Sentinel("bar error")
)

func init() {
	openZipFile()
	LoginAndGetData(uuid.New().String(), "pass123", "secret.txt")
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
