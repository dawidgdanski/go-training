package language

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
)

func init() {
	openZipFile()
}

func openZipFile() {
	data := []byte("This is not a zip file")
	fileReader := bytes.NewReader(data)
	_, err := zip.NewReader(fileReader, fileReader.Size())
	if errors.Is(err, zip.ErrFormat) {
		fmt.Println("Told you so")
	}

}
