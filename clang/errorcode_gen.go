package clang

// #include "./clang-c/CXErrorCode.h"
// #include "go-clang.h"
import "C"
import "fmt"

// ErrorCode error codes returned by libclang routines.
//
// Zero (CXError_Success) is the only error code indicating success. Other
// error codes, including not yet assigned non-zero values, indicate errors.
type ErrorCode uint32

const (
	// Error_Success no error.
	Error_Success ErrorCode = C.CXError_Success
	// Error_Failure a generic error code, no further details are available.
	//
	// Errors of this kind can get their own specific error codes in future
	// libclang versions.
	Error_Failure = C.CXError_Failure
	// Error_Crashed libclang crashed while performing the requested operation.
	Error_Crashed = C.CXError_Crashed
	// Error_InvalidArguments the function detected that the arguments violate the function contract.
	Error_InvalidArguments = C.CXError_InvalidArguments
	// Error_ASTReadError an AST deserialization error has occurred.
	Error_ASTReadError = C.CXError_ASTReadError
)

func (ec ErrorCode) Spelling() string {
	switch ec {
	case Error_Success:
		return "Error=Success"
	case Error_Failure:
		return "Error=Failure"
	case Error_Crashed:
		return "Error=Crashed"
	case Error_InvalidArguments:
		return "Error=InvalidArguments"
	case Error_ASTReadError:
		return "Error=ASTReadError"
	}

	return fmt.Sprintf("ErrorCode unknown %d", int(ec))
}

func (ec ErrorCode) String() string {
	return ec.Spelling()
}
