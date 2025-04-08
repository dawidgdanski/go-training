package language

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func CountLetters() {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("String length:", counts)
}

func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func GzipCountLetters() {
	err := gzipCountLetters()
	if err != nil {
		slog.Error("error with gzipCountLetters", "msg", err)
	}
}

func gzipCountLetters() error {
	r, closer, err := buildGZipReader("my_data.txt.gz")
	if err != nil {
		return err
	}
	defer closer()
	counts, err := countLetters(r)
	if err != nil {
		return err
	}
	fmt.Println("my_data.txt.gz:", counts)
	return nil
}

type TodoItem struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func decodeResponseBody(res *http.Response) TodoItem {
	var data TodoItem
	err := json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	return data
}

func createTodoListRequest() *http.Request {
	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("X-My-Client", "Learning Go")
	return req
}

func JsonEncodingAndDecodingExample() {
	err := processPerson()
	if err != nil {
		panic(err)
	}
}

func processPerson() error {
	type AnotherPerson struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	toFile := AnotherPerson{
		Name: "Fred",
		Age:  40,
	}
	// Write it out
	tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())
	err = json.NewEncoder(tmpFile).Encode(toFile)
	if err != nil {
		return err
	}
	err = tmpFile.Close()
	if err != nil {
		return err
	}

	// Read it back in again
	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		return err
	}
	var fromFile AnotherPerson

	err = json.NewDecoder(tmpFile2).Decode(&fromFile)
	if err != nil {
		return err
	}
	err = tmpFile2.Close()
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", fromFile)
	return nil
}
