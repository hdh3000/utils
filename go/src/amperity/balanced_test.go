package amperity

import (
	"testing"
)

func TestBalanced(t *testing.T) {
	testCases := []struct {
		str      string
		expected bool
	}{
		{"[abc]", true},
		{"[[", false},
		{"[)", false},
		{"{[a]bq()}", true},
		{"[(])", false},
		{"]()", false},
	}

	for i, tc := range testCases {
		t.Logf("testing %q\n", tc.str)
		if balanced(tc.str) != tc.expected {
			t.Fatalf("Failed on caese [%d], expected %v, got %v on value: %q", i, tc.expected, !tc.expected, tc.str)
		}
	}
}
