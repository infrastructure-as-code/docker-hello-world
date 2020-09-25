FROM alpine AS qemu
ENV QEMU_URL https://github.com/balena-io/qemu/releases/download/v3.0.0%2Bresin/qemu-3.0.0+resin-arm.tar.gz
RUN apk add curl && curl -L ${QEMU_URL} | tar zxvf - -C . --strip-components 1

FROM arm32v7/golang:1.15-alpine AS builder
COPY --from=qemu qemu-arm-static /usr/bin
WORKDIR /go/src/github.com/infrastructure-as-code/docker-hello-world
ENV GIN_MODE debug
COPY Makefile *.go ./
RUN apk update
RUN	apk upgrade
RUN	apk add alpine-sdk
RUN make all

FROM arm32v7/alpine:latest
LABEL maintainer "Vince Tse <vincetse@users.noreply.github.com>"
COPY --from=builder /go/src/github.com/infrastructure-as-code/docker-hello-world/hello_world .
ENV GIN_MODE release
EXPOSE 8080
ENTRYPOINT ["/hello_world"]
