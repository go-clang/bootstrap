package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// SourceRange identifies a half-open character range in the source code.
//
// Use clang_getRangeStart() and clang_getRangeEnd() to retrieve the
// starting and end locations from a source range, respectively.
type SourceRange struct {
	c C.CXSourceRange
}

// GetNullRange retrieve a NULL (invalid) source range.
func NewNullRange() SourceRange {
	return SourceRange{C.clang_getNullRange()}
}

// EqualRanges determine whether two ranges are equivalent.
//
// Returns non-zero if the ranges are the same, zero if they differ.
func (sr SourceRange) Equal(sr2 SourceRange) bool {
	o := C.clang_equalRanges(sr.c, sr2.c)

	return o != C.uint(0)
}

// IsNull returns non-zero if range is null.
func (sr SourceRange) IsNull() bool {
	o := C.clang_Range_isNull(sr.c)

	return o != C.int(0)
}

// GetRangeStart retrieve a source location representing the first character within a source range.
func (sr SourceRange) Start() SourceLocation {
	return SourceLocation{C.clang_getRangeStart(sr.c)}
}

// GetRangeEnd retrieve a source location representing the last character within a source range.
func (sr SourceRange) End() SourceLocation {
	return SourceLocation{C.clang_getRangeEnd(sr.c)}
}
