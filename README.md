# TLS - Simple TLS Client/Server

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/tls)](https://pkg.go.dev/github.com/go-zoox/tls)
[![Build Status](https://github.com/go-zoox/tls/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/tls/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/tls)](https://goreportcard.com/report/github.com/go-zoox/tls)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/tls/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/tls?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/tls.svg)](https://github.com/go-zoox/tls/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/tls.svg?label=Release)](https://github.com/go-zoox/tls/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/tls
```

## Getting Started

```go
// client
import (
  "testing"
  "github.com/go-zoox/tls"
)

func main(t *testing.T) {
	// @TODO
  // 1. send negotiate to server
  // 2. exchange data with negotiate secret
}
```

```go
// server
func main(t *testing.T) {
	// 1. wait for negotiate from client
  // 2. exchange data with negotiate secret
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).
