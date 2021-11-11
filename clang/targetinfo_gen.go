package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// An opaque type representing target information for a given translation unit.
type TargetInfo struct {
	c C.CXTargetInfo
}

// Destroy the CXTargetInfo object.
func (ti TargetInfo) Dispose() {
	C.clang_TargetInfo_dispose(ti.c)
}

/*
	Get the normalized target triple as a string.

	Returns the empty string in case of any error.
*/
func (ti TargetInfo) Triple() string {
	o := cxstring{C.clang_TargetInfo_getTriple(ti.c)}
	defer o.Dispose()

	return o.String()
}

/*
	Get the pointer width of the target in bits.

	Returns -1 in case of error.
*/
func (ti TargetInfo) PointerWidth() int32 {
	return int32(C.clang_TargetInfo_getPointerWidth(ti.c))
}
