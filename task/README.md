Exercise #7: CLI Task Manager

## Requirements

https://github.com/gophercises/task

## TODOs

- [x] Build the CLI shell using lib from [this list](https://github.com/avelino/awesome-go#command-line). 
- [x] Add feat: add - adds a new task to our list
- [x] Add feat: list - lists all of our incomplete tasks
- [x] Add feat: do - marks a task as complete
- [x] Write the [BoltDB](https://github.com/boltdb/bolt) interactions
- [x] Add `rm` command to remove a task
- [ ] Add unit tests

## Setup

To run this project, you need to install [Go](https://go.dev/doc/install), an editor/IDE such as VSCode/Goland.

```bash
# Run these commands in the root folder
go mod tidy
# List all tasks, default -c=false
go run main.go list -c=true
# Add a task with name and description
go run main.go add "read a book" "read the first 30 pages"
# Accomplish a task
go run main.go do "09eb5e0b-8b7a-4527-b5b6-534bd4257840"
# Remove a task
go run main.go rm "09eb5e0b-8b7a-4527-b5b6-534bd4257840"
```
