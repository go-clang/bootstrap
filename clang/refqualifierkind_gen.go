package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type RefQualifierKind uint32

const (
	// RefQualifier_None no ref-qualifier was provided.
	RefQualifier_None RefQualifierKind = C.CXRefQualifier_None
	// RefQualifier_LValue an lvalue ref-qualifier was provided (&).
	RefQualifier_LValue = C.CXRefQualifier_LValue
	// RefQualifier_RValue an rvalue ref-qualifier was provided (&&).
	RefQualifier_RValue = C.CXRefQualifier_RValue
)

func (rqk RefQualifierKind) Spelling() string {
	switch rqk {
	case RefQualifier_None:
		return "RefQualifier=None"
	case RefQualifier_LValue:
		return "RefQualifier=LValue"
	case RefQualifier_RValue:
		return "RefQualifier=RValue"
	}

	return fmt.Sprintf("RefQualifierKind unknown %d", int(rqk))
}

func (rqk RefQualifierKind) String() string {
	return rqk.Spelling()
}
