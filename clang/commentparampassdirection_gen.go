package clang

// #include "./clang-c/Documentation.h"
// #include "go-clang.h"
import "C"
import "fmt"

// CommentParamPassDirection describes parameter passing direction for parameter or arg command.
type CommentParamPassDirection uint32

const (
	// CommentParamPassDirection_In the parameter is an input parameter.
	CommentParamPassDirection_In CommentParamPassDirection = C.CXCommentParamPassDirection_In
	// CommentParamPassDirection_Out the parameter is an output parameter.
	CommentParamPassDirection_Out = C.CXCommentParamPassDirection_Out
	// CommentParamPassDirection_InOut the parameter is an input and output parameter.
	CommentParamPassDirection_InOut = C.CXCommentParamPassDirection_InOut
)

func (cppd CommentParamPassDirection) Spelling() string {
	switch cppd {
	case CommentParamPassDirection_In:
		return "CommentParamPassDirection=In"
	case CommentParamPassDirection_Out:
		return "CommentParamPassDirection=Out"
	case CommentParamPassDirection_InOut:
		return "CommentParamPassDirection=InOut"
	}

	return fmt.Sprintf("CommentParamPassDirection unknown %d", int(cppd))
}

func (cppd CommentParamPassDirection) String() string {
	return cppd.Spelling()
}
