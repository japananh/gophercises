package task

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
	"time"
)

type Task struct {
	IsCompleted bool   `json:"is_done"`
	Id          string `json:"id"`
	Title       string `json:"name"`
	Desc        string `json:"desc"`
	CreatedDate int64  `json:"created_date"`
}

type Filter struct {
	IsCompleted bool `json:"is_completed,omitempty"`
}

func Add(title string, desc string) error {
	if title == "" || desc == "" {
		return fmt.Errorf("title or description couldn't be empty")
	}

	dupTask, err := Get(title)
	if err != nil {
		return fmt.Errorf("could not add task: %v", err)
	}
	if dupTask.Title == title {
		return fmt.Errorf("title cannot be duplicated")
	}

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

	fmt.Println("Successfully added:", title)

	return nil
}

func Remove(title string) error {
	return db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte(TodoBucketName)).Delete([]byte(title))
		if err != nil {
			return fmt.Errorf("could not delete task: %v", err)
		}
		fmt.Println("Successfully deleted:", title)
		return nil
	})
}

func Resolve(id string) error {
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
		fmt.Println("Successfully deleted:", id)
		return nil
	})
}

func Get(id string) (task *Task, err error) {
	if id == "" {
		return nil, fmt.Errorf("title shouldn't be empty")
	}

	var found Task

	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte(TodoBucketName)).Get([]byte(id))
		if err = json.Unmarshal(b, &found); err != nil {
			return err
		}
		task = &found
		return nil
	}); err != nil {
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
