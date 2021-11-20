package clang

// #include "./clang-c/Rewrite.h"
// #include "go-clang.h"
import "C"
import "unsafe"

type Rewriter struct {
	c C.CXRewriter
}

// Insert the specified string at the specified location in the original buffer.
func (r Rewriter) InsertTextBefore(loc SourceLocation, insert string) {
	c_insert := C.CString(insert)
	defer C.free(unsafe.Pointer(c_insert))

	C.clang_CXRewriter_insertTextBefore(r.c, loc.c, c_insert)
}

// Replace the specified range of characters in the input with the specified replacement.
func (r Rewriter) ReplaceText(toBeReplaced SourceRange, replacement string) {
	c_replacement := C.CString(replacement)
	defer C.free(unsafe.Pointer(c_replacement))

	C.clang_CXRewriter_replaceText(r.c, toBeReplaced.c, c_replacement)
}

// Remove the specified range.
func (r Rewriter) RemoveText(toBeRemoved SourceRange) {
	C.clang_CXRewriter_removeText(r.c, toBeRemoved.c)
}

// Save all changed files to disk. Returns 1 if any files were not saved successfully, returns 0 otherwise.
func (r Rewriter) OverwriteChangedFiles() int32 {
	return int32(C.clang_CXRewriter_overwriteChangedFiles(r.c))
}

// Write out rewritten version of the main file to stdout.
func (r Rewriter) WriteMainFileToStdOut() {
	C.clang_CXRewriter_writeMainFileToStdOut(r.c)
}

// Free the given CXRewriter.
func (r Rewriter) CXRewriter_Dispose() {
	C.clang_CXRewriter_dispose(r.c)
}
