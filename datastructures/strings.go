package datastructures

import "fmt"

func init() {
	stringsAndByteArrays()
}

func stringsAndByteArrays() {
	aString := "the spice must flow"
	fmt.Println("STRA", aString)
	bytesConvertedFromString := []byte(aString)
	fmt.Println("bytesConvertedFromString", bytesConvertedFromString)
	stringFromBinary := string(bytesConvertedFromString)
	fmt.Println("stringFromBinary", stringFromBinary)
}
