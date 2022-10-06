all: migrate run

migrate:
	goose -dir db/migrations postgres "postgresql://postgres:postgres@127.0.0.1:5430/gophercises_phone?sslmode=disable" up

run:
	go run main.go

test:
	go test --cover ./...

start:
	docker compose up --build -d

stop:
	docker compose down
