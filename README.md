# prototpl

## Overview

`prototpl` is a plugin for protobuf compilator (`protoc`) which allows to generate files using custom templates and rules.

Now only default template for go is supported. Also only go structs are generated.

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

Install `protoc-gen-tpl` plugin using (go)[https://golang.org]
```
go get github.com/apoprotsky/proto-gen-tpl
```

## How to use

Example how to generate go code from your proto files
```sh
protoc \
    --tpl_out=examples/out \
    --tpl_opt=module=github.com/apoprotsky/prototpl/examples/ \
    examples/proto/*.proto \
```
Parameters `module` has [same behaviour](https://developers.google.com/protocol-buffers/docs/reference/go-generated#invocation) as for go plugin

## Roadmap

- [ ] Go support
  - [x] Generate `go` files from `proto` files
  - [x] Generate structs types from messages
  - [ ] Generate structs fields tags
    - [x] Generate json tags
    - [ ] Generate custom tags
  - [ ] Generate types from enumerations
  - [ ] Generate constants from enumerations
  - Supported struct fields types:
    - [x] Scalar types (string, numbers)
    - [x] Arrays
    - [ ] Pointer to generated from message type
    - [ ] Generated from enumeration type
- [ ] Typescript support
  - [ ] Generate `ts` files from `proto` files
  - [ ] Generate interfaces from messages
  - Supported interfaces fields types:
    - [ ] Scalar types (string, numbers)
    - [ ] Arrays
    - [ ] Generated from message type
    - [ ] Generated from enumeration type
- [ ] PHP support
  - [ ] Generate `php` files from `proto` files
  - [ ] Generate classes from messages
