.PHONY: run migrate test

run:
	go run cmd/main.go

migrate:
	go run cmd/main.go migrate

test :
	go test -v ./test