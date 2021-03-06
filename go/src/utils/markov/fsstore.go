package markov

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// fs is the FileSystemStore.
type fs struct {
	dir string
}

// new FSStore creates a store that uses your filesystem to create a persistent model.
func NewFSStore(dir string) (ModelStore, error) {
	d, err := os.Stat(dir)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to find dir...\n%s", err)
	}

	if os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0740); err != nil {
			return nil, fmt.Errorf("failed to create dir...\n%s", err)
		}
	} else if !d.IsDir() {

		return nil, fmt.Errorf("%s exists and is not a dir", dir)
	}

	return &fs{dir}, nil

}

func (fs *fs) Set(key string, counts map[string]int) error {
	f, err := os.Create(path.Join(fs.dir, key))
	if err != nil {
		return fmt.Errorf("failed to create file...\n%s", err)
	}
	defer f.Close()

	b, err := json.Marshal(counts)
	if err != nil {
		return fmt.Errorf("failed to marshal counts...\n%s", err)
	}

	if _, err = f.Write(b); err != nil {
		return fmt.Errorf("failed to write counts...\n%s", err)
	}

	return nil
}

func (fs *fs) Get(key string) (map[string]int, error) {
	f, err := os.Open(path.Join(fs.dir, key))
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to open file...\n%s", err)
	}
	defer f.Close()
	if os.IsNotExist(err) {
		return map[string]int{}, nil
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read counts...\n%s", err)
	}

	var counts = make(map[string]int)
	if err := json.Unmarshal(b, &counts); err != nil {
		return nil, fmt.Errorf("failed to unmarshal counts...\n%s", err)
	}

	return counts, nil
}
