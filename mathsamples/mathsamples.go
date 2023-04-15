package mathsamples

import (
	"fmt"
	"math/rand"
	"sort"
)

func WorstScores() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	copy(worst, scores[:5])
	fmt.Println(worst)
}
