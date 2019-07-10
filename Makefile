BIN="ci-test"

build:
	go build -o $(BIN)

test:
	go test -v ./...