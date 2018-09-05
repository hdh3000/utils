package nthfib

func Fib(n int) int {
	if n == 0 {
		return 0
	}

	twoAgo, oneAgo, current := 0, 1, 1
	for i := 2; i <= n; i++ {
		current = twoAgo + oneAgo

		// Drop twoAgo to get ready for the next cycle
		twoAgo = oneAgo
		oneAgo = current
	}

	return current
}

func ProductOfArray(a []int) []int {
	if len(a) == 1 {
		return []int{0}
	}

	// If we have zero 0's we are in good shape
	// If we have one zero, we have a set of things to do
	// If we have two zeros, everything will be zero
	var zeroCount int

	var product int // Keep track of the total product, without zeros
	var startedProduct bool
	for _, v := range a {
		if !startedProduct && v != 0 {
			product = v
			startedProduct = true
			continue
		}

		if v == 0 {
			zeroCount++
			continue

		}

		// Collect the product without zeros..
		product = product * v

	}

	out := make([]int, len(a))
	if zeroCount > 1 {
		return out // It will ALWAYS be zero
	}

	for i, v := range a {
		if v == 0 {
			out[i] = product
			continue
		}

		if zeroCount == 1 {
			continue
		}

		out[i] = int(product / v)

	}

	return out
}

func CheckPalindrome(s string) bool {
	var offset int
	for {

		if len(s) <= 2*offset {
			break
		}

		if s[len(s)-1-offset] != s[offset] {
			return false
		}

		offset++
	}

	return true

}

//  [(0, 1), (3, 5), (4, 8), (10, 12), (9, 10)]

//  [(0, 1), (3, 5), (3, 8), (9, 10), (10, 12), (10, 13)]

// Overlap is
// "ending after start"
// "starting before end"

// 1 - - - - - - - - 9 - - -
// - 2 - - 3 - - - - - - - -
// - - - 3 - 5 - - - - - - -
// - - - - 4 - - - 8 - - - -

// rand7() gives me 1 1/7 of the time
// rand5() gives me 1 1/5 of the time

// rand7() gives me 7 1/7
// rand5() gives me 7 0/5

//
//  | - - - - - - | - - - - - - | - - - - - - | - - - - - - | - - - - - - |
//  | - - - - | - - - - | - - - - | - - - - | - - - - | - - - - | - - - - |
