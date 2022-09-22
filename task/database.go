package task

import (
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

func init() {
	// Open my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *bolt.DB) {
		_ = db.Close()
	}(db)
}
