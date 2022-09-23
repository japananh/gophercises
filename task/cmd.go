package task

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
)

func Add(input string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte(TodoBucketName)).Put([]byte(input), []byte(input))
		if err != nil {
			return fmt.Errorf("could not insert entry: %v", err)
		}
		return nil
	})
	fmt.Println("Successfully added:", input)
	return err
}

func Resolve(input string) error {
	return db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte(TodoBucketName)).Delete([]byte(input))
		if err != nil {
			return fmt.Errorf("could not delete task: %v", err)
		}
		fmt.Println("Successfully deleted:", input)
		return nil
	})
}

func List() error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("TODOS"))
		err := b.ForEach(func(k, v []byte) error {
			fmt.Printf("- key: %s\n  value: %s\n", string(k), string(v))
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
}
