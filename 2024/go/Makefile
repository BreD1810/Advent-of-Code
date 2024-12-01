.PHONY: build run test

build:
	go build -o aoc-2024 cmd/aoc-2024/main.go

run: build
	./aoc-2024 $(DAY)

test:
	go test -count=1 ./...
