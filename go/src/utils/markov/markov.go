// Persistent uses the local file system to model the markov
package markov

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"sync"
)

func NewModelBuilder(s ModelStore, workers int) ModelBuilder {
	mb := &modelBuilder{
		s:      s,
		jobs:   make(chan []string, workers),
		errs:   make(chan error, workers),
		closed: false,
		wg:     &sync.WaitGroup{},
	}

	// start worker pool
	for i := 0; i < workers; i++ {
		mb.wg.Add(1)
		go func() {
			defer mb.wg.Done()
			if err := worker(mb); err != nil {
				mb.errs <- err
			}
		}()
	}

	return mb
}

type modelBuilder struct {
	s      ModelStore
	jobs   chan []string
	errs   chan error
	wg     *sync.WaitGroup
	closed bool
}

func (mb *modelBuilder) Add(tokens []string) (err error) {
	if mb.closed {
		return fmt.Errorf("cannot call Add() method after model builder has been closed")
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("caught panic on adding...\n%s", r)
		}
	}()
	mb.jobs <- tokens
	return nil
}

func (mb *modelBuilder) Errs() []error {
	if mb.closed {
		return []error{fmt.Errorf("cannot call Errs() method after model builder has been closed")}
	}
	var errs []error
	for {
		select {
		case err := <-mb.errs:
			errs = append(errs, err)
		default:
			return errs
		}
	}
}

func (mb *modelBuilder) Close() []error {
	if mb.closed {
		return nil
	}
	mb.closed = true
	close(mb.jobs)
	close(mb.errs)
	errs := mb.Errs()
	mb.wg.Wait()
	return errs
}

func worker(pm *modelBuilder) error {
	for tokens := range pm.jobs {
		// If the set of input tokens is: [foo, bar, baz, bing]
		// store all the intermediate counts:
		// [foo, bar, baz]:{bing:1}
		// [foo, bar]:{baz:1}
		// [foo]: {baz:1}
		for i := len(tokens) - 1; i > 0; i-- {
			pathKey := makeKey(tokens[:i])
			// Lock this particular key such that we don't attempt
			// simultaneous reads/writes to that key in the store
			countKey := tokens[i]
			counts, err := pm.s.Get(pathKey)
			if err != nil {
				return fmt.Errorf("failed to get counts for %q...\n%s", pathKey, err)
			}
			counts[countKey] += 1

			if err := pm.s.Set(pathKey, counts); err != nil {
				return fmt.Errorf("failed to get set counts for %q:%q...\n%s", pathKey, countKey, err)
			}
		}
	}
	return nil
}

func NewModelExecer(s ModelStore) ModelExecer {
	return &modelExecer{s}
}

type modelExecer struct {
	s ModelStore
}

func (me *modelExecer) GetChoices(tokens []string) (Choices, error) {
	for i := 0; i < len(tokens); i++ {
		pathKey := makeKey(tokens[i:])
		counts, err := me.s.Get(pathKey)
		if err != nil {
			return nil, fmt.Errorf("failed to get counts for %q...\n%s", pathKey, err)
		}

		if len(counts) == 0 {
			// no word probable word choices, so try with fewer tokens
			continue
		} else {

			total := 0
			for k := range counts {
				total += counts[k]
			}

			c := make(Choices, len(counts))
			i := 0
			for k := range counts {
				c[i] = Choice{Prob: float64(counts[k]) / float64(total), Outcome: k}
				i++
			}

			return c, nil
		}
	}

	return nil, fmt.Errorf("there is no probable value for the token sequence %v", tokens)
}

func makeKey(tokens []string) string {
	// TODO(henry): Perhaps there is a more efficient way of doing this?
	return fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(tokens, ""))))
}
