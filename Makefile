.PHONY: all test docker/test

export ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

export CC := clang
export CXX := clang++

ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(ARGS):;@:) # turn arguments into do-nothing targets
export ARGS

LLVM_LIBDIR?=$(shell llvm-config --libdir)

all: test

test:
	CGO_LDFLAGS="-L${LLVM_LIBDIR} -Wl,-rpath,${LLVM_LIBDIR}" go test -v -race -timeout 60s ./...

LLVM_VERSION?=4

docker/test:
	docker container run --rm -it --mount type=bind,src=$(CURDIR),dst=/go/src/github.com/go-clang/bootstrap -w /go/src/github.com/go-clang/bootstrap goclang/base:${LLVM_VERSION} make test
