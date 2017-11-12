# Copied from https://github.com/prometheus/client_golang/blob/master/examples/simple/Dockerfile

FROM golang:1.9.2 AS builder
WORKDIR /go/src/github.com/infrastructure-as-code/docker-hello-world
ENV GIN_MODE debug
COPY Makefile *.go ./
RUN make all

FROM scratch
LABEL maintainer "Vince Tse <thelazyenginerd@gmail.com>"
COPY --from=builder /go/src/github.com/infrastructure-as-code/docker-hello-world/hello_world .
ENV GIN_MODE release
EXPOSE 8080
ENTRYPOINT ["/hello_world"]
