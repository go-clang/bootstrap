package clang

// #include "./clang-c/BuildSystem.h"
// #include "go-clang.h"
import "C"

// Object encapsulating information about overlaying virtual file/directories over the real file system.
type VirtualFileOverlay struct {
	c C.CXVirtualFileOverlay
}
