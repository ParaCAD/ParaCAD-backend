package sqldb

import (
	"fmt"
	"time"
)

type CacheItem struct {
	Key     string    `db:"key"`
	Value   []byte    `db:"value"`
	Created time.Time `db:"created"`
}

func (db *SQLDB) CacheGetModel(key string) ([]byte, error) {
	fmt.Println("CacheGetModel called, but cache is not implemented yet. " + key)
	// TODO: Implement cache
	return nil, nil
}

func (db *SQLDB) CacheSetModel(key string, model []byte) error {
	fmt.Println("CacheSetModel called, but cache is not implemented yet. " + key)
	return nil
}

func (db *SQLDB) cacheCleanerJob() {
	db.clearCache()

	ticker := time.NewTicker(24 * time.Hour)
	for range ticker.C {
		db.clearCache()
	}
}

func (db *SQLDB) clearCache() {
}
