package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// TokenKind describes a kind of token.
type TokenKind uint32

const (
	// Token_Punctuation a token that contains some kind of punctuation.
	Token_Punctuation TokenKind = C.CXToken_Punctuation
	// Token_Keyword a language keyword.
	Token_Keyword = C.CXToken_Keyword
	// Token_Identifier an identifier (that is not a keyword).
	Token_Identifier = C.CXToken_Identifier
	// Token_Literal a numeric, string, or character literal.
	Token_Literal = C.CXToken_Literal
	// Token_Comment a comment.
	Token_Comment = C.CXToken_Comment
)

func (tk TokenKind) Spelling() string {
	switch tk {
	case Token_Punctuation:
		return "Token=Punctuation"
	case Token_Keyword:
		return "Token=Keyword"
	case Token_Identifier:
		return "Token=Identifier"
	case Token_Literal:
		return "Token=Literal"
	case Token_Comment:
		return "Token=Comment"
	}

	return fmt.Sprintf("TokenKind unknown %d", int(tk))
}

func (tk TokenKind) String() string {
	return tk.Spelling()
}
