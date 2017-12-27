package markov

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// mem is the in memory file store
type sqlStore struct {
	db        *sql.DB
	modelName string
}

func NeSQLStore(db *sql.DB, modelName string) (ModelStore, error) {
	if _, err := db.Exec("DROP TABLE IF EXISTS ?", modelName); err != nil {
		return nil, fmt.Errorf("failed to drop table %s...\n%s", modelName, err)
	}

	if _, err := db.Exec("CREATE TABLE ? (id CHAR(64) PRIMARY KEY, counts VARCHAR)", modelName); err != nil {
		return nil, fmt.Errorf("failed to create table %s...\n%s", modelName, err)
	}

	return &sqlStore{db, modelName}, nil
}

func (store *sqlStore) Set(key string, counts map[string]int) error {
	b, err := json.Marshal(counts)
	if err != nil {
		return fmt.Errorf("failed to marshal counts...\n%s", err)
	}

	if _, err := store.db.Exec("INSERT INTO ? (id, counts) VALUES (?, ?)", store.modelName, key, string(b)); err != nil {
		return fmt.Errorf("failed to insert counts...\n%s", err)
	}
	return nil
}

func (store *sqlStore) Get(key string) (map[string]int, error) {
	res, err := store.db.Query("SELECT counts FROM ? WHERE id=?", store.modelName, key)
	if err != nil {
		return nil, fmt.Errorf("failed to query counts...\n%s", err)
	}

	if res.Next() {
		var raw string
		if err := res.Scan(&raw); err != nil {
			return nil, fmt.Errorf("failed to scan counts...\n%s", err)
		}

		out := map[string]int{}
		if err := json.Unmarshal([]byte(raw), &out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal counts...\n%s", err)
		}
		return out, nil
	}

	if res.Err() != nil {
		return nil, fmt.Errorf("failed to prepare counts...\n%s", err)
	}

	return make(map[string]int), nil
}
