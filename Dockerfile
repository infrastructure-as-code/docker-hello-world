FROM ghcr.io/vincetse/scratch
LABEL maintainer "Vince Tse <vincetse@users.noreply.github.com>"
COPY ./hello_world .
ENV GIN_MODE release
EXPOSE 8080
ENTRYPOINT ["/hello_world"]
