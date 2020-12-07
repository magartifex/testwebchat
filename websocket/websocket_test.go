package websocket

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const countCycleGenerateUserIDDouble int = 100000

func BenchmarkGenerateUserIDDouble(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	var double int

	data := make(map[string]*interface{})

	b.ResetTimer()

	for i := 0; i < countCycleGenerateUserIDDouble; i++ {
		s := generateUserID(nil)

		if _, ok := data[s]; ok {
			double++
			continue
		}

		data[s] = nil
	}

	fmt.Println("BenchmarkGenerateUserIDDouble double", double)
}
