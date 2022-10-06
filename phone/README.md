# Exercise #8: Phone Number Normalizer

## Requirements

https://github.com/gophercises/phone/blob/master/README.md

## TODO

- [x] Setup DB
- [x] Normalize phone number in DB (format: `##########`)
- [x] Remove duplicate data
- [ ] Write unit test

## Setup

I used Postgres database. To run migration, you need to install [goose](https://github.com/pressly/goose)

```bash
# Create a new migration
cd phone
goose -dir db/migrations postgres "postgresql://postgres:postgres@127.0.0.1:5430/gophercises_phone?sslmode=disable" create init sql

# Run migration
goose -dir db/migrations postgres "postgresql://postgres:postgres@127.0.0.1:5430/gophercises_phone?sslmode=disable" up
```

To run this project, you need to install [Go](https://go.dev/doc/install), an editor/IDE such as VSCode/Goland.

```bash
# Run these command in the root folder
go mod tidy
go run main.go
```
