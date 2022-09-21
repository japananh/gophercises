Exercise #5: Sitemap Builder

## Requirements

https://github.com/gophercises/sitemap/blob/master/README.md

## TODO

- [x] Output `xml`
- [x] Handle cyclical case
- [x] Useful packages: `net/http`, `gophercises/link`, `encoding/xml`, `flag`
- [x] Add a `domain` flag
- [x] Add a `depth` flag using BFS (breath-first search algorithm)
- [ ] Write test

## Setup

To run this project, you need to install [Go](https://go.dev/doc/install), an editor/IDE such as VSCode/Goland.

```bash
# Run these command in the root folder
go mod tidy
go run main.go -depth=<your-depth> -url=<your-domain>
```