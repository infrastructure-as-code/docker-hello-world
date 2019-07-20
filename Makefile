all: deps lint test
	CGO_ENABLED=0 go build -a -o hello_world

test:
	GIN_MODE=debug go test

lint:
	test -z `gofmt -s -l .`
	go vet ./...
	golint -set_exit_status `go list ./...`

deps:
	go get -d
	go get github.com/stretchr/testify
	go get golang.org/x/lint/golint
