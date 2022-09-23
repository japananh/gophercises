Exercise #7: CLI Task Manager

## Requirements

https://github.com/gophercises/task

## TODOs

- [x] Build the CLI shell using lib from [this list](https://github.com/avelino/awesome-go#command-line). 
- [x] Add feat: add - adds a new task to our list
- [x] Add feat: list - lists all of our incomplete tasks
- [x] Add feat: do - marks a task as complete
- [x] Write the [BoltDB](https://github.com/boltdb/bolt) interactions

## Setup

To run this project, you need to install [Go](https://go.dev/doc/install), an editor/IDE such as VSCode/Goland.

```bash
# Run these commands in the root folder
go mod tidy
# List all tasks 
go run main.go list
# Add a task
go run main.go add "read a book"
# Accomplish a task
go run main.go do "read a book"
```
