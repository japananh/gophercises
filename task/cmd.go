package task

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
	"time"
)

type Task struct {
	IsCompleted bool   `json:"is_completed"`
	Id          string `json:"id"`
	Title       string `json:"name"`
	Desc        string `json:"desc"`
	CreatedDate int64  `json:"created_date"`
}

type Filter struct {
	IsCompleted bool `json:"is_completed"`
}

func Add(title string, desc string) error {
	// Validate data
	if title == "" || desc == "" {
		return fmt.Errorf("title or description couldn't be empty")
	}

	// Find duplicated task
	dupTask, err := Get(title)
	if err != nil {
		return fmt.Errorf("could not add task: %v", err)
	}
	if dupTask != nil && dupTask.Title == title {
		return fmt.Errorf("title cannot be duplicated")
	}

	// Prepare data
	id := uuid.New().String()
	task := Task{
		Id:          id,
		IsCompleted: false,
		Title:       title,
		Desc:        desc,
		CreatedDate: time.Now().UnixNano(),
	}
	taskByte, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("could not marshal json: %v", err)
	}

	// Create task
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte(TodoBucketName)).Put([]byte(id), taskByte)
		if err != nil {
			return fmt.Errorf("could not insert entry: %v", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not add task: %v", err)
	}

	fmt.Println("Successfully added:", title, id)

	return nil
}

func Remove(id string) error {
	return db.Update(func(tx *bolt.Tx) error {
		if err := tx.Bucket([]byte("DB")).Bucket([]byte(TodoBucketName)).Delete([]byte(id)); err != nil {
			return fmt.Errorf("could not delete task %s: %v", id, err)
		}
		fmt.Println("Successfully deleted task id:", id)
		return nil
	})
}

func Resolve(id string) error {
	// Check for existed task
	found, err := Get(id)
	if err != nil {
		return fmt.Errorf("task not found")
	}

	found.IsCompleted = true

	foundByte, err := json.Marshal(found)
	if err != nil {
		return fmt.Errorf("could not resolve task: %v", err)
	}

	// Find task id
	return db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte(TodoBucketName)).Put([]byte(found.Id), foundByte)
		if err != nil {
			return fmt.Errorf("could not delete task: %v", err)
		}
		fmt.Println("Successfully resolved task id:", id)
		return nil
	})
}

func Get(id string) (task *Task, err error) {
	if id == "" {
		return nil, fmt.Errorf("task id shouldn't be empty")
	}

	var found Task
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte(TodoBucketName)).Get([]byte(id))
		if len(b) == 0 {
			return nil
		}
		if err = json.Unmarshal(b, &found); err != nil {
			return err
		}
		task = &found
		return nil
	})
	if err != nil {
		return nil, err
	}

	return
}

func List(filter *Filter) error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte(TodoBucketName))
		err := b.ForEach(func(k, v []byte) error {
			var task Task
			if err := json.Unmarshal(v, &task); err != nil {
				return fmt.Errorf("could not unmarshal task: %v", err)
			}

			if task.IsCompleted == filter.IsCompleted {
				fmt.Printf("- key: %s\n  value: %v\n", string(k), task)
			}

			return nil
		})
		if err != nil {
			return fmt.Errorf("could not list task: %v", err)
		}
		return nil
	})
}
