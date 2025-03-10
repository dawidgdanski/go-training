package embed

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed passwords.txt
var passwords string

func PasswordCheck() {
	pwds := strings.Split(passwords, "\n")
	if len(os.Args) > 1 {
		for _, password := range pwds {
			if password == os.Args[1] {
				fmt.Println("Password correct")
			}
		}
	} else {
		fmt.Println("Password incorrect")
	}
}
