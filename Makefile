.PHONY: run build test

run:
	go run main.go

build:
	go build -o fpt-clique main.go
