![op](https://raw.githubusercontent.com/txn2/op/master/mast.jpg)

WIP: op proxy for TXN2


## Release Packaging

Build test release:
```bash
goreleaser --skip-publish --rm-dist --skip-validate
```

Build and release:
```bash
GITHUB_TOKEN=$GITHUB_TOKEN goreleaser --rm-
```