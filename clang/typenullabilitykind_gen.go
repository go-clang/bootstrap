package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type TypeNullabilityKind uint32

const (
	// TypeNullability_NonNull values of this type can never be null.
	TypeNullability_NonNull TypeNullabilityKind = C.CXTypeNullability_NonNull
	// TypeNullability_Nullable values of this type can be null.
	TypeNullability_Nullable = C.CXTypeNullability_Nullable
	// TypeNullability_Unspecified whether values of this type can be null is (explicitly) unspecified. This captures a (fairly rare) case where we can't conclude anything about the nullability of the type even though it has been considered.
	TypeNullability_Unspecified = C.CXTypeNullability_Unspecified
	// TypeNullability_Invalid nullability is not applicable to this type.
	TypeNullability_Invalid = C.CXTypeNullability_Invalid
	// TypeNullability_NullableResult generally behaves like Nullable, except when used in a block parameter that was imported into a swift async method. There, swift will assume that the parameter can get null even if no error occurred. _Nullable parameters are assumed to only get null on error.
	TypeNullability_NullableResult = C.CXTypeNullability_NullableResult
)

func (tnk TypeNullabilityKind) Spelling() string {
	switch tnk {
	case TypeNullability_NonNull:
		return "TypeNullability=NonNull"
	case TypeNullability_Nullable:
		return "TypeNullability=Nullable"
	case TypeNullability_Unspecified:
		return "TypeNullability=Unspecified"
	case TypeNullability_Invalid:
		return "TypeNullability=Invalid"
	case TypeNullability_NullableResult:
		return "TypeNullability=NullableResult"
	}

	return fmt.Sprintf("TypeNullabilityKind unknown %d", int(tnk))
}

func (tnk TypeNullabilityKind) String() string {
	return tnk.Spelling()
}
