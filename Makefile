all: deps test
	CGO_ENABLED=0 go build -a -o hello_world

test:
	GIN_MODE=debug go test

deps:
	go get -u
	go get github.com/stretchr/testify
