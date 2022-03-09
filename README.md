# Sandbox

[![Go Reference](https://pkg.go.dev/badge/github.com/FedericoSchonborn/sandbox.svg)](https://pkg.go.dev/github.com/FedericoSchonborn/sandbox)
[![Go](https://github.com/FedericoSchonborn/sandbox/actions/workflows/go.yml/badge.svg)](https://github.com/FedericoSchonborn/sandbox/actions/workflows/go.yml)
[![Rust](https://github.com/FedericoSchonborn/sandbox/actions/workflows/rust.yml/badge.svg)](https://github.com/FedericoSchonborn/sandbox/actions/workflows/rust.yml)

## Notes

### Go

Almost all Go packages on this repository require Go 1.18 or later, as they use
type parameters. To install Go 1.18, with a previously installed Go toolchain
run:

```sh
$ go install golang.org/dl/go1.18rc1@latest
# Installs the "go1.18rc1" command.
$ go1.18rc1 download
# Fetches and builds the Go toolchain.
```
