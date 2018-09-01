package sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	gen := newRandIntPerm(time.Now().Unix())
	lens := []int{0, 1, 2, 3, 4, 20, 30, 50}
	for _, i := range lens {
		a := gen(i)
		t.Log("I", a)
		QuickSort(sortableInt{a: a})
		t.Log("O", a)
		checkSorted(t, a)
		t.Log("")
	}
}

func checkSorted(t *testing.T, a []int) {
	for i := range a {
		if i == 0 {
			continue
		}
		if a[i] < a[i-1] {
			t.Fatalf("not sorted at %d, %v", i, a)
		}
	}
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

func TestHeap(t *testing.T) {
	gen := newRandIntPerm(time.Now().Unix())
	lens := []int{0, 1, 2, 3, 4, 20, 30, 50}
	for _, i := range lens {
		a := gen(i)
		t.Log("I ", a)
		h := heapify(a)
		t.Log("H ", h.heap)
		for i := range h.heap {
			leftChild := h.leftChild(i)
			rightChild := h.rightChild(i)

			if leftChild != -1 {
				if h.heap[leftChild] > h.heap[i] {
					t.Fatalf("intial heap incorrect, %d > %d", h.heap[leftChild], h.heap[i])
				}
			}

			if rightChild != -1 {
				if h.heap[rightChild] > h.heap[i] {
					t.Fatalf("intial heap incorrect, %d > %d", h.heap[leftChild], h.heap[i])
				}
			}
		}

		HeapSort(a)
		t.Log("O ", a)
		checkSorted(t, a)
		t.Log("")
	}
}

func newRandIntPerm(seed int64) func(len int) []int {
	r := rand.New(rand.NewSource(seed))
	return func(len int) []int {
		out := make([]int, len)
		for i := 0; i < len; i++ {
			out[i] = r.Intn(len + 50)
		}
		return out
	}
}
