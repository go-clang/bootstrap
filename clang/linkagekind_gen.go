package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// LinkageKind describe the linkage of the entity referred to by a cursor.
type LinkageKind uint32

const (
	// Linkage_Invalid this value indicates that no linkage information is available for a provided CXCursor.
	Linkage_Invalid LinkageKind = C.CXLinkage_Invalid
	// Linkage_NoLinkage this is the linkage for variables, parameters, and so on that have automatic storage. This covers normal (non-extern) local variables.
	Linkage_NoLinkage = C.CXLinkage_NoLinkage
	// Linkage_Internal this is the linkage for static variables and static functions.
	Linkage_Internal = C.CXLinkage_Internal
	// Linkage_UniqueExternal this is the linkage for entities with external linkage that live in C++ anonymous namespaces.
	Linkage_UniqueExternal = C.CXLinkage_UniqueExternal
	// Linkage_External this is the linkage for entities with true, external linkage.
	Linkage_External = C.CXLinkage_External
)

func (lk LinkageKind) Spelling() string {
	switch lk {
	case Linkage_Invalid:
		return "Linkage=Invalid"
	case Linkage_NoLinkage:
		return "Linkage=NoLinkage"
	case Linkage_Internal:
		return "Linkage=Internal"
	case Linkage_UniqueExternal:
		return "Linkage=UniqueExternal"
	case Linkage_External:
		return "Linkage=External"
	}

	return fmt.Sprintf("LinkageKind unknown %d", int(lk))
}

func (lk LinkageKind) String() string {
	return lk.Spelling()
}
