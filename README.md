![op](https://raw.githubusercontent.com/txn2/op/master/mast.jpg)
[![op Release](https://img.shields.io/github/release/txn2/op.svg)](https://github.com/txn2/query/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/txn2/op)](https://goreportcard.com/report/github.com/txn2/op)
[![GoDoc](https://godoc.org/github.com/txn2/op?status.svg)](https://godoc.org/github.com/txn2/op)
[![Docker Container Image Size](https://shields.beevelop.com/docker/image/image-size/txn2/op/latest.svg)](https://hub.docker.com/r/txn2/op/)
[![Docker Container Layers](https://shields.beevelop.com/docker/image/layers/txn2/op/latest.svg)](https://hub.docker.com/r/txn2/op/)

WIP: op proxy for TXN2


## Release Packaging

Build test release:
```bash
goreleaser --skip-publish --rm-dist --skip-validate
```

Build and release:
```bash
GITHUB_TOKEN=$GITHUB_TOKEN goreleaser --rm-dist
```