BINARY_NAME=marky
.DEFAULT_GOAL := build

build:
	GOARCH=arm64 GOOS=darwin go build -o ./target/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ./target/${BINARY_NAME}-windows main.go

clean:
	go clean
	rm  ./target/${BINARY_NAME}-darwin
	rm  ./target/${BINARY_NAME}-linux
	rm  ./target/${BINARY_NAME}-windows
