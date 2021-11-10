package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Evaluation result of a cursor
type EvalResult struct {
	c C.CXEvalResult
}

// Returns the kind of the evaluated result.
func (er EvalResult) Kind() EvalResultKind {
	return EvalResultKind(C.clang_EvalResult_getKind(er.c))
}

// Returns the evaluation result as integer if the kind is Int.
func (er EvalResult) AsInt() int32 {
	return int32(C.clang_EvalResult_getAsInt(er.c))
}

// Returns the evaluation result as a long long integer if the kind is Int. This prevents overflows that may happen if the result is returned with clang_EvalResult_getAsInt.
func (er EvalResult) AsLongLong() int64 {
	return int64(C.clang_EvalResult_getAsLongLong(er.c))
}

// Returns a non-zero value if the kind is Int and the evaluation result resulted in an unsigned integer.
func (er EvalResult) IsUnsignedInt() bool {
	o := C.clang_EvalResult_isUnsignedInt(er.c)

	return o != C.uint(0)
}

// Returns the evaluation result as an unsigned integer if the kind is Int and clang_EvalResult_isUnsignedInt is non-zero.
func (er EvalResult) AsUnsigned() uint64 {
	return uint64(C.clang_EvalResult_getAsUnsigned(er.c))
}

// Returns the evaluation result as double if the kind is double.
func (er EvalResult) AsDouble() float64 {
	return float64(C.clang_EvalResult_getAsDouble(er.c))
}

// Returns the evaluation result as a constant string if the kind is other than Int or float. User must not free this pointer, instead call clang_EvalResult_dispose on the CXEvalResult returned by clang_Cursor_Evaluate.
func (er EvalResult) AsStr() string {
	return C.GoString(C.clang_EvalResult_getAsStr(er.c))
}

// Disposes the created Eval memory.
func (er EvalResult) Dispose() {
	C.clang_EvalResult_dispose(er.c)
}
