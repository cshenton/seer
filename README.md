# Seer: Forecasting Microservice
[![Build Status](https://travis-ci.org/cshenton/seer.svg?branch=master)](https://travis-ci.org/cshenton/seer)
[![Coverage Status](https://coveralls.io/repos/github/cshenton/seer/badge.svg?branch=master)](https://coveralls.io/github/cshenton/seer?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/cshenton/seer)](https://goreportcard.com/report/github.com/cshenton/seer)
[![Docker Automated Build](https://img.shields.io/docker/build/cshenton/seer.svg)](https://hub.docker.com/r/cshenton/seer/)


## What is this?

Seer is a service that can do real-time forecasting on time series data. You can
run it using one of the provided docker images, then talk to it using a `gRPC`
client.

It allows you to simply stream in your data as you receive it, then generate
up to date forecasts.


## Why did I build it?

Seer is the open-source release of a product I built to provide time series
forecasting as a service. The original product was a multi-tenant cloud API, but
this release is a single-tenant version of that product intended for internal
corporate or personal use.


## How do I use it?

To run a copy of the server locally, use the latest docker image:

```
docker run -d -p 8080:8080 cshenton/seer
```

Then use one of the available clients to stream in data and start forecasting:

- [python](https://github.com/cshenton/seer-python)


## Roadmap

- Python client
- Speed and accuracy benchmarks


## Other

#### Generating server snippets
```
protoc -I seer/ seer/seer.proto --go_out=plugins=grpc:seer
```