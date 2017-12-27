// Package markov provides tooling to build and "run" markov models from arbitrary text.
package markov

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

// ModelBuilder builds a markov model
// Once closed, none of its methods can be called.
type ModelBuilder interface {
	// Add adds new choices to the model
	Add(tokens []string) error
	// Close the builder
	Close() []error
}

// ModelExecer executes a markov model
type ModelExecer interface {
	// Get choices gives you your next Choice given your current prefixes
	GetChoices(tokens []string) (Choices, error)
}

// A ModelStore stores the learned parameters.
// It must be safe for concurrent reads/writes.
// values for keys are guaranteed to be written in serial.
// That is key "abc", will always be Get/Set from the same thread.
type ModelStore interface {
	Get(key string) (map[string]int, error)
	Set(key string, counts map[string]int) error
}

type Chooser interface {
	// Choose picks a choice
	Choose(Choices) string
}

// Choices represents the next possible state of the model.
type Choices map[string]int

type Choice struct {
	// The chance that this outcome would happen
	Prob float64
	// The outcome that the probability describes
	Outcome string
}

// NewRandomChooser returns a chooser that chooses the next token at pseudo-random
// weighted by the probability of the token being next
func NewRandomChooser() Chooser {
	return &randomChooser{rand: rand.New(rand.NewSource(time.Now().Unix()))}
}

type randomChooser struct {
	rand *rand.Rand
}

func (m *randomChooser) Choose(c Choices) string {
	total := 0.0
	for _, v := range c {
		total += float64(v)
	}

	random := m.rand.Float64() * total
	sum := float64(0)
	var last string
	for k, v := range c {
		last = k
		if random <= sum {
			return k
		}
		sum += float64(v)
	}
	return last
}

// Use Ingest to bring in the entire contents of a reader.
// Must check for errors using m.Close()
func Ingest(m ModelBuilder, r io.Reader, splitter bufio.SplitFunc, chainLength int) error {
	s := bufio.NewScanner(r)
	s.Split(splitter)
	var tokens []string
	for s.Scan() {
		text := strings.ToLower(s.Text())

		if len(tokens) <= chainLength {
			tokens = append(tokens, text)
			continue
		}

		if err := m.Add(tokens); err != nil {
			return fmt.Errorf("failed to add tokens...\n%s", err)
		}

		tokens = tokens[1:]
		tokens = append(tokens, text)

		if errs := s.Err(); errs != nil {
			return fmt.Errorf("failed to scan tokens...\n%s", errs)
		}
	}
	return nil
}

type printFunc func(token string, pos int) string

func Exec(m ModelExecer, c Chooser, seed []string, chainLength int, printFunc printFunc, msgLength int) (io.Reader, error) {
	previous := seed
	for i := len(seed); i < chainLength; i++ {
		choices, err := m.GetChoices(previous)
		if err != nil {
			return nil, fmt.Errorf("failed to get choices...\n%s", err)
		}
		next := c.Choose(choices)
		previous = append(previous, next)
	}

	buf := &bytes.Buffer{}
	for i := 0; i < msgLength; i++ {
		choices, err := m.GetChoices(previous)
		if err != nil {
			return nil, fmt.Errorf("failed to get choices...\n%s", err)
		}
		next := c.Choose(choices)
		previous = append(previous[1:], next)

		buf.WriteString(printFunc(next, i))
	}

	return buf, nil
}
