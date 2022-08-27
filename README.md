# go-examples
[![CircleCI](https://circleci.com/gh/wdstar/go-examples.svg?style=shield)](https://circleci.com/gh/wdstar/go-examples)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=wdstar_go-examples&metric=alert_status)](https://sonarcloud.io/dashboard?id=wdstar_go-examples)

This source's dependency is managed by Go module.

## requirements

- Go >= 1.18

## Test

```
$ export GO111MODULE=on
$ go mod download
$ go test -cover ./...
```
