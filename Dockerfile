# Copied from https://github.com/prometheus/client_golang/blob/master/examples/simple/Dockerfile

FROM golang:1.9.2 AS builder
WORKDIR /go/src/github.com/infrastructure-as-code/docker-hello-world
COPY hello-world.go .
RUN go get -d
RUN CGO_ENABLED=0 GOOS=linux go build -a hello-world.go

FROM scratch
LABEL maintainer "Vince Tse <thelazyenginerd@gmail.com>"
COPY --from=builder /go/src/github.com/infrastructure-as-code/docker-hello-world/hello-world .
EXPOSE 8080
ENTRYPOINT ["/hello-world"]
