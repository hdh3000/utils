// Package markov provides tooling to build and "run" markov models from arbitrary text.
package markov

import (
	"math/rand"
	"time"
)

// ModelBuilder builds a markov model
type ModelBuilder interface {
	// Add adds new choices to the model
	Add(tokens []string)
	// Get any errors that have occurred while adding
	Errs() []error
	// Close the builder
	Close() []error
}

// ModelExecr executes a markov model
type ModelExecer interface {
	// Get choices gives you your next Choice given your current prefixes
	GetChoices(tokens []string) (Choices, error)
}

type Chooser interface {
	// Choose picks a choice
	Choose(Choices) string
}

// Choices represents the next possible state of the model.
type Choices []Choice

type Choice struct{
	// The chance that this outcome would happen
	Prob float64
	// The outcome that the probability describes
	Outcome string
}


// NewRandomChooser returns a chooser that chooses the next token at pseudo-random
// weighted by the probability of the token being next
func NewRandomChooser() Chooser {
	return &randomChooser{rand : rand.New(rand.NewSource(time.Now().Unix()))}
}

type randomChooser struct {
	rand *rand.Rand
}

func (m *randomChooser) Choose(c Choices) string {
	random := m.rand.Float64()
	sum := float64(0)
	for i := range c {
		if random <= sum {
			return c[i].Outcome
		}
		sum += c[i].Prob
	}

	if len(c) > 0 {
		return c[len(c)-1].Outcome
	}

	return ""
}




