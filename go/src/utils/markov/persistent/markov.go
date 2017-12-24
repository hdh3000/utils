// Persistent uses the local file system to model the markov
package persistent

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"utils/markov"
	"sync"
)

type Store interface {
	Get(key string) (map[string]int, error)
	Set(key string, counts map[string]int) error
}

func NewModelBuilder(s Store, workers int) markov.ModelBuilder {
 	pm := &persitentModel{
		s:s,
		locks: map[string]*sync.Mutex{},
		Mutex: &sync.Mutex{},
		jobs: make(chan struct {tokens []string}, workers),
		errs: make(chan error, workers),
	}

	// start workers
	for i := 0; i < workers; i++ {
		go worker(pm)
	}
	return pm
}

func NewModelExecer(s Store, numWorkers int) markov.ModelExecer {

}

type persitentModel struct {
	s Store
	locks map[string]*sync.Mutex
	*sync.Mutex
	jobs chan struct{tokens []string}
	errs chan error
	wg *sync.WaitGroup
}

func (pm *persitentModel) Add(tokens []string) error {
	for i := len(tokens) - 1; i > 0; i-- {
		pathKey := makeKey(tokens[:i])
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
	return nil
}



func worker(pm *persitentModel) {
	// If the set of input tokens is: [foo, bar, baz, bing]
	// store all the intermediate counts:
	// [foo, bar, baz]:{bing:1}
	// [foo, bar]:{baz:1}
	// [foo]: {baz:1}

	for j := range pm.jobs {
		pm.wg.Add(1)
		defer pm.wg.Done()

		tokens := j.tokens
		for i := len(tokens) - 1; i > 0; i-- {
			pathKey := makeKey(tokens[:i])

			pm.Lock() // acquire the lock to mutate the map of mutexes
			if _, ok := pm.locks[pathKey]; !ok {
				pm.locks[pathKey] = &sync.Mutex{}
			}
			pm.Unlock()


			pm.locks[pathKey].Lock()
			defer pm.locks[pathKey].Unlock()

			countKey := tokens[i]
			counts, err := pm.s.Get(pathKey)
			if err != nil {
				pm.errs <- fmt.Errorf("failed to get counts for %q...\n%s", pathKey, err)
			}
			counts[countKey] += 1
			if err := pm.s.Set(pathKey, counts); err != nil {
				pm.errs <- fmt.Errorf("failed to get set counts for %q:%q...\n%s", pathKey, countKey, err)
			}
		}
	}
}

func (pm *persitentModel) GetChoices(tokens []string) (markov.Choices, error) {
	for i := 0; i < len(tokens); i++ {
		pathKey := makeKey(tokens[i:])
		counts, err := pm.s.Get(pathKey)
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

			c := make(markov.Choices, len(counts))
			i := 0
			for k := range counts {
				c[i] = markov.Choice{Prob: float64(counts[k]) / float64(total), Outcome: k}
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
