.PHONY: run migrate

run:
	go run cmd/main.go

migrate:
	go run cmd/main.go migrate
