package sort

func HeapSort(a []int) {
	h := heapify(a)
	for i := 1; i < len(h.heap); i++ {
		h.swap(0, h.len()-1)
		h.ignore++
		h.siftDown(0)
	}
}

type heap struct {
	heap   []int
	ignore int
}

func (h *heap) len() int {
	// If the heap is [0,1,0,2] and
	// h.ignore = 1, then heap length is three
	// (2 is out of the heap)
	// This lets us both heapify, and then sort our array in place
	return len(h.heap) - h.ignore
}

func (h *heap) leftChild(i int) int {
	lc := 2*i + 1
	if lc >= h.len() {
		return -1
	} else {
		return lc
	}
}

func (h *heap) rightChild(i int) int {
	rc := 2*i + 2
	if rc >= h.len() {
		return -1
	} else {
		return rc
	}
}

func (h *heap) swap(i, j int) {
	temp := h.heap[i]
	h.heap[i] = h.heap[j]
	h.heap[j] = temp
}

func heapify(a []int) *heap {
	if len(a) < 2 {
		// Anything of length 0, or 1 is by definition sorted
		return &heap{heap: a}
	}

	h := &heap{heap: a, ignore: len(a)}
	for i := 0; i < len(a); i++ {
		h.ignore-- // introduce a new un-ignored value into the heap
		for j := h.len() - 1; j > 0; j-- {
			// shift the value introduced into the heap
			// up to the front, and re-align the heap
			h.swap(j-1, j)
		}

		h.siftDown(0)
	}
	return h
}

// Sift down starts from a position in the tree,
// and ensures that that element is correctly inserted in the tree
func (h *heap) siftDown(i int) {
	if i == -1 {
		return
	}

	leftChild := h.leftChild(i)
	rightChild := h.rightChild(i)

	if leftChild != -1 && h.heap[leftChild] > h.heap[i] {
		h.swap(i, leftChild)
	}

	if rightChild != -1 && h.heap[rightChild] > h.heap[i] {
		h.swap(i, rightChild)
	}

	h.siftDown(leftChild)
	h.siftDown(rightChild)

}
