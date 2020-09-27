package clang

// #include "./clang-c/BuildSystem.h"
// #include "go-clang.h"
import "C"

// Object encapsulating information about a module.map file.
type ModuleMapDescriptor struct {
	c C.CXModuleMapDescriptor
}
