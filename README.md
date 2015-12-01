# go-clang bootstrap [![GoDoc](https://godoc.org/github.com/go-clang/bootstrap?status.png)](https://godoc.org/github.com/go-clang/bootstrap) [![Build Status](https://travis-ci.org/go-clang/bootstrap.svg?branch=master)](https://travis-ci.org/go-clang/bootstrap)

Native Go bindings for Clang's C API.

## Install/Update

```bash
CGO_LDFLAGS="-L`llvm-config --libdir`" \
  go get -u github.com/go-clang/bootstrap/...
```

## Example/Usage

An example on how to use the AST visitor of Clang can be found in [/cmd/go-clang-dump/main.go](/cmd/go-clang-dump/main.go)

## I need bindings for a different Clang version

These can be found at: [bindings](https://github.com/go-clang/gen#where-are-the-bindings).

## I found a bug/missing a feature in go-clang

Use the issue tracker of the https://github.com/go-clang/gen repository.

## How is the binding generated?

The go-clang binding generation is in its own repository and can be found at [gen](https://github.com/go-clang/gen).
