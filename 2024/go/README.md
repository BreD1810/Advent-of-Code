Advent of Code - 2024 (Golang)
===
To run a day use:
```shell
make run DAY=1
# Alternatively with Go
go run cmd/aoc-2024/main.go 1
# Alternatively with the binary
make build
./aoc-2024 1
```

## Tests

To run all tests:
```shell
make test
```

To run specific tests:
```shell
 go test -count=1 -run day1 ./...
```
