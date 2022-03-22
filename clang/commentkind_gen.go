package clang

// #include "./clang-c/Documentation.h"
// #include "go-clang.h"
import "C"
import "fmt"

// CommentKind describes the type of the comment AST node (CXComment). A comment node can be considered block content (e. g., paragraph), inline content (plain text) or neither (the root AST node).
type CommentKind uint32

const (
	// Comment_Null null comment. No AST node is constructed at the requested location because there is no text or a syntax error.
	Comment_Null CommentKind = C.CXComment_Null
	// Comment_Text plain text. Inline content.
	Comment_Text = C.CXComment_Text
	// Comment_InlineCommand a command with word-like arguments that is considered inline content.
	//
	// For example: \command.
	Comment_InlineCommand = C.CXComment_InlineCommand
	// Comment_HTMLStartTag hTML start tag with attributes (name-value pairs). Considered
	// inline content.
	//
	// For example:
	//  <br> <br /> <a href="http://example.org/">
	Comment_HTMLStartTag = C.CXComment_HTMLStartTag
	// Comment_HTMLEndTag hTML end tag. Considered inline content.
	//
	// For example:
	//  </a>
	Comment_HTMLEndTag = C.CXComment_HTMLEndTag
	// Comment_Paragraph a paragraph, contains inline comment. The paragraph itself is block content.
	Comment_Paragraph = C.CXComment_Paragraph
	// Comment_BlockCommand a command that has zero or more word-like arguments (number of
	// word-like arguments depends on command name) and a paragraph as an
	// argument. Block command is block content.
	//
	// Paragraph argument is also a child of the block command.
	//
	// For example: \has 0 word-like arguments and a paragraph argument.
	//
	// AST nodes of special kinds that parser knows about (e. g., \\param
	// command) have their own node kinds.
	Comment_BlockCommand = C.CXComment_BlockCommand
	// Comment_ParamCommand a parameter or arg command that describes the function parameter
	// (name, passing direction, description).
	//
	// For example: parameter [in] ParamName description.
	Comment_ParamCommand = C.CXComment_ParamCommand
	// Comment_TParamCommand a template param command that describes a template parameter (name and
	// description).
	//
	// For example: template param T description.
	Comment_TParamCommand = C.CXComment_TParamCommand
	// Comment_VerbatimBlockCommand a verbatim block command (e. g., preformatted code). Verbatim
	// block has an opening and a closing command and contains multiple lines of
	// text (CXComment_VerbatimBlockLine child nodes).
	//
	// For example:
	//  aaa
	Comment_VerbatimBlockCommand = C.CXComment_VerbatimBlockCommand
	// Comment_VerbatimBlockLine a line of text that is contained within a CXComment_VerbatimBlockCommand node.
	Comment_VerbatimBlockLine = C.CXComment_VerbatimBlockLine
	// Comment_VerbatimLine a verbatim line command. Verbatim line has an opening command, a single line of text (up to the newline after the opening command) and has no closing command.
	Comment_VerbatimLine = C.CXComment_VerbatimLine
	// Comment_FullComment a full comment attached to a declaration, contains block content.
	Comment_FullComment = C.CXComment_FullComment
)

func (ck CommentKind) Spelling() string {
	switch ck {
	case Comment_Null:
		return "Comment=Null"
	case Comment_Text:
		return "Comment=Text"
	case Comment_InlineCommand:
		return "Comment=InlineCommand"
	case Comment_HTMLStartTag:
		return "Comment=HTMLStartTag"
	case Comment_HTMLEndTag:
		return "Comment=HTMLEndTag"
	case Comment_Paragraph:
		return "Comment=Paragraph"
	case Comment_BlockCommand:
		return "Comment=BlockCommand"
	case Comment_ParamCommand:
		return "Comment=ParamCommand"
	case Comment_TParamCommand:
		return "Comment=TParamCommand"
	case Comment_VerbatimBlockCommand:
		return "Comment=VerbatimBlockCommand"
	case Comment_VerbatimBlockLine:
		return "Comment=VerbatimBlockLine"
	case Comment_VerbatimLine:
		return "Comment=VerbatimLine"
	case Comment_FullComment:
		return "Comment=FullComment"
	}

	return fmt.Sprintf("CommentKind unknown %d", int(ck))
}

func (ck CommentKind) String() string {
	return ck.Spelling()
}
