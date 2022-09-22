package task

import "fmt"

var todos []string

func Add(input string) bool {
	if input == "" {
		return false
	}
	if todos == nil {
		todos = []string{}
	}
	todos = append(todos, input)
	fmt.Println(todos)
	return true
}

func Resolve(input string) bool {
	if input == "" || todos == nil || len(todos) == 0 {
		return false
	}
	tasks := make([]string, 0)
	for _, t := range todos {
		if input != t {
			tasks = append(tasks, t)
		}
	}
	fmt.Println(todos)
	return true
}

func List() []string {
	return todos
}
