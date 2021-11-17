package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type TypeNullabilityKind uint32

const (
	// Values of this type can never be null.
	TypeNullability_NonNull TypeNullabilityKind = C.CXTypeNullability_NonNull
	// Values of this type can be null.
	TypeNullability_Nullable = C.CXTypeNullability_Nullable
	// Whether values of this type can be null is (explicitly) unspecified. This captures a (fairly rare) case where we can't conclude anything about the nullability of the type even though it has been considered.
	TypeNullability_Unspecified = C.CXTypeNullability_Unspecified
	// Nullability is not applicable to this type.
	TypeNullability_Invalid = C.CXTypeNullability_Invalid
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
	}

	return fmt.Sprintf("TypeNullabilityKind unknown %d", int(tnk))
}

func (tnk TypeNullabilityKind) String() string {
	return tnk.Spelling()
}
