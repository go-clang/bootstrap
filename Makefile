.PHONY: all test docker/test

export CC := clang
export CXX := clang++

LLVM_LIBDIR?=$(shell llvm-config --libdir)
LLVM_VERSION?=4

all: test

test:
	CGO_LDFLAGS="-L${LLVM_LIBDIR} -Wl,-rpath,${LLVM_LIBDIR}" go test -v -race ./...

docker/test:
	docker container run --rm -it --mount type=bind,src=$(CURDIR),dst=/go/src/github.com/go-clang/bootstrap -w /go/src/github.com/go-clang/bootstrap goclang/base:${LLVM_VERSION} make test
