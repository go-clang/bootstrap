package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// IdxClientFile the client's data object that is associated with a CXFile.
type IdxClientFile struct {
	c C.CXIdxClientFile
}
