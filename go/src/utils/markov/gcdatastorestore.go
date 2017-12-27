package markov

import (
	"fmt"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

// mem is the in memory file store
type dataStore struct {
	client    *datastore.Client
	modelName string
}

const kind = "markov-model"

func NewDataStore(projectID string, modelName string) (ModelStore, error) {
	client, err := datastore.NewClient(context.Background(), projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create datastore client...\n%s", err)
	}
	return &dataStore{client, modelName}, nil
}

func (store *dataStore) Set(key string, counts map[string]int) error {
	dsk := datastore.NameKey("markov-model", key, nil)
	dsk.Namespace = store.modelName
	if _, err := store.client.Put(context.Background(), dsk, &counts); err != nil {
		return fmt.Errorf("failed to insert entitiy...\n%s", err)
	}
	return nil
}

func (store *dataStore) Get(key string) (map[string]int, error) {
	dsk := datastore.NameKey("markov-model", key, nil)
	dsk.Namespace = store.modelName
	dst := map[string]int{}
	if err := store.client.Get(context.Background(), dsk, &dst); err != nil {
		return nil, fmt.Errorf("failed to get entitiy...\n%s", err)
	}

	return dst, nil
}
