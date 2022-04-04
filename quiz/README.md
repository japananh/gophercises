# Exercise #1: Quiz Game

## Requirements

https://github.com/gophercises/quiz/blob/master/README.md

## TODO

- [x] Create `problem.csv` file
- [x] Custom csv filename via flag
- [x] Output total numbers of correct questions and how many questions in total
- [x] Add timer (30 seconds) via a flag
- [x] Filter invalid answers/questions
- [x] Print answer to the screen after user press Enter
- [ ] Continue loop when user press Enter
- [ ] Write test

## Setup

To run this project, you need to install [Go](https://go.dev/doc/install), an editor/IDE such as VSCode/Goland.

```bash
# Run these command in the root folder
go mod tidy
go run main.go -path=./problem.csv -time=30
```
