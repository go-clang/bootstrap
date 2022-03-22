package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// ClientData opaque pointer representing client data that will be passed through to various callbacks and visitors.
type ClientData struct {
	c C.CXClientData
}
