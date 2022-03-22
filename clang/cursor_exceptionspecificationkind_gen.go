package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Cursor_ExceptionSpecificationKind describes the exception specification of a cursor.
//
// A negative value indicates that the cursor is not a function declaration.
type Cursor_ExceptionSpecificationKind uint32

const (
	// Cursor_ExceptionSpecificationKind_None the cursor has no exception specification.
	Cursor_ExceptionSpecificationKind_None Cursor_ExceptionSpecificationKind = C.CXCursor_ExceptionSpecificationKind_None
	// Cursor_ExceptionSpecificationKind_DynamicNone the cursor has exception specification throw()
	Cursor_ExceptionSpecificationKind_DynamicNone = C.CXCursor_ExceptionSpecificationKind_DynamicNone
	// Cursor_ExceptionSpecificationKind_Dynamic the cursor has exception specification throw(T1, T2)
	Cursor_ExceptionSpecificationKind_Dynamic = C.CXCursor_ExceptionSpecificationKind_Dynamic
	// Cursor_ExceptionSpecificationKind_MSAny the cursor has exception specification throw(...).
	Cursor_ExceptionSpecificationKind_MSAny = C.CXCursor_ExceptionSpecificationKind_MSAny
	// Cursor_ExceptionSpecificationKind_BasicNoexcept the cursor has exception specification basic noexcept.
	Cursor_ExceptionSpecificationKind_BasicNoexcept = C.CXCursor_ExceptionSpecificationKind_BasicNoexcept
	// Cursor_ExceptionSpecificationKind_ComputedNoexcept the cursor has exception specification computed noexcept.
	Cursor_ExceptionSpecificationKind_ComputedNoexcept = C.CXCursor_ExceptionSpecificationKind_ComputedNoexcept
	// Cursor_ExceptionSpecificationKind_Unevaluated the exception specification has not yet been evaluated.
	Cursor_ExceptionSpecificationKind_Unevaluated = C.CXCursor_ExceptionSpecificationKind_Unevaluated
	// Cursor_ExceptionSpecificationKind_Uninstantiated the exception specification has not yet been instantiated.
	Cursor_ExceptionSpecificationKind_Uninstantiated = C.CXCursor_ExceptionSpecificationKind_Uninstantiated
	// Cursor_ExceptionSpecificationKind_Unparsed the exception specification has not been parsed yet.
	Cursor_ExceptionSpecificationKind_Unparsed = C.CXCursor_ExceptionSpecificationKind_Unparsed
	// Cursor_ExceptionSpecificationKind_NoThrow the cursor has a __declspec(nothrow) exception specification.
	Cursor_ExceptionSpecificationKind_NoThrow = C.CXCursor_ExceptionSpecificationKind_NoThrow
)

func (cesk Cursor_ExceptionSpecificationKind) Spelling() string {
	switch cesk {
	case Cursor_ExceptionSpecificationKind_None:
		return "Cursor=ExceptionSpecificationKind_None"
	case Cursor_ExceptionSpecificationKind_DynamicNone:
		return "Cursor=ExceptionSpecificationKind_DynamicNone"
	case Cursor_ExceptionSpecificationKind_Dynamic:
		return "Cursor=ExceptionSpecificationKind_Dynamic"
	case Cursor_ExceptionSpecificationKind_MSAny:
		return "Cursor=ExceptionSpecificationKind_MSAny"
	case Cursor_ExceptionSpecificationKind_BasicNoexcept:
		return "Cursor=ExceptionSpecificationKind_BasicNoexcept"
	case Cursor_ExceptionSpecificationKind_ComputedNoexcept:
		return "Cursor=ExceptionSpecificationKind_ComputedNoexcept"
	case Cursor_ExceptionSpecificationKind_Unevaluated:
		return "Cursor=ExceptionSpecificationKind_Unevaluated"
	case Cursor_ExceptionSpecificationKind_Uninstantiated:
		return "Cursor=ExceptionSpecificationKind_Uninstantiated"
	case Cursor_ExceptionSpecificationKind_Unparsed:
		return "Cursor=ExceptionSpecificationKind_Unparsed"
	case Cursor_ExceptionSpecificationKind_NoThrow:
		return "Cursor=ExceptionSpecificationKind_NoThrow"
	}

	return fmt.Sprintf("Cursor_ExceptionSpecificationKind unknown %d", int(cesk))
}

func (cesk Cursor_ExceptionSpecificationKind) String() string {
	return cesk.Spelling()
}
