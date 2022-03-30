# Exercise #3: Choose your own adventure

## Requirements

https://github.com/gophercises/cyoa

## TODO

- [x] Create `story.json` file.
- [x] Use the `html/template` package to create your HTML pages.
- [x] Create an `http.Handler` to handle the web requests instead of a handler function.
- [x] Use the `encoding/json` package to decode the JSON file.
- [ ] Create a command-line version of `cyoa`
- [ ] Consider how you would alter your program in order to support stories starting form a story-defined arc. That is, what if all stories didn't start on an arc named intro? How would you redesign your program or restructure the JSON? This bonus exercises is meant to be as much of a thought exercise as an actual coding one.
- [ ] Write test

## Setup

To run this project, you need to install [Go](https://go.dev/doc/install), an editor/IDE such as VSCode/Goland.

```bash
go mod tidy
go run main.go
```
