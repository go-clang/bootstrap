package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// LoadDiag_Error describes the kind of error that occurred (if any) in a call to clang_loadDiagnostics.
type LoadDiag_Error int32

const (
	// LoadDiag_None indicates that no error occurred.
	LoadDiag_None LoadDiag_Error = C.CXLoadDiag_None
	// LoadDiag_Unknown indicates that an unknown error occurred while attempting to deserialize diagnostics.
	LoadDiag_Unknown = C.CXLoadDiag_Unknown
	// LoadDiag_CannotLoad indicates that the file containing the serialized diagnostics could not be opened.
	LoadDiag_CannotLoad = C.CXLoadDiag_CannotLoad
	// LoadDiag_InvalidFile indicates that the serialized diagnostics file is invalid or corrupt.
	LoadDiag_InvalidFile = C.CXLoadDiag_InvalidFile
)

func (lde LoadDiag_Error) Spelling() string {
	switch lde {
	case LoadDiag_None:
		return "LoadDiag=None"
	case LoadDiag_Unknown:
		return "LoadDiag=Unknown"
	case LoadDiag_CannotLoad:
		return "LoadDiag=CannotLoad"
	case LoadDiag_InvalidFile:
		return "LoadDiag=InvalidFile"
	}

	return fmt.Sprintf("LoadDiag_Error unknown %d", int(lde))
}

func (lde LoadDiag_Error) String() string {
	return lde.Spelling()
}

func (lde LoadDiag_Error) Error() string {
	return lde.Spelling()
}
