package nthfib

import "testing"

func TestNthFib(t *testing.T) {
	testCases := []struct {
		n, exp int
	}{
		{n: 0, exp: 0},
		{n: 1, exp: 1},
		{n: 2, exp: 1},
		{n: 3, exp: 2},
		{n: 4, exp: 3},
		{n: 11, exp: 89},
	}

	for _, tc := range testCases {
		if f := Fib(tc.n); f != tc.exp {
			t.Fatalf("Wrong number for n of %d, expected %d, got %d", tc.n, tc.exp, f)
		}
	}
}

func TestProductOfArray(t *testing.T) {
	testCases := []struct {
		n   []int
		exp []int
	}{
		{n: []int{}, exp: []int{}},
		//{n: []int{1}, exp: []int{}},
		{n: []int{1, 7, 3, 4}, exp: []int{84, 12, 28, 21}},
		{n: []int{1, -7, 3, -4}, exp: []int{84, -12, 28, -21}},
		{n: []int{0, -7, 3, -4}, exp: []int{84, 0, 0, 0}},
	}

	for i, tc := range testCases {
		pa := ProductOfArray(tc.n)
		for j, v := range pa {
			if tc.exp[j] != v {
				t.Fatalf("Wrong value for output (case %d) expected %d, got %d", i, tc.exp[j], v)
			}
		}
	}
}

func TestCheckPalindrome(t *testing.T) {
	testCases := []struct {
		v   string
		exp bool
	}{
		{"civic", true},
		{"civil", false},
		{"livci", false},
		{"AAABBBAAA", true},
	}

	for _, tc := range testCases {
		if f := CheckPalindrome(tc.v); f != tc.exp {
			t.Fatalf("thought %s was something it wasnt (expected %s, but got %s)", tc.v, tc.exp, f)
		}
	}

}
