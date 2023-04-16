package filesystem

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(source, destination string) (written int64, err error) {
	src, srcOpenError := os.Open(source)
	if srcOpenError != nil {
		fmt.Println("Source File not found:", source)
		return
	}
	defer src.Close()

	dst, dstOpenError := os.Create(destination)
	if dstOpenError != nil {
		fmt.Println("Destination File not created:", source)
		return
	}
	defer dst.Close()

	return io.Copy(src, dst)
}
