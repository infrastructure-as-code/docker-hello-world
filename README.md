[![GHCR Build Status](https://github.com/infrastructure-as-code/docker-hello-world/actions/workflows/ghcr.yml/badge.svg?branch=master)](https://github.com/infrastructure-as-code/docker-hello-world/actions/workflows/ghcr.yml)
[![Docker Hub Build Status](https://github.com/infrastructure-as-code/docker-hello-world/actions/workflows/dockerhub.yml/badge.svg?branch=master)](https://github.com/infrastructure-as-code/docker-hello-world/actions/workflows/dockerhub.yml)


# ghcr.io/infrastructure-as-code/hello-world

A [Prometheus](https://prometheus.io/)-instrumented Docker "Hello World" web server.  This image began life as [infrastructureascode/hello-world](https://hub.docker.com/r/infrastructureascode/hello-world) on Docker Hub, and is now also available on the GitHub Container Registry as `ghcr.io/infrastructure-as-code/hello-world`.


## Images

| Registry Name | Image Name |
|---------------|------------|
| GitHub Container Registry | `ghcr.io/infrastructure-as-code/hello-world` |
| Docker Hub | `infrastructureascode/hello-world` |


## Features

1. Always returns a HTTP 200 status code and a "Hello, World!" message at the `/` path.
1. Has a metrics endpoint at `/metrics` that returns Prometheus metrics.
1. Has a health check endpoint, `/health`, that returns an empty response and a HTTP 200 response.

## Building

```
docker build --rm -t ghcr.io/infrastructure-as-code/hello-world .
```

## Releases

Images are [built with GitHub Actions](https://github.com/infrastructure-as-code/docker-hello-world/actions/workflows/build-images.yml) upon tagging/push, and pushed directly to the GitHub Container Registry.  You can look at the [packages page](https://github.com/infrastructure-as-code/docker-hello-world/pkgs/container/hello-world) for the latest tagged version.

Additionally, the `sha1sum` of the binary in each image is emitted during the build process (look for "Show binary info") in the build logs in case provenance is a concern.


## Usage

```
# start the container
docker run \
  --detach \
  --name hello-world \
  --publish 8000:8080 \
  ghcr.io/infrastructure-as-code/hello-world

# curl the container
curl http://0.0.0.0:8000/

# curl the health check endpoint which returns an empty response
curl http://0.0.0.0:8000/health

# curl Prometheus metrics
curl http://0.0.0.0:8000/metrics
```
