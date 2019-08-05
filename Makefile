sp: build

all: clean lint test build

clean:
	rm -f sp

build:
	go build -o sp cmd/sp/main.go

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run

test:
	go test -v -cover -race ./...
