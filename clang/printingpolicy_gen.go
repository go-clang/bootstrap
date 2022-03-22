package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// PrintingPolicy opaque pointer representing a policy that controls pretty printing for clang_getCursorPrettyPrinted.
type PrintingPolicy struct {
	c C.CXPrintingPolicy
}

// Property get a property value for the given printing policy.
func (pp PrintingPolicy) Property(property PrintingPolicyProperty) uint32 {
	return uint32(C.clang_PrintingPolicy_getProperty(pp.c, C.enum_CXPrintingPolicyProperty(property)))
}

// SetProperty set a property value for the given printing policy.
func (pp PrintingPolicy) SetProperty(property PrintingPolicyProperty, value uint32) {
	C.clang_PrintingPolicy_setProperty(pp.c, C.enum_CXPrintingPolicyProperty(property), C.uint(value))
}

// Dispose release a printing policy.
func (pp PrintingPolicy) Dispose() {
	C.clang_PrintingPolicy_dispose(pp.c)
}
