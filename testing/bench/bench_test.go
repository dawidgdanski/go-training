package bench

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Set Up")
	makeData()
	exitVal := m.Run()
	fmt.Println("Tear Down")
	_ = os.Remove("testdata/data.txt")
	os.Exit(exitVal)
}

// makeData makes our data file for us. Rather than checking in a large file, we recreate it for the test.
// By setting the random seed to the same value every time, we ensure that we generate the same file every time.
// This random seed generates a file that's 65,204 bytes long.
func makeData() {
	file, err := os.Create("testdata/data.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	randomGenerator := rand.New(rand.NewSource(1))
	for i := 0; i < 10000; i++ {
		data := makeWord(randomGenerator.Intn(10) + 1)
		_, _ = file.Write(data)
	}
}

func makeWord(l int) []byte {
	out := make([]byte, l+1)
	for i := 0; i < l; i++ {
		out[i] = 'a' + byte(rand.Intn(26))
	}
	out[l] = '\n'
	return out
}

func TestFileLen(t *testing.T) {
	result, err := FileLen("testdata/data.txt", 1)
	if err != nil {
		t.Fatal(err)
	}
	if result < 1 {
		t.Error("Expected not empty, got", result)
	}
}

var blackhole int

func BenchmarkFileLen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := FileLen("testdata/data.txt", 1)
		if err != nil {
			b.Fatal(err)
		}
		blackhole = result
	}
}

func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := FileLen("testdata/data.txt", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
