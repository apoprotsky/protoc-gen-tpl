# protoc-gen-tpl

[![Go Report Card](https://goreportcard.com/badge/github.com/apoprotsky/protoc-gen-tpl)](https://goreportcard.com/report/github.com/apoprotsky/protoc-gen-tpl)

## Overview

`protoc-gen-tpl` is a plugin for protobuf compiler (`protoc`) which allows to generate files using custom templates and rules.

See in [checklist](#checklist) implemented and planned features.

## Prerequisites

You need protobuf compiler. See instructions on [Protocol Buffers site](https://developers.google.com/protocol-buffers).

Install on `Linux` using `apt`
```sh
apt install protobuf-compiler
```

Install on `macOS` using [Homebrew](https://brew.sh)
```sh
brew install protobuf
```

## Installation

Install `protoc-gen-tpl` plugin using [go](https://golang.org)
```
go get github.com/apoprotsky/protoc-gen-tpl
```

## How to use

Example how to generate go code from proto files
```sh
mkdir examples/out
protoc \
  --tpl_out=examples/out \
  --tpl_opt=prefix=github.com/apoprotsky/protoc-gen-tpl/examples/ \
  --tpl_opt=lang=go \
  --tpl_opt=lang=ts \
  --tpl_opt=lang=php \
  examples/proto/*.proto
```
Option `prefix` has [same behaviour](https://developers.google.com/protocol-buffers/docs/reference/go-generated#invocation) as `module=$PREFIX` for go plugin.

Option `lang` designates which languages to use for output files. This option can be specified multiple times.

## Checklist

### Go

- [x] Generate `go` files from `proto` files
- [x] Generate structs types from messages
- [x] Generate struct fields tags
  - [x] Generate json tags
  - [x] Generate custom tags
- [ ] Generate types from enumerations
- [ ] Generate constants from enumerations
- [ ] Supported struct fields types:
  - [x] Scalar types (string, numbers)
  - [x] Arrays
  - [ ] Message type
  - [ ] Enumeration type

### Typescript

- [x] Generate `ts` files from `proto` files
- [x] Generate interfaces from messages
- [ ] Generate enumerations
- [ ] Supported interface fields types:
  - [x] Scalar types (string, boolean, number, bigint)
  - [ ] Arrays
  - [ ] Message type
  - [ ] Enumeration type

### PHP

- [x] Generate `php` files from `proto` files
- [x] Generate classes from messages
- [ ] Generate constants from enumerations
