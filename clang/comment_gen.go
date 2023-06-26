package clang

// #include "./clang-c/Documentation.h"
// #include "go-clang.h"
import "C"

// Comment a parsed comment.
type Comment struct {
	c C.CXComment
}

// Kind parameter Comment AST node of any kind.
//
// Returns the type of the AST node.
func (c Comment) Kind() CommentKind {
	return CommentKind(C.clang_Comment_getKind(c.c))
}

// NumChildren parameter Comment AST node of any kind.
//
// Returns number of children of the AST node.
func (c Comment) NumChildren() uint32 {
	return uint32(C.clang_Comment_getNumChildren(c.c))
}

// Child parameter Comment AST node of any kind.
//
// Parameter ChildIdx child index (zero-based).
//
// Returns the specified child of the AST node.
func (c Comment) Child(childIdx uint32) Comment {
	return Comment{C.clang_Comment_getChild(c.c, C.uint(childIdx))}
}

// IsWhitespace a CXComment_Paragraph node is considered whitespace if it contains
// only CXComment_Text nodes that are empty or whitespace.
//
// Other AST nodes (except CXComment_Paragraph and CXComment_Text) are
// never considered whitespace.
//
// Returns non-zero if Comment is whitespace.
func (c Comment) IsWhitespace() bool {
	o := C.clang_Comment_isWhitespace(c.c)

	return o != C.uint(0)
}

// InlineContentComment_HasTrailingNewline returns non-zero if Comment is inline content and has a newline immediately following it in the comment text. Newlines between paragraphs do not count.
func (c Comment) InlineContentComment_HasTrailingNewline() bool {
	o := C.clang_InlineContentComment_hasTrailingNewline(c.c)

	return o != C.uint(0)
}

// TextComment_getText parameter Comment a CXComment_Text AST node.
//
// Returns text contained in the AST node.
func (c Comment) TextComment_getText() string {
	o := cxstring{C.clang_TextComment_getText(c.c)}
	defer o.Dispose()

	return o.String()
}

// InlineCommandComment_getCommandName parameter Comment a CXComment_InlineCommand AST node.
//
// Returns name of the inline command.
func (c Comment) InlineCommandComment_getCommandName() string {
	o := cxstring{C.clang_InlineCommandComment_getCommandName(c.c)}
	defer o.Dispose()

	return o.String()
}

// InlineCommandComment_getRenderKind parameter Comment a CXComment_InlineCommand AST node.
//
// Returns the most appropriate rendering mode, chosen on command
// semantics in Doxygen.
func (c Comment) InlineCommandComment_getRenderKind() CommentInlineCommandRenderKind {
	return CommentInlineCommandRenderKind(C.clang_InlineCommandComment_getRenderKind(c.c))
}

// InlineCommandComment_getNumArgs parameter Comment a CXComment_InlineCommand AST node.
//
// Returns number of command arguments.
func (c Comment) InlineCommandComment_getNumArgs() uint32 {
	return uint32(C.clang_InlineCommandComment_getNumArgs(c.c))
}

// InlineCommandComment_getArgText parameter Comment a CXComment_InlineCommand AST node.
//
// Parameter ArgIdx argument index (zero-based).
//
// Returns text of the specified argument.
func (c Comment) InlineCommandComment_getArgText(argIdx uint32) string {
	o := cxstring{C.clang_InlineCommandComment_getArgText(c.c, C.uint(argIdx))}
	defer o.Dispose()

	return o.String()
}

// HTMLTagComment_getTagName parameter Comment a CXComment_HTMLStartTag or CXComment_HTMLEndTag AST
// node.
//
// Returns HTML tag name.
func (c Comment) HTMLTagComment_getTagName() string {
	o := cxstring{C.clang_HTMLTagComment_getTagName(c.c)}
	defer o.Dispose()

	return o.String()
}

// HTMLStartTagComment_IsSelfClosing parameter Comment a CXComment_HTMLStartTag AST node.
//
// Returns non-zero if tag is self-closing (for example, &lt;br /&gt;).
func (c Comment) HTMLStartTagComment_IsSelfClosing() bool {
	o := C.clang_HTMLStartTagComment_isSelfClosing(c.c)

	return o != C.uint(0)
}

// HTMLStartTag_getNumAttrs parameter Comment a CXComment_HTMLStartTag AST node.
//
// Returns number of attributes (name-value pairs) attached to the start tag.
func (c Comment) HTMLStartTag_getNumAttrs() uint32 {
	return uint32(C.clang_HTMLStartTag_getNumAttrs(c.c))
}

// HTMLStartTag_getAttrName parameter Comment a CXComment_HTMLStartTag AST node.
//
// Parameter AttrIdx attribute index (zero-based).
//
// Returns name of the specified attribute.
func (c Comment) HTMLStartTag_getAttrName(attrIdx uint32) string {
	o := cxstring{C.clang_HTMLStartTag_getAttrName(c.c, C.uint(attrIdx))}
	defer o.Dispose()

	return o.String()
}

// HTMLStartTag_getAttrValue parameter Comment a CXComment_HTMLStartTag AST node.
//
// Parameter AttrIdx attribute index (zero-based).
//
// Returns value of the specified attribute.
func (c Comment) HTMLStartTag_getAttrValue(attrIdx uint32) string {
	o := cxstring{C.clang_HTMLStartTag_getAttrValue(c.c, C.uint(attrIdx))}
	defer o.Dispose()

	return o.String()
}

// BlockCommandComment_getCommandName parameter Comment a CXComment_BlockCommand AST node.
//
// Returns name of the block command.
func (c Comment) BlockCommandComment_getCommandName() string {
	o := cxstring{C.clang_BlockCommandComment_getCommandName(c.c)}
	defer o.Dispose()

	return o.String()
}

// BlockCommandComment_getNumArgs parameter Comment a CXComment_BlockCommand AST node.
//
// Returns number of word-like arguments.
func (c Comment) BlockCommandComment_getNumArgs() uint32 {
	return uint32(C.clang_BlockCommandComment_getNumArgs(c.c))
}

// BlockCommandComment_getArgText parameter Comment a CXComment_BlockCommand AST node.
//
// Parameter ArgIdx argument index (zero-based).
//
// Returns text of the specified word-like argument.
func (c Comment) BlockCommandComment_getArgText(argIdx uint32) string {
	o := cxstring{C.clang_BlockCommandComment_getArgText(c.c, C.uint(argIdx))}
	defer o.Dispose()

	return o.String()
}

// BlockCommandComment_getParagraph parameter Comment a CXComment_BlockCommand or
// CXComment_VerbatimBlockCommand AST node.
//
// Returns paragraph argument of the block command.
func (c Comment) BlockCommandComment_getParagraph() Comment {
	return Comment{C.clang_BlockCommandComment_getParagraph(c.c)}
}

// ParamCommandComment_getParamName parameter Comment a CXComment_ParamCommand AST node.
//
// Returns parameter name.
func (c Comment) ParamCommandComment_getParamName() string {
	o := cxstring{C.clang_ParamCommandComment_getParamName(c.c)}
	defer o.Dispose()

	return o.String()
}

// ParamCommandComment_IsParamIndexValid parameter Comment a CXComment_ParamCommand AST node.
//
// Returns non-zero if the parameter that this AST node represents was found
// in the function prototype and clang_ParamCommandComment_getParamIndex
// function will return a meaningful value.
func (c Comment) ParamCommandComment_IsParamIndexValid() bool {
	o := C.clang_ParamCommandComment_isParamIndexValid(c.c)

	return o != C.uint(0)
}

// ParamCommandComment_getParamIndex parameter Comment a CXComment_ParamCommand AST node.
//
// Returns zero-based parameter index in function prototype.
func (c Comment) ParamCommandComment_getParamIndex() uint32 {
	return uint32(C.clang_ParamCommandComment_getParamIndex(c.c))
}

// ParamCommandComment_IsDirectionExplicit parameter Comment a CXComment_ParamCommand AST node.
//
// Returns non-zero if parameter passing direction was specified explicitly in
// the comment.
func (c Comment) ParamCommandComment_IsDirectionExplicit() bool {
	o := C.clang_ParamCommandComment_isDirectionExplicit(c.c)

	return o != C.uint(0)
}

// ParamCommandComment_getDirection parameter Comment a CXComment_ParamCommand AST node.
//
// Returns parameter passing direction.
func (c Comment) ParamCommandComment_getDirection() CommentParamPassDirection {
	return CommentParamPassDirection(C.clang_ParamCommandComment_getDirection(c.c))
}

// TParamCommandComment_getParamName parameter Comment a CXComment_TParamCommand AST node.
//
// Returns template parameter name.
func (c Comment) TParamCommandComment_getParamName() string {
	o := cxstring{C.clang_TParamCommandComment_getParamName(c.c)}
	defer o.Dispose()

	return o.String()
}

// TParamCommandComment_IsParamPositionValid parameter Comment a CXComment_TParamCommand AST node.
//
// Returns non-zero if the parameter that this AST node represents was found
// in the template parameter list and
// clang_TParamCommandComment_getDepth and
// clang_TParamCommandComment_getIndex functions will return a meaningful
// value.
func (c Comment) TParamCommandComment_IsParamPositionValid() bool {
	o := C.clang_TParamCommandComment_isParamPositionValid(c.c)

	return o != C.uint(0)
}

// TParamCommandComment_getDepth parameter Comment a CXComment_TParamCommand AST node.
//
// Returns zero-based nesting depth of this parameter in the template parameter list.
//
// For example,
//
//	template<typename C, template<typename T> class TT>
//	void test(TT<int> aaa);
//
// for C and TT nesting depth is 0,
// for T nesting depth is 1.
func (c Comment) TParamCommandComment_getDepth() uint32 {
	return uint32(C.clang_TParamCommandComment_getDepth(c.c))
}

// TParamCommandComment_getIndex parameter Comment a CXComment_TParamCommand AST node.
//
// Returns zero-based parameter index in the template parameter list at a
// given nesting depth.
//
// For example,
//
//	template<typename C, template<typename T> class TT>
//	void test(TT<int> aaa);
//
// for C and TT nesting depth is 0, so we can ask for index at depth 0:
// at depth 0 C's index is 0, TT's index is 1.
//
// For T nesting depth is 1, so we can ask for index at depth 0 and 1:
// at depth 0 T's index is 1 (same as TT's),
// at depth 1 T's index is 0.
func (c Comment) TParamCommandComment_getIndex(depth uint32) uint32 {
	return uint32(C.clang_TParamCommandComment_getIndex(c.c, C.uint(depth)))
}

// VerbatimBlockLineComment_getText parameter Comment a CXComment_VerbatimBlockLine AST node.
//
// Returns text contained in the AST node.
func (c Comment) VerbatimBlockLineComment_getText() string {
	o := cxstring{C.clang_VerbatimBlockLineComment_getText(c.c)}
	defer o.Dispose()

	return o.String()
}

// VerbatimLineComment_getText parameter Comment a CXComment_VerbatimLine AST node.
//
// Returns text contained in the AST node.
func (c Comment) VerbatimLineComment_getText() string {
	o := cxstring{C.clang_VerbatimLineComment_getText(c.c)}
	defer o.Dispose()

	return o.String()
}

// HTMLTagComment_getAsString convert an HTML tag AST node to string.
//
// Parameter Comment a CXComment_HTMLStartTag or CXComment_HTMLEndTag AST
// node.
//
// Returns string containing an HTML tag.
func (c Comment) HTMLTagComment_getAsString() string {
	o := cxstring{C.clang_HTMLTagComment_getAsString(c.c)}
	defer o.Dispose()

	return o.String()
}

// FullComment_getAsHTML convert a given full parsed comment to an HTML fragment.
//
// Specific details of HTML layout are subject to change. Don't try to parse
// this HTML back into an AST, use other APIs instead.
//
// Currently the following CSS classes are used:
// The "para-brief" for paragraph and equivalent commands;
// The "para-returns" for returns paragraph and equivalent commands;
// The "word-returns" for the "Returns" word in returns paragraph.
//
// Function argument documentation is rendered as a \<dl\> list with arguments
// sorted in function prototype order. CSS classes used:
// The "param-name-index-NUMBER" for parameter name (\<dt\>);
// The "param-descr-index-NUMBER" for parameter description (\<dd\>);
// The "param-name-index-invalid" and "param-descr-index-invalid" are used if
// parameter index is invalid.
//
// Template parameter documentation is rendered as a \<dl\> list with
// parameters sorted in template parameter list order. CSS classes used:
// The "tparam-name-index-NUMBER" for parameter name (\<dt\>);
// The "tparam-descr-index-NUMBER" for parameter description (\<dd\>);
// The "tparam-name-index-other" and "tparam-descr-index-other" are used for
// names inside template template parameters;
// The "tparam-name-index-invalid" and "tparam-descr-index-invalid" are used if
// parameter position is invalid.
//
// Parameter Comment a CXComment_FullComment AST node.
//
// Returns string containing an HTML fragment.
func (c Comment) FullComment_getAsHTML() string {
	o := cxstring{C.clang_FullComment_getAsHTML(c.c)}
	defer o.Dispose()

	return o.String()
}

// FullComment_getAsXML convert a given full parsed comment to an XML document.
//
// A Relax NG schema for the XML can be found in comment-xml-schema.rng file
// inside clang source tree.
//
// Parameter Comment a CXComment_FullComment AST node.
//
// Returns string containing an XML document.
func (c Comment) FullComment_getAsXML() string {
	o := cxstring{C.clang_FullComment_getAsXML(c.c)}
	defer o.Dispose()

	return o.String()
}

func (c Comment) TranslationUnit() TranslationUnit {
	return TranslationUnit{c.c.TranslationUnit}
}
