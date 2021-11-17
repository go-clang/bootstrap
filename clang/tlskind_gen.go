package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// Describe the "thread-local storage (TLS) kind" of the declaration referred to by a cursor.
type TLSKind uint32

const (
	TLS_None    TLSKind = C.CXTLS_None
	TLS_Dynamic         = C.CXTLS_Dynamic
	TLS_Static          = C.CXTLS_Static
)

func (tlsk TLSKind) Spelling() string {
	switch tlsk {
	case TLS_None:
		return "TLS=None"
	case TLS_Dynamic:
		return "TLS=Dynamic"
	case TLS_Static:
		return "TLS=Static"
	}

	return fmt.Sprintf("TLSKind unknown %d", int(tlsk))
}

func (tlsk TLSKind) String() string {
	return tlsk.Spelling()
}
