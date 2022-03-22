package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type NameRefFlags uint32

const (
	// NameRange_WantQualifier include the nested-name-specifier, e.g. Foo:: in x.Foo::y, in the range.
	NameRange_WantQualifier NameRefFlags = C.CXNameRange_WantQualifier
	// NameRange_WantTemplateArgs include the explicit template arguments, e.g. \<int> in x.f<int>, in the range.
	NameRange_WantTemplateArgs = C.CXNameRange_WantTemplateArgs
	// NameRange_WantSinglePiece if the name is non-contiguous, return the full spanning range.
	//
	// Non-contiguous names occur in Objective-C when a selector with two or more
	// parameters is used, or in C++ when using an operator:
	//  [object doSomething:here withValue:there]; // Objective-C
	//  return some_vector[1]; // C++
	NameRange_WantSinglePiece = C.CXNameRange_WantSinglePiece
)

func (nrf NameRefFlags) Spelling() string {
	switch nrf {
	case NameRange_WantQualifier:
		return "NameRange=WantQualifier"
	case NameRange_WantTemplateArgs:
		return "NameRange=WantTemplateArgs"
	case NameRange_WantSinglePiece:
		return "NameRange=WantSinglePiece"
	}

	return fmt.Sprintf("NameRefFlags unknown %d", int(nrf))
}

func (nrf NameRefFlags) String() string {
	return nrf.Spelling()
}
