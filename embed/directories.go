package embed

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

//go:embed help
var helpInfo embed.FS

func HelpFilesCheck() {
	printHelpFiles()
	printFileIfExists()
}

func printFileIfExists() {
	fileName := getFileNameArgument()
	data, err := helpInfo.ReadFile("help/" + fileName)
	if err != nil {
		fmt.Printf("An attempt to open a file '%s' has failed\n", fileName)
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}

func getFileNameArgument() string {
	if len(os.Args) < 2 {
		return "test"
	} else {
		return os.Args[1]
	}
}

func printHelpFiles() {
	fmt.Println("Contents:")
	err := fs.WalkDir(helpInfo, ".", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			_, fileName, _ := strings.Cut(path, string(os.PathSeparator))
			fmt.Println(fileName)
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
