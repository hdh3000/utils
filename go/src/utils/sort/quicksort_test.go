package sort

import (
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {

	var checkSorted = func(a []int) {
		for i := range a {
			if i == 0 {
				continue
			}
			if a[i] < a[i-1] {
				t.Fatalf("not sorted at %d, %v", i, a)
			}
		}
	}

	for i := 0; i < 20; i++ {
		a := rand.Perm(i)
		t.Log(a)
		QuickSort(sortableInt{a: a})
		t.Log(a)
		checkSorted(a)
	}

	var repeats = []int{2, 14, 11, 6, 12, 5, 3, 15, 16, 9, 14, 10, 10, 7, 13, 4, 8}
	t.Log(repeats)
	QuickSort(sortableInt{a: repeats})
	t.Log(repeats)
	checkSorted(repeats)
}

type sortableInt struct {
	a []int
}

func (s sortableInt) Less(i, j int) bool {
	return s.a[i] < s.a[j]
}

func (s sortableInt) Len() int {
	return len(s.a)
}

func (s sortableInt) Swap(i, j int) {
	temp := s.a[i]
	s.a[i] = s.a[j]
	s.a[j] = temp
}
