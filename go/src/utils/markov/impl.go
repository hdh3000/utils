package markov
//
//import (
//	"bufio"
//)
//
//var _ Model = &model{}
//
//func NewModel() (Model) {
//	return &model{
//		Root: &node{
//			Children:map[string]*node{},
//			Count:map[string]int{},
//		},
//	}
//}
//
//type model struct {
//	Root  *node
//	scanner bufio.Scanner
//}
//
//type node struct {
//	Children map[string]*node
//	Count map[string]int
//}
//
//func (m *model) Add(tokens []string) error {
//	increment(tokens, m.Root)
//	return nil
//}
//
//func increment(tokens []string, parent *node) {
//	if len(tokens) == 0 {
//		return
//	}
//
//	childKey := tokens[0]
//
//	// Get the next node
//	child, ok := parent.Children[childKey]
//	if !ok {
//		// Need to make sure we have a node to work on.
//		child = &node{
//			Children: make(map[string]*node),
//			Count: make(map[string]int),
//		}
//		parent.Children[childKey] = child
//	}
//
//	parent.Count[childKey] += 1
//
//	increment(tokens[1:], parent.Children[childKey])
//}
//
//func find(tokens []string, parent *node) (map[string]int) {
//	if len(tokens) == 1 {
//		return parent.Count
//	}
//
//	// Get the next node
//	child, ok := parent.Children[tokens[0]]
//	if !ok {
//		// Need to make sure we have a node to work on.
//		child = &node{
//			Children: map[string]*node{},
//			Count: map[string]int{},
//		}
//		parent.Children[tokens[0]] = child
//	}
//
//
//	return find(tokens[1:], child)
//}
//
//
//func (m *model) GetChoices(tokens []string) (Choices, error) {
//	probs := find(append(tokens, ""), m.Root)
//	if len(probs) == 0 {
//		return m.GetChoices(tokens[1:])
//	}
//
//	total := 0
//	for k  := range probs {
//		total += probs[k]
//	}
//
//	c := make(Choices, len(probs))
//	i := 0
//
//	for k := range probs {
//		c[i] = Choice{Prob:float64(probs[k])/float64(total), Outcome: k}
//		i++
//	}
//
//	return c, nil
//}
//
//
//
//
