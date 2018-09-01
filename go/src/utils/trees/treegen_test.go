package trees

import (
	"testing"
)

func TestGenTree(t *testing.T) {
	n := GenTree(6, false)
	BreadthFirstAllNodes(n)
}
