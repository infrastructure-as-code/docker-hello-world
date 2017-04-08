# infrastructureascode/hello-world

A Docker "Hello World" web server inspired by the [tutum/hello-world](https://hub.docker.com/r/tutum/hello-world/) image which seems to have become unmaintained, and returns HTTP/400 more often than it should.

## Features

1. Always returns a HTTP 200 status code and a "Hello, world!" message.
1. Has a health check endpoint, `/health`, that returns an empty response.

## Building

```
docker build --rm infrastructureascode/hello-world .
```

## Usage

```
# start the container
docker run \
  --detach \
  --name hello-world \
  --publish 8000:80 \
  infrastructureascode/hello-world

# curl the container
curl http://0.0.0.0:8000/

# curl the health check endpoint which returns an empty response
curl http://0.0.0.0:8000/health
```
