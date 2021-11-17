package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

/*
	Roles that are attributed to symbol occurrences.

	Internal: this currently mirrors low 9 bits of clang::index::SymbolRole with
	higher bits zeroed. These high bits may be exposed in the future.
*/
type SymbolRole uint32

const (
	SymbolRole_None        SymbolRole = C.CXSymbolRole_None
	SymbolRole_Declaration            = C.CXSymbolRole_Declaration
	SymbolRole_Definition             = C.CXSymbolRole_Definition
	SymbolRole_Reference              = C.CXSymbolRole_Reference
	SymbolRole_Read                   = C.CXSymbolRole_Read
	SymbolRole_Write                  = C.CXSymbolRole_Write
	SymbolRole_Call                   = C.CXSymbolRole_Call
	SymbolRole_Dynamic                = C.CXSymbolRole_Dynamic
	SymbolRole_AddressOf              = C.CXSymbolRole_AddressOf
	SymbolRole_Implicit               = C.CXSymbolRole_Implicit
)

func (sr SymbolRole) Spelling() string {
	switch sr {
	case SymbolRole_None:
		return "SymbolRole=None"
	case SymbolRole_Declaration:
		return "SymbolRole=Declaration"
	case SymbolRole_Definition:
		return "SymbolRole=Definition"
	case SymbolRole_Reference:
		return "SymbolRole=Reference"
	case SymbolRole_Read:
		return "SymbolRole=Read"
	case SymbolRole_Write:
		return "SymbolRole=Write"
	case SymbolRole_Call:
		return "SymbolRole=Call"
	case SymbolRole_Dynamic:
		return "SymbolRole=Dynamic"
	case SymbolRole_AddressOf:
		return "SymbolRole=AddressOf"
	case SymbolRole_Implicit:
		return "SymbolRole=Implicit"
	}

	return fmt.Sprintf("SymbolRole unknown %d", int(sr))
}

func (sr SymbolRole) String() string {
	return sr.Spelling()
}
