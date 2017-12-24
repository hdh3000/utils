package markov

import (
	"testing"
	"fmt"
)

func TestBuilder_Add(t *testing.T) {
	inter, _ := NewModel()
	b := inter.(*model)

	testCases := []struct {
		tokens []string
		apply int
	} {
		{[]string{"cheese", "pizza", "sauce"}, 5},
		{[]string{"cheese", "pizza", "peppers"}, 2},
		{[]string{"pizza", "pizza", "peppers"}, 6},
	}

	for _, c := range testCases {
		for i :=0; i < c.apply; i++ {
			b.Add(c.tokens)
		}

	}

	for _, c := range testCases {
		got := b.Root.Children[c.tokens[0]].Children[c.tokens[1]].Count[c.tokens[2]]
		if got != c.apply {
			t.Errorf("incorrect word count wanted %d, got %d", got, c.apply)
		}
	}
}

func TestShallowBuilder_Add(t *testing.T) {
	inter, _ := NewModel()
	b := inter.(*model)

	testCases := []struct {
		tokens []string
		apply int
	} {
		{[]string{"cheese", "pizza"}, 5},
		{[]string{"cheese", "sauce",}, 2},
		{[]string{"pizza", "pizza", }, 6},
	}

	for _, c := range testCases {
		for i :=0; i < c.apply; i++ {
			b.Add(c.tokens)
		}

	}

	for _, c := range testCases {
		got := b.Root.Children[c.tokens[0]].Count[c.tokens[1]]
		if got != c.apply {
			t.Errorf("incorrect word count wanted %d, got %d", got, c.apply)
		}
	}
}

func TestModel_GetChoices(t *testing.T) {
	inter, _  := NewModel()
	b := inter.(*model)

	testCases := []struct {
		tokens []string
		apply int
	} {
		{[]string{"cheese", "pizza", "peppers"}, 10},
		{[]string{"cheese", "pizza", "sauce"}, 20},
	}

	for _, c := range testCases {
		for i :=0; i < c.apply; i++ {
			b.Add(c.tokens)
		}

	}

	// TODO(henry): make this a real test
	fmt.Println(b.GetChoices([]string{"cheese", "pizza"}))
}