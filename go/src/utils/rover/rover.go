package rover

import (
	"math"
	"sort"
	"strings"
	"sync"
)

// ImgMetadata hold the information received from the rover
type ImgMetadata struct {
	Size   int
	Chunks []Chunk
}

// ImgRequest is the set of chunks we should request from the rover
type ImgRequest struct {
	chunks []Chunk
}

func (imgReq *ImgRequest) String() string {
	var ids = make([]string, len(imgReq.chunks))
	for i, c := range imgReq.chunks {
		ids[i] = c.Id
	}
	sort.Strings(ids)
	return strings.Join(ids, "\n")
}

// FindSmallestChunkSetRecursively takes ImgMetadata and returns the smallest set of chunks
// (by total bytes in the chunks) that can recreate the entire image
func FindSmallestChunkSetRecursively(imd *ImgMetadata) *ImgRequest {
	// The naive approach is to simply test all combinations of chunksById
	// and return the smallest combo that covers the img.
	// Lets try that...
	//
	// Drawbacks here are that this doesn't scale particularly well... as the list size grows
	// the number of combos will sky rocket, for the given tests though, it seems ok.
	//

	// First thing is to eliminate duplicates (assumes unique ids)
	var chunksById = make(map[string]Chunk)
	for i, c := range imd.Chunks {
		chunksById[c.Id] = imd.Chunks[i]
	}

	var uniqueChunks chunkSelectors
	for k := range chunksById {
		uniqueChunks = append(uniqueChunks, chunkSelector{Chunk: chunksById[k]})
	}

	// Now that we have our set of unique Ids, we need to generate unique combinations of those ids.
	// We do this by recursively creating a tree of possibilities.
	// If we have n chunks, we want to check all combinations of chunks sizes 1-n

	// bestComboForLengthN takes a given size of combination (n) and produces all the unique sets of chunks that are that size.
	var bestComboForLengthN func(n int, imgSize int, chunks chunkSelectors, index, countSelected, currentBest int) (int, []chunkSelector)
	bestComboForLengthN = func(n int, imgSize int, chunks chunkSelectors, index, countSelected, currentBest int) (int, []chunkSelector) {
		if n == countSelected {
			if chunks.fullyCovers(imgSize) {
				return chunks.size(), chunks
			}
			return 0, nil
		}

		if index == len(chunks) {
			// We went down a branch of the tree that didn't yield any valid combos
			return 0, nil
		}

		if chunks.size() > currentBest && currentBest != 0 {
			// We are already over size, and not even done. Lets get out of here.
			return 0, nil
		}

		optionTrue := make(chunkSelectors, len(chunks))
		copy(optionTrue, chunks)
		optionTrue[index].selected = true
		tBestSize, tBest := bestComboForLengthN(n, imgSize, optionTrue, index+1, countSelected+1, currentBest)

		optionFalse := make(chunkSelectors, len(chunks))
		copy(optionFalse, chunks)
		optionFalse[index].selected = false
		fBestSize, fBest := bestComboForLengthN(n, imgSize, optionFalse, index+1, countSelected, currentBest)

		if fBest == nil {
			return tBestSize, tBest
		}

		if tBest == nil {
			return fBestSize, fBest
		}

		if tBestSize < fBestSize {
			return tBestSize, tBest
		} else {
			return fBestSize, fBest
		}

	}

	var bestSize int
	var bestChunks chunkSelectors
	for i := 1; i <= len(uniqueChunks); i++ {
		pBestSize, pBestChunks := bestComboForLengthN(i, imd.Size, uniqueChunks, 0, 0, bestSize)
		if pBestChunks != nil && bestChunks == nil {
			bestChunks = pBestChunks
			bestSize = pBestSize
			continue
		}

		if pBestChunks == nil {
			// Nothing at this size covered it
			continue
		}

		if pBestSize < bestSize {
			// Len check is necessary because we could have an empty chunk,
			// and we don't want to include that as it would waste our time.
			bestSize = pBestSize
			bestChunks = pBestChunks
		}

	}

	var chunks []Chunk
	for _, c := range bestChunks {
		if c.selected {
			chunks = append(chunks, c.Chunk)
		}
	}

	return &ImgRequest{chunks: chunks}
}

// FindSmallestChunkSetSequentially takes ImgMetadata and returns the smallest set of chunks
// (by total bytes in the chunks) that can recreate the entire image by iterating through all possible combinations
// Unlike the recursive approach, where we build up a tree of possible combinations (and can therefore exit early when we determine that
// all the leaves of a particular branch will be too big, the sequential approach must go through each combination.
func FindSmallestChunkSetSequentially(imd *ImgMetadata, numWorkers int) *ImgRequest {
	// First thing is to eliminate duplicates (assumes unique ids)
	var chunksById = make(map[string]Chunk)
	for i, c := range imd.Chunks {
		chunksById[c.Id] = imd.Chunks[i]
	}

	var uniqueChunks chunkSelectors
	for k := range chunksById {
		uniqueChunks = append(uniqueChunks, chunkSelector{Chunk: chunksById[k]})
	}

	// Total number of possible combinations of sizes 0-n is 2^n
	totalCombos := int(math.Pow(2, float64(len(uniqueChunks))))

	// Using Gray codes, we can generate a unique combination from an index.
	// Because we know the total number of combinations in advance, and can provide a unique
	// index for each, its possible to shard out the process to a number of workers.
	// We will follow the fan-out-fan-in technique
	// Note the range is (start, end]
	var worker = func(chunks chunkSelectors, imgSize int, start, end int) []int {
		var bestSize int
		var bestChunks []int

		for i := start; i < end; i++ {
			seq := newGray(i, len(chunks))
			chunks.set(seq)

			size := chunks.size()
			if (size < bestSize || bestChunks == nil) && chunks.fullyCovers(imgSize) {
				bestSize = size
				bestChunks = seq
			}
		}

		return bestChunks

	}

	if numWorkers > totalCombos {
		// This produces as many workers as actually makes sense...
		numWorkers = totalCombos
	}

	workerRange := int(totalCombos / numWorkers)
	remainder := totalCombos % numWorkers

	bestChan := make(chan []int, numWorkers)
	wg := sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		start := i * workerRange
		end := start + workerRange

		if i == numWorkers {
			// Extend the last worker to cover the overage, and make sure we get the very last combo.
			end += remainder + 1
		}

		wg.Add(1)
		go func() {
			// Note it is important here to copy the slice, as the worker manipulates its own copy.
			bestChan <- worker(append(chunkSelectors{}, uniqueChunks...), imd.Size, start, end)
			wg.Done()
		}()
	}

	// Wait till our workers are done
	wg.Wait()
	close(bestChan)

	// Compare the results from each worker, and output the best one.
	var best []int
	var bestSize int

	for combo := range bestChan {
		if combo == nil {
			continue
		}

		uniqueChunks.set(combo)
		size := uniqueChunks.size()
		if size < bestSize || bestSize == 0 {
			best = combo
			bestSize = size
		}
	}

	// set our chunks to best, and return them
	uniqueChunks.set(best)

	var chunks []Chunk
	for _, c := range uniqueChunks {
		if c.selected {
			chunks = append(chunks, c.Chunk)
		}
	}

	return &ImgRequest{
		chunks: chunks,
	}

}

// newGray takes a decimal, and output width to produce a Gray code corresponding to that decimal number
// the value should be <= 2^width, or the number won't "fit"
func newGray(value, width int) []int {
	// First convert the decimal to binary, then binary to Gray
	// TODO(henry): probably some more efficient way of doing this.

	var decimalToBin func(int, []int) (int, []int)
	decimalToBin = func(in int, prev []int) (int, []int) {
		next := int(in / 2)
		remainder := in % 2

		if next == 0 {
			return 0, append([]int{remainder}, prev...)
		}

		return decimalToBin(next, append([]int{remainder}, prev...))

	}
	_, binary := decimalToBin(value, nil)

	pad := make([]int, width-len(binary))
	binary = append(pad, binary...)

	// Gray code from binary.
	gray := make([]int, width)
	gray[0] = binary[0]
	for i := 1; i < width; i++ {
		gray[i] = binary[i-1] ^ binary[i]
	}

	return gray

}

// Chunk corresponds to a piece of the image.
type Chunk struct {
	Id    string
	Start int
	Len   int
}

type chunkSelectors []chunkSelector

// chunkSelector is a datatype to help with generating combinations of chunks
type chunkSelector struct {
	Chunk
	selected bool
}

// fullyCovers checks if the image is covered by the set of selected chunks.
func (cc chunkSelectors) fullyCovers(size int) bool {
	if cc.size() < size {
		return false
	}

	// Really naive way of doing this, probably better to sort and check overlap.
	// Though my guess is that this isn't where the real performance gains are.
	// TODO(henry): comeback and do that if you have time.
	for i := 0; i < size; i++ {
		byteCovered := false

		// is the byte covered in at least one of our chunks?
		for _, c := range cc {
			if !c.selected {
				continue
			}

			if c.Start <= i && c.Start+c.Len > i {
				byteCovered = true
				break
			}
		}

		if !byteCovered {
			return false
		}
	}

	return true
}

// size returns the size of the chunks
func (cc chunkSelectors) size() int {
	sum := 0
	for _, c := range cc {
		if c.selected {
			sum += c.Len
		}
	}
	return sum
}

func (cc chunkSelectors) set(vv []int) {
	for i, v := range vv {
		cc[i].selected = v == 1
	}
}
