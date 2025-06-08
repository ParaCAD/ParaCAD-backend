package sqldb

import (
	"database/sql"
	"log"
	"time"
)

type CacheItem struct {
	Key     string    `db:"key"`
	Value   []byte    `db:"value"`
	Created time.Time `db:"created"`
}

func (db *SQLDB) CacheGetModel(key string) ([]byte, error) {
	var cachedModel CacheItem
	query := `
	SELECT key, value, created
	FROM cache
	WHERE key = $1
	`
	err := db.db.Get(&cachedModel, query, key)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return cachedModel.Value, nil
}

func (db *SQLDB) CacheSetModel(key string, model []byte) error {
	query := `
	INSERT INTO cache 
	(key, value, created)
	VALUES
	($1, $2, $3)
	`

	_, err := db.db.Exec(query, key, model, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (db *SQLDB) cacheCleanerJob() {
	err := db.clearCache()
	if err != nil {
		log.Printf("Error clearing cache: %s", err.Error())
	}

	ticker := time.NewTicker(24 * time.Hour)
	for range ticker.C {
		err = db.clearCache()
		if err != nil {
			log.Printf("Error clearing cache: %s", err.Error())
		}
	}
}

func (db *SQLDB) clearCache() error {
	query := `
	DELETE FROM cache
	WHERE created < $1
	`

	expirationTime := time.Now().Add(-24 * time.Hour)

	_, err := db.db.Exec(query, expirationTime)
	if err != nil {
		return err
	}

	return nil
}
