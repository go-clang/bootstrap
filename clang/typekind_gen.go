package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Describes the kind of type
type TypeKind uint32

const (
	// Represents an invalid type (e.g., where no type is available).
	Type_Invalid TypeKind = C.CXType_Invalid
	// A type whose specific kind is not exposed via this interface.
	Type_Unexposed = C.CXType_Unexposed
	// A type whose specific kind is not exposed via this interface.
	Type_Void = C.CXType_Void
	// A type whose specific kind is not exposed via this interface.
	Type_Bool = C.CXType_Bool
	// A type whose specific kind is not exposed via this interface.
	Type_Char_U = C.CXType_Char_U
	// A type whose specific kind is not exposed via this interface.
	Type_UChar = C.CXType_UChar
	// A type whose specific kind is not exposed via this interface.
	Type_Char16 = C.CXType_Char16
	// A type whose specific kind is not exposed via this interface.
	Type_Char32 = C.CXType_Char32
	// A type whose specific kind is not exposed via this interface.
	Type_UShort = C.CXType_UShort
	// A type whose specific kind is not exposed via this interface.
	Type_UInt = C.CXType_UInt
	// A type whose specific kind is not exposed via this interface.
	Type_ULong = C.CXType_ULong
	// A type whose specific kind is not exposed via this interface.
	Type_ULongLong = C.CXType_ULongLong
	// A type whose specific kind is not exposed via this interface.
	Type_UInt128 = C.CXType_UInt128
	// A type whose specific kind is not exposed via this interface.
	Type_Char_S = C.CXType_Char_S
	// A type whose specific kind is not exposed via this interface.
	Type_SChar = C.CXType_SChar
	// A type whose specific kind is not exposed via this interface.
	Type_WChar = C.CXType_WChar
	// A type whose specific kind is not exposed via this interface.
	Type_Short = C.CXType_Short
	// A type whose specific kind is not exposed via this interface.
	Type_Int = C.CXType_Int
	// A type whose specific kind is not exposed via this interface.
	Type_Long = C.CXType_Long
	// A type whose specific kind is not exposed via this interface.
	Type_LongLong = C.CXType_LongLong
	// A type whose specific kind is not exposed via this interface.
	Type_Int128 = C.CXType_Int128
	// A type whose specific kind is not exposed via this interface.
	Type_Float = C.CXType_Float
	// A type whose specific kind is not exposed via this interface.
	Type_Double = C.CXType_Double
	// A type whose specific kind is not exposed via this interface.
	Type_LongDouble = C.CXType_LongDouble
	// A type whose specific kind is not exposed via this interface.
	Type_NullPtr = C.CXType_NullPtr
	// A type whose specific kind is not exposed via this interface.
	Type_Overload = C.CXType_Overload
	// A type whose specific kind is not exposed via this interface.
	Type_Dependent = C.CXType_Dependent
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCId = C.CXType_ObjCId
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCClass = C.CXType_ObjCClass
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCSel = C.CXType_ObjCSel
	// A type whose specific kind is not exposed via this interface.
	Type_Float128 = C.CXType_Float128
	// A type whose specific kind is not exposed via this interface.
	Type_Half = C.CXType_Half
	// A type whose specific kind is not exposed via this interface.
	Type_Float16 = C.CXType_Float16
	// A type whose specific kind is not exposed via this interface.
	Type_ShortAccum = C.CXType_ShortAccum
	// A type whose specific kind is not exposed via this interface.
	Type_Accum = C.CXType_Accum
	// A type whose specific kind is not exposed via this interface.
	Type_LongAccum = C.CXType_LongAccum
	// A type whose specific kind is not exposed via this interface.
	Type_UShortAccum = C.CXType_UShortAccum
	// A type whose specific kind is not exposed via this interface.
	Type_UAccum = C.CXType_UAccum
	// A type whose specific kind is not exposed via this interface.
	Type_ULongAccum = C.CXType_ULongAccum
	// A type whose specific kind is not exposed via this interface.
	Type_FirstBuiltin = C.CXType_FirstBuiltin
	// A type whose specific kind is not exposed via this interface.
	Type_LastBuiltin = C.CXType_LastBuiltin
	// A type whose specific kind is not exposed via this interface.
	Type_Complex = C.CXType_Complex
	// A type whose specific kind is not exposed via this interface.
	Type_Pointer = C.CXType_Pointer
	// A type whose specific kind is not exposed via this interface.
	Type_BlockPointer = C.CXType_BlockPointer
	// A type whose specific kind is not exposed via this interface.
	Type_LValueReference = C.CXType_LValueReference
	// A type whose specific kind is not exposed via this interface.
	Type_RValueReference = C.CXType_RValueReference
	// A type whose specific kind is not exposed via this interface.
	Type_Record = C.CXType_Record
	// A type whose specific kind is not exposed via this interface.
	Type_Enum = C.CXType_Enum
	// A type whose specific kind is not exposed via this interface.
	Type_Typedef = C.CXType_Typedef
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCInterface = C.CXType_ObjCInterface
	// A type whose specific kind is not exposed via this interface.
	Type_ObjCObjectPointer = C.CXType_ObjCObjectPointer
	// A type whose specific kind is not exposed via this interface.
	Type_FunctionNoProto = C.CXType_FunctionNoProto
	// A type whose specific kind is not exposed via this interface.
	Type_FunctionProto = C.CXType_FunctionProto
	// A type whose specific kind is not exposed via this interface.
	Type_ConstantArray = C.CXType_ConstantArray
	// A type whose specific kind is not exposed via this interface.
	Type_Vector = C.CXType_Vector
	// A type whose specific kind is not exposed via this interface.
	Type_IncompleteArray = C.CXType_IncompleteArray
	// A type whose specific kind is not exposed via this interface.
	Type_VariableArray = C.CXType_VariableArray
	// A type whose specific kind is not exposed via this interface.
	Type_DependentSizedArray = C.CXType_DependentSizedArray
	// A type whose specific kind is not exposed via this interface.
	Type_MemberPointer = C.CXType_MemberPointer
	// A type whose specific kind is not exposed via this interface.
	Type_Auto = C.CXType_Auto
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_Elaborated = C.CXType_Elaborated
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_Pipe = C.CXType_Pipe
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dRO = C.CXType_OCLImage1dRO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dArrayRO = C.CXType_OCLImage1dArrayRO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dBufferRO = C.CXType_OCLImage1dBufferRO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dRO = C.CXType_OCLImage2dRO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayRO = C.CXType_OCLImage2dArrayRO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dDepthRO = C.CXType_OCLImage2dDepthRO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayDepthRO = C.CXType_OCLImage2dArrayDepthRO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dMSAARO = C.CXType_OCLImage2dMSAARO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayMSAARO = C.CXType_OCLImage2dArrayMSAARO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dMSAADepthRO = C.CXType_OCLImage2dMSAADepthRO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayMSAADepthRO = C.CXType_OCLImage2dArrayMSAADepthRO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage3dRO = C.CXType_OCLImage3dRO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dWO = C.CXType_OCLImage1dWO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dArrayWO = C.CXType_OCLImage1dArrayWO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dBufferWO = C.CXType_OCLImage1dBufferWO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dWO = C.CXType_OCLImage2dWO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayWO = C.CXType_OCLImage2dArrayWO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dDepthWO = C.CXType_OCLImage2dDepthWO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayDepthWO = C.CXType_OCLImage2dArrayDepthWO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dMSAAWO = C.CXType_OCLImage2dMSAAWO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayMSAAWO = C.CXType_OCLImage2dArrayMSAAWO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dMSAADepthWO = C.CXType_OCLImage2dMSAADepthWO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayMSAADepthWO = C.CXType_OCLImage2dArrayMSAADepthWO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage3dWO = C.CXType_OCLImage3dWO
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dRW = C.CXType_OCLImage1dRW
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dArrayRW = C.CXType_OCLImage1dArrayRW
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage1dBufferRW = C.CXType_OCLImage1dBufferRW
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dRW = C.CXType_OCLImage2dRW
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayRW = C.CXType_OCLImage2dArrayRW
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dDepthRW = C.CXType_OCLImage2dDepthRW
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayDepthRW = C.CXType_OCLImage2dArrayDepthRW
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dMSAARW = C.CXType_OCLImage2dMSAARW
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayMSAARW = C.CXType_OCLImage2dArrayMSAARW
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dMSAADepthRW = C.CXType_OCLImage2dMSAADepthRW
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage2dArrayMSAADepthRW = C.CXType_OCLImage2dArrayMSAADepthRW
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLImage3dRW = C.CXType_OCLImage3dRW
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLSampler = C.CXType_OCLSampler
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLEvent = C.CXType_OCLEvent
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLQueue = C.CXType_OCLQueue
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLReserveID = C.CXType_OCLReserveID
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_ObjCObject = C.CXType_ObjCObject
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_ObjCTypeParam = C.CXType_ObjCTypeParam
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_Attributed = C.CXType_Attributed
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCMcePayload = C.CXType_OCLIntelSubgroupAVCMcePayload
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImePayload = C.CXType_OCLIntelSubgroupAVCImePayload
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCRefPayload = C.CXType_OCLIntelSubgroupAVCRefPayload
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCSicPayload = C.CXType_OCLIntelSubgroupAVCSicPayload
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCMceResult = C.CXType_OCLIntelSubgroupAVCMceResult
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeResult = C.CXType_OCLIntelSubgroupAVCImeResult
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCRefResult = C.CXType_OCLIntelSubgroupAVCRefResult
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCSicResult = C.CXType_OCLIntelSubgroupAVCSicResult
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeResultSingleRefStreamout = C.CXType_OCLIntelSubgroupAVCImeResultSingleRefStreamout
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeResultDualRefStreamout = C.CXType_OCLIntelSubgroupAVCImeResultDualRefStreamout
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeSingleRefStreamin = C.CXType_OCLIntelSubgroupAVCImeSingleRefStreamin
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_OCLIntelSubgroupAVCImeDualRefStreamin = C.CXType_OCLIntelSubgroupAVCImeDualRefStreamin
	/*
		Represents a type that was referred to using an elaborated type keyword.

		E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	*/
	Type_ExtVector = C.CXType_ExtVector
)

// Retrieve the spelling of a given CXTypeKind.
func (tk TypeKind) Spelling() string {
	o := cxstring{C.clang_getTypeKindSpelling(C.enum_CXTypeKind(tk))}
	defer o.Dispose()

	return o.String()
}

func (tk TypeKind) String() string {
	return tk.Spelling()
}
