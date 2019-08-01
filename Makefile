all: build

build:
	go build -o sp cmd/sp/main.go

test:
	go test ./... -v -cover