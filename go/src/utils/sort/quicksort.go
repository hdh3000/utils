package sort

import (
	"sort"
)

func QuickSort(a sort.Interface) {
	var isRelativeySorted = func(pivot, comparison int) bool {
		if pivot < comparison {
			return a.Less(pivot, comparison)
		}

		if pivot > comparison {
			// Pivot should be less than the comparison because its to the left
			return !a.Less(pivot, comparison)
		}

		// if the pivot and the comparison are the same index, this is true
		panic("shouldn't happen")
	}

	var sorter func(sort.Interface, int, int)
	sorter = func(a sort.Interface, start, end int) {
		// First select a pivot, and place it correctly in the array
		// (all numbers left are <= than, all numbers right are >=)
		pivotIndex := int((end-start)/2) + start

		comparisonIndex := start
		for comparisonIndex < pivotIndex {
			if isRelativeySorted(pivotIndex, comparisonIndex) {
				comparisonIndex++

			} else {
				pivotIndex--
				a.Swap(pivotIndex, pivotIndex+1)
				if pivotIndex != comparisonIndex {
					a.Swap(pivotIndex+1, comparisonIndex)
				}
			}
		}

		// Now run the same process the other way
		comparisonIndex = end - 1
		for comparisonIndex > pivotIndex {
			if isRelativeySorted(pivotIndex, comparisonIndex) {
				comparisonIndex--
			} else {
				pivotIndex++
				a.Swap(pivotIndex, pivotIndex-1)
				if pivotIndex != comparisonIndex {
					a.Swap(pivotIndex-1, comparisonIndex)
				}
			}
		}

		if pivotIndex-start > 1 {
			sorter(a, start, pivotIndex)
		}

		if end-pivotIndex > 1 {
			// Exclude the pivot index from the sort, as it is in the right place...
			sorter(a, pivotIndex+1, end)
		}
	}

	sorter(a, 0, a.Len())
}
