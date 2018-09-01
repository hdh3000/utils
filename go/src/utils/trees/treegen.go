package trees

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Name     string
	Level    int
	Parent   *Node
	Children []*Node
}

func (n *Node) NewChild(name string) *Node {
	c := &Node{
		Name:     name,
		Level:    n.Level + 1,
		Parent:   n,
		Children: nil,
	}
	n.Children = append(n.Children, c)
	return c
}

func GenTree(depth int, binary bool) *Node {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var addChildren func(*Node, int, int, bool)
	addChildren = func(n *Node, maxDepth int, maxChildren int, balanced bool) {
		if n.Level == maxDepth {
			return
		}

		numChildren := 0
		if balanced {
			numChildren = maxChildren
		} else {
			numChildren = rand.Intn(maxChildren + 1)
		}

		for i := 0; i < numChildren; i++ {
			addChildren(n.NewChild(n.Name+string(alphabet[i])), maxDepth, maxChildren, balanced)
		}
	}

	n := &Node{
		Name:  "_",
		Level: 0,
	}

	childCount := func() int {
		if binary {
			return 2
		} else {
			return 10
		}
	}

	addChildren(n, depth, childCount(), binary)

	return n

}

func DepthFirstAllNodes(n *Node) []*Node {
	var out []*Node
	for i := range n.Children {
		out = append(out, DepthFirstAllNodes(n.Children[i])...)
	}
	return out
}

func BreadthFirstAllNodes(n ...*Node) {
	toVisit := []*Node{}

	for _, n := range n {
		fmt.Println(n.Level, n.Name)
		toVisit = append(toVisit, n.Children...)
	}

	if len(toVisit) == 0 {
		return
	}

	BreadthFirstAllNodes(toVisit...)
}
