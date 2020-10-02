FROM alpine AS qemu
ENV QEMU_URL https://github.com/balena-io/qemu/releases/download/v4.0.0%2Bbalena2/qemu-4.0.0.balena2-arm.tar.gz
RUN apk add curl && curl -sL ${QEMU_URL} | tar zxvf - -C . --strip-components 1

FROM arm32v6/golang:1.15-alpine AS builder
COPY --from=qemu qemu-arm-static /usr/bin
WORKDIR /go/src/github.com/infrastructure-as-code/docker-hello-world
ENV GIN_MODE debug
COPY Makefile *.go ./
RUN apk update
RUN	apk upgrade
RUN	apk add alpine-sdk
RUN make all

FROM infrastructureascode/scratch
LABEL maintainer "Vince Tse <vincetse@users.noreply.github.com>"
COPY --from=builder /go/src/github.com/infrastructure-as-code/docker-hello-world/hello_world .
ENV GIN_MODE release
EXPOSE 8080
ENTRYPOINT ["/hello_world"]
