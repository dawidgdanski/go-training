package language

import "fmt"

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

const (
	Field1 = 0
	Field2 = 1 + iota
	Field3 = 20
	Field4
	Field5 = iota
)

type BitField int

const (
	BitField1 BitField = 1 << iota
	BitField2
	BitField3
	BitField4
	BitField5
	BitField6
)

func init() {
	fmt.Println("Enums: ", Field1, Field2, Field3, Field4, Field5)
	fmt.Println("BitFields: ", BitField1, BitField2, BitField3, BitField4, BitField5, BitField6)
}
