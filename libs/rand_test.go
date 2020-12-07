package libs

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unicode/utf8"
)

func TestRandStringBytes(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"null",
			args{
				0,
			},
		},
		{
			"plus",
			args{
				5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandStringBytes(tt.args.n); utf8.RuneCountInString(got) != int(tt.args.n) {
				t.Errorf("RandStringBytes() = %v", got)
			}
		})
	}
}

const (
	countRandStringSpeed       uint = 128
	countRandStringDouble      uint = 16
	countCycleRandStringDouble int  = 1000000
)

func BenchmarkRandStringSpeed(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RandStringBytes(countRandStringSpeed)
	}
}

func BenchmarkRandStringDouble(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	var double int

	data := make(map[string]*interface{})

	b.ResetTimer()

	for i := 0; i < countCycleRandStringDouble; i++ {
		s := RandStringBytes(countRandStringDouble)

		if _, ok := data[s]; ok {
			double++
			continue
		}

		data[s] = nil
	}

	fmt.Println("BenchmarkRandStringBytes double", double)
}
