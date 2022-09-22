package task

import (
	"fmt"
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

const TodoBucketName = "TODOS"

func InitDatabase() error {
	// Open my.db data file in your current directory.
	// It will be created if it doesn't exist.
	var err error
	if db, err = bolt.Open("test.db", 0600, &bolt.Options{Timeout: 3 * time.Second}); err != nil {
		log.Fatal(err)
	}
	if err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte(TodoBucketName))
		if err != nil {
			return fmt.Errorf("could not create todos bucket: %v", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done", db)
	return nil
}
