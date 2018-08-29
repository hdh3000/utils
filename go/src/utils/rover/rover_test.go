package rover

import (
	"testing"
)

func TestNewBinary(t *testing.T) {
	width := 8
	tests := []struct {
		decimal int
		gray    []int
	}{
		{
			decimal: 0,
			gray:    []int{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			decimal: 1,
			gray:    []int{0, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			decimal: 8,
			gray:    []int{0, 0, 0, 0, 1, 1, 0, 0},
		},
		{
			decimal: 100,
			gray:    []int{0, 1, 0, 1, 0, 1, 1, 0},
		},
	}

	for _, expected := range tests {
		gray := newGray(expected.decimal, width)

		for i := range expected.gray {
			if expected.gray[i] != gray[i] {
				t.Fatalf("bad gray value, expected %v, got %v", expected.gray, gray)
			}
		}
	}
}
