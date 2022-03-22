package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

// CompletionContext bits that represent the context under which completion is occurring.
//
// The enumerators in this enumeration may be bitwise-OR'd together if multiple
// contexts are occurring simultaneously.
type CompletionContext uint32

const (
	// CompletionContext_Unexposed the context for completions is unexposed, as only Clang results should be included. (This is equivalent to having no context bits set.)
	CompletionContext_Unexposed CompletionContext = C.CXCompletionContext_Unexposed
	// CompletionContext_AnyType completions for any possible type should be included in the results.
	CompletionContext_AnyType = C.CXCompletionContext_AnyType
	// CompletionContext_AnyValue completions for any possible value (variables, function calls, etc.) should be included in the results.
	CompletionContext_AnyValue = C.CXCompletionContext_AnyValue
	// CompletionContext_ObjCObjectValue completions for values that resolve to an Objective-C object should be included in the results.
	CompletionContext_ObjCObjectValue = C.CXCompletionContext_ObjCObjectValue
	// CompletionContext_ObjCSelectorValue completions for values that resolve to an Objective-C selector should be included in the results.
	CompletionContext_ObjCSelectorValue = C.CXCompletionContext_ObjCSelectorValue
	// CompletionContext_CXXClassTypeValue completions for values that resolve to a C++ class type should be included in the results.
	CompletionContext_CXXClassTypeValue = C.CXCompletionContext_CXXClassTypeValue
	// CompletionContext_DotMemberAccess completions for fields of the member being accessed using the dot operator should be included in the results.
	CompletionContext_DotMemberAccess = C.CXCompletionContext_DotMemberAccess
	// CompletionContext_ArrowMemberAccess completions for fields of the member being accessed using the arrow operator should be included in the results.
	CompletionContext_ArrowMemberAccess = C.CXCompletionContext_ArrowMemberAccess
	// CompletionContext_ObjCPropertyAccess completions for properties of the Objective-C object being accessed using the dot operator should be included in the results.
	CompletionContext_ObjCPropertyAccess = C.CXCompletionContext_ObjCPropertyAccess
	// CompletionContext_EnumTag completions for enum tags should be included in the results.
	CompletionContext_EnumTag = C.CXCompletionContext_EnumTag
	// CompletionContext_UnionTag completions for union tags should be included in the results.
	CompletionContext_UnionTag = C.CXCompletionContext_UnionTag
	// CompletionContext_StructTag completions for struct tags should be included in the results.
	CompletionContext_StructTag = C.CXCompletionContext_StructTag
	// CompletionContext_ClassTag completions for C++ class names should be included in the results.
	CompletionContext_ClassTag = C.CXCompletionContext_ClassTag
	// CompletionContext_Namespace completions for C++ namespaces and namespace aliases should be included in the results.
	CompletionContext_Namespace = C.CXCompletionContext_Namespace
	// CompletionContext_NestedNameSpecifier completions for C++ nested name specifiers should be included in the results.
	CompletionContext_NestedNameSpecifier = C.CXCompletionContext_NestedNameSpecifier
	// CompletionContext_ObjCInterface completions for Objective-C interfaces (classes) should be included in the results.
	CompletionContext_ObjCInterface = C.CXCompletionContext_ObjCInterface
	// CompletionContext_ObjCProtocol completions for Objective-C protocols should be included in the results.
	CompletionContext_ObjCProtocol = C.CXCompletionContext_ObjCProtocol
	// CompletionContext_ObjCCategory completions for Objective-C categories should be included in the results.
	CompletionContext_ObjCCategory = C.CXCompletionContext_ObjCCategory
	// CompletionContext_ObjCInstanceMessage completions for Objective-C instance messages should be included in the results.
	CompletionContext_ObjCInstanceMessage = C.CXCompletionContext_ObjCInstanceMessage
	// CompletionContext_ObjCClassMessage completions for Objective-C class messages should be included in the results.
	CompletionContext_ObjCClassMessage = C.CXCompletionContext_ObjCClassMessage
	// CompletionContext_ObjCSelectorName completions for Objective-C selector names should be included in the results.
	CompletionContext_ObjCSelectorName = C.CXCompletionContext_ObjCSelectorName
	// CompletionContext_MacroName completions for preprocessor macro names should be included in the results.
	CompletionContext_MacroName = C.CXCompletionContext_MacroName
	// CompletionContext_NaturalLanguage natural language completions should be included in the results.
	CompletionContext_NaturalLanguage = C.CXCompletionContext_NaturalLanguage
	// CompletionContext_IncludedFile #include file completions should be included in the results.
	CompletionContext_IncludedFile = C.CXCompletionContext_IncludedFile
	// CompletionContext_Unknown the current context is unknown, so set all contexts.
	CompletionContext_Unknown = C.CXCompletionContext_Unknown
)

func (cc CompletionContext) Spelling() string {
	switch cc {
	case CompletionContext_Unexposed:
		return "CompletionContext=Unexposed"
	case CompletionContext_AnyType:
		return "CompletionContext=AnyType"
	case CompletionContext_AnyValue:
		return "CompletionContext=AnyValue"
	case CompletionContext_ObjCObjectValue:
		return "CompletionContext=ObjCObjectValue"
	case CompletionContext_ObjCSelectorValue:
		return "CompletionContext=ObjCSelectorValue"
	case CompletionContext_CXXClassTypeValue:
		return "CompletionContext=CXXClassTypeValue"
	case CompletionContext_DotMemberAccess:
		return "CompletionContext=DotMemberAccess"
	case CompletionContext_ArrowMemberAccess:
		return "CompletionContext=ArrowMemberAccess"
	case CompletionContext_ObjCPropertyAccess:
		return "CompletionContext=ObjCPropertyAccess"
	case CompletionContext_EnumTag:
		return "CompletionContext=EnumTag"
	case CompletionContext_UnionTag:
		return "CompletionContext=UnionTag"
	case CompletionContext_StructTag:
		return "CompletionContext=StructTag"
	case CompletionContext_ClassTag:
		return "CompletionContext=ClassTag"
	case CompletionContext_Namespace:
		return "CompletionContext=Namespace"
	case CompletionContext_NestedNameSpecifier:
		return "CompletionContext=NestedNameSpecifier"
	case CompletionContext_ObjCInterface:
		return "CompletionContext=ObjCInterface"
	case CompletionContext_ObjCProtocol:
		return "CompletionContext=ObjCProtocol"
	case CompletionContext_ObjCCategory:
		return "CompletionContext=ObjCCategory"
	case CompletionContext_ObjCInstanceMessage:
		return "CompletionContext=ObjCInstanceMessage"
	case CompletionContext_ObjCClassMessage:
		return "CompletionContext=ObjCClassMessage"
	case CompletionContext_ObjCSelectorName:
		return "CompletionContext=ObjCSelectorName"
	case CompletionContext_MacroName:
		return "CompletionContext=MacroName"
	case CompletionContext_NaturalLanguage:
		return "CompletionContext=NaturalLanguage"
	case CompletionContext_IncludedFile:
		return "CompletionContext=IncludedFile"
	case CompletionContext_Unknown:
		return "CompletionContext=Unknown"
	}

	return fmt.Sprintf("CompletionContext unknown %d", int(cc))
}

func (cc CompletionContext) String() string {
	return cc.Spelling()
}
