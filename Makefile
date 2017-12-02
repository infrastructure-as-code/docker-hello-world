all: deps test
	CGO_ENABLED=0 go build -a -o hello_world
	upx --brute hello_world

test:
	GIN_MODE=debug go test

deps:
	go get -d
	go get github.com/stretchr/testify
