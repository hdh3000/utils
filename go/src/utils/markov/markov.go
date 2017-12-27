// Persistent uses the local file system to model the markov
package markov

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"sync"
)

type job struct {
	id       string
	countKey string
}

func NewModelBuilder(s ModelStore, workers int) ModelBuilder {
	mb := &modelBuilder{
		s:       s,
		jobChan: make([]chan job, workers),
		errChan: make(chan error, workers),
		closed:  false,
		workers: workers,
		wg:      &sync.WaitGroup{},
	}

	// start worker pool
	for i := 0; i < workers; i++ {
		mb.jobChan[i] = make(chan job) // ensure that we have a non-nil chan.
		j := i                         // hold the index in j, as i will change by the time we start the worker
		go func() {
			mb.wg.Add(1)
			defer mb.wg.Done()
			if err := worker(mb, j); err != nil {
				mb.errChan <- fmt.Errorf("failed to execute worker...\n%s", err)
			}
		}()
	}
	return mb
}

type modelBuilder struct {
	s       ModelStore
	jobChan []chan job
	errChan chan error
	wg      *sync.WaitGroup
	workers int
	closed  bool
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

	for i := len(tokens) - 1; i > 0; i-- {
		// If the set of input tokens is: [foo, bar, baz, bing]
		// store all the intermediate counts:
		// [foo, bar, baz]:{bing:1}
		// [foo, bar]:{baz:1}
		// [foo]: {baz:1}
		pathKey, keySum := makeKey(tokens[:i])
		c := keySum % (mb.workers - 1)
		mb.jobChan[c] <- job{id: pathKey, countKey: tokens[i]}
	}

	return nil
}

func (mb *modelBuilder) errs() []error {
	if mb.closed {
		return []error{fmt.Errorf("cannot call errs() method after model builder has been closed")}
	}
	var errs []error
	for err := range mb.errChan {
		errs = append(errs, err)
	}

	return errs
}

func (mb *modelBuilder) Close() []error {
	if mb.closed {
		return nil
	}

	for i, _ := range mb.jobChan {
		close(mb.jobChan[i])
	}
	mb.wg.Wait()      // wait for all the workers to finish
	close(mb.errChan) // close errors
	errs := mb.errs() // read off the error chan.
	mb.closed = true
	return errs
}

func worker(pm *modelBuilder, number int) error {
	for tokens := range pm.jobChan[number] {
		counts, err := pm.s.Get(tokens.id)
		if err != nil {
			return fmt.Errorf("failed get count for %s...\n%s", tokens.id, err)
		}
		counts[tokens.countKey] += 1

		if err := pm.s.Set(tokens.id, counts); err != nil {
			return fmt.Errorf("failed to get set counts for %q:%q...\n%s", tokens.id, tokens.countKey, err)
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
		pathKey, _ := makeKey(tokens[i:])
		counts, err := me.s.Get(pathKey)
		if err != nil {
			return nil, fmt.Errorf("failed to get counts for %q...\n%s", pathKey, err)
		}
		if counts != nil {
			return counts, nil
		}
	}
	return nil, fmt.Errorf("there is no probable value for the token sequence %v", tokens)
}

func makeKey(tokens []string) (string, int) {
	// TODO(henry): Perhaps there is a more efficient way of doing this?
	sum := sha256.Sum256([]byte(strings.Join(tokens, "")))
	total := 0
	for _, b := range sum {
		total += int(b)
	}

	return fmt.Sprintf("%x", sum), total
}
