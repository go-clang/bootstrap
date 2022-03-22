package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type IdxContainerInfo struct {
	c *C.CXIdxContainerInfo
}

// Index_getClientContainer for retrieving a custom CXIdxClientContainer attached to a container.
func (ici *IdxContainerInfo) ClientContainer() IdxClientContainer {
	return IdxClientContainer{C.clang_index_getClientContainer(ici.c)}
}

// Index_setClientContainer for setting a custom CXIdxClientContainer attached to a container.
func (ici *IdxContainerInfo) SetClientContainer(icc IdxClientContainer) {
	C.clang_index_setClientContainer(ici.c, icc.c)
}

func (ici IdxContainerInfo) Cursor() Cursor {
	return Cursor{ici.c.cursor}
}
