BIN="ci-test"

build:
	go build -o $(BIN)

test:
	go test -v ./...

test-cover:
	go test -v -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html