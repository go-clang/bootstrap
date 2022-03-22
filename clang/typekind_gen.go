package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// TypeKind describes the kind of type
type TypeKind uint32

const (
	// Type_Invalid represents an invalid type (e.g., where no type is available).
	Type_Invalid TypeKind = C.CXType_Invalid
	// Type_Unexposed a type whose specific kind is not exposed via this interface.
	Type_Unexposed = C.CXType_Unexposed
	// Type_Void a type whose specific kind is not exposed via this interface.
	Type_Void = C.CXType_Void
	// Type_Bool a type whose specific kind is not exposed via this interface.
	Type_Bool = C.CXType_Bool
	// Type_Char_U a type whose specific kind is not exposed via this interface.
	Type_Char_U = C.CXType_Char_U
	// Type_UChar a type whose specific kind is not exposed via this interface.
	Type_UChar = C.CXType_UChar
	// Type_Char16 a type whose specific kind is not exposed via this interface.
	Type_Char16 = C.CXType_Char16
	// Type_Char32 a type whose specific kind is not exposed via this interface.
	Type_Char32 = C.CXType_Char32
	// Type_UShort a type whose specific kind is not exposed via this interface.
	Type_UShort = C.CXType_UShort
	// Type_UInt a type whose specific kind is not exposed via this interface.
	Type_UInt = C.CXType_UInt
	// Type_ULong a type whose specific kind is not exposed via this interface.
	Type_ULong = C.CXType_ULong
	// Type_ULongLong a type whose specific kind is not exposed via this interface.
	Type_ULongLong = C.CXType_ULongLong
	// Type_UInt128 a type whose specific kind is not exposed via this interface.
	Type_UInt128 = C.CXType_UInt128
	// Type_Char_S a type whose specific kind is not exposed via this interface.
	Type_Char_S = C.CXType_Char_S
	// Type_SChar a type whose specific kind is not exposed via this interface.
	Type_SChar = C.CXType_SChar
	// Type_WChar a type whose specific kind is not exposed via this interface.
	Type_WChar = C.CXType_WChar
	// Type_Short a type whose specific kind is not exposed via this interface.
	Type_Short = C.CXType_Short
	// Type_Int a type whose specific kind is not exposed via this interface.
	Type_Int = C.CXType_Int
	// Type_Long a type whose specific kind is not exposed via this interface.
	Type_Long = C.CXType_Long
	// Type_LongLong a type whose specific kind is not exposed via this interface.
	Type_LongLong = C.CXType_LongLong
	// Type_Int128 a type whose specific kind is not exposed via this interface.
	Type_Int128 = C.CXType_Int128
	// Type_Float a type whose specific kind is not exposed via this interface.
	Type_Float = C.CXType_Float
	// Type_Double a type whose specific kind is not exposed via this interface.
	Type_Double = C.CXType_Double
	// Type_LongDouble a type whose specific kind is not exposed via this interface.
	Type_LongDouble = C.CXType_LongDouble
	// Type_NullPtr a type whose specific kind is not exposed via this interface.
	Type_NullPtr = C.CXType_NullPtr
	// Type_Overload a type whose specific kind is not exposed via this interface.
	Type_Overload = C.CXType_Overload
	// Type_Dependent a type whose specific kind is not exposed via this interface.
	Type_Dependent = C.CXType_Dependent
	// Type_ObjCId a type whose specific kind is not exposed via this interface.
	Type_ObjCId = C.CXType_ObjCId
	// Type_ObjCClass a type whose specific kind is not exposed via this interface.
	Type_ObjCClass = C.CXType_ObjCClass
	// Type_ObjCSel a type whose specific kind is not exposed via this interface.
	Type_ObjCSel = C.CXType_ObjCSel
	// Type_Float128 a type whose specific kind is not exposed via this interface.
	Type_Float128 = C.CXType_Float128
	// Type_Half a type whose specific kind is not exposed via this interface.
	Type_Half = C.CXType_Half
	// Type_Float16 a type whose specific kind is not exposed via this interface.
	Type_Float16 = C.CXType_Float16
	// Type_ShortAccum a type whose specific kind is not exposed via this interface.
	Type_ShortAccum = C.CXType_ShortAccum
	// Type_Accum a type whose specific kind is not exposed via this interface.
	Type_Accum = C.CXType_Accum
	// Type_LongAccum a type whose specific kind is not exposed via this interface.
	Type_LongAccum = C.CXType_LongAccum
	// Type_UShortAccum a type whose specific kind is not exposed via this interface.
	Type_UShortAccum = C.CXType_UShortAccum
	// Type_UAccum a type whose specific kind is not exposed via this interface.
	Type_UAccum = C.CXType_UAccum
	// Type_ULongAccum a type whose specific kind is not exposed via this interface.
	Type_ULongAccum = C.CXType_ULongAccum
	// Type_BFloat16 a type whose specific kind is not exposed via this interface.
	Type_BFloat16 = C.CXType_BFloat16
	// Type_FirstBuiltin a type whose specific kind is not exposed via this interface.
	Type_FirstBuiltin = C.CXType_FirstBuiltin
	// Type_LastBuiltin a type whose specific kind is not exposed via this interface.
	Type_LastBuiltin = C.CXType_LastBuiltin
	// Type_Complex a type whose specific kind is not exposed via this interface.
	Type_Complex = C.CXType_Complex
	// Type_Pointer a type whose specific kind is not exposed via this interface.
	Type_Pointer = C.CXType_Pointer
	// Type_BlockPointer a type whose specific kind is not exposed via this interface.
	Type_BlockPointer = C.CXType_BlockPointer
	// Type_LValueReference a type whose specific kind is not exposed via this interface.
	Type_LValueReference = C.CXType_LValueReference
	// Type_RValueReference a type whose specific kind is not exposed via this interface.
	Type_RValueReference = C.CXType_RValueReference
	// Type_Record a type whose specific kind is not exposed via this interface.
	Type_Record = C.CXType_Record
	// Type_Enum a type whose specific kind is not exposed via this interface.
	Type_Enum = C.CXType_Enum
	// Type_Typedef a type whose specific kind is not exposed via this interface.
	Type_Typedef = C.CXType_Typedef
	// Type_ObjCInterface a type whose specific kind is not exposed via this interface.
	Type_ObjCInterface = C.CXType_ObjCInterface
	// Type_ObjCObjectPointer a type whose specific kind is not exposed via this interface.
	Type_ObjCObjectPointer = C.CXType_ObjCObjectPointer
	// Type_FunctionNoProto a type whose specific kind is not exposed via this interface.
	Type_FunctionNoProto = C.CXType_FunctionNoProto
	// Type_FunctionProto a type whose specific kind is not exposed via this interface.
	Type_FunctionProto = C.CXType_FunctionProto
	// Type_ConstantArray a type whose specific kind is not exposed via this interface.
	Type_ConstantArray = C.CXType_ConstantArray
	// Type_Vector a type whose specific kind is not exposed via this interface.
	Type_Vector = C.CXType_Vector
	// Type_IncompleteArray a type whose specific kind is not exposed via this interface.
	Type_IncompleteArray = C.CXType_IncompleteArray
	// Type_VariableArray a type whose specific kind is not exposed via this interface.
	Type_VariableArray = C.CXType_VariableArray
	// Type_DependentSizedArray a type whose specific kind is not exposed via this interface.
	Type_DependentSizedArray = C.CXType_DependentSizedArray
	// Type_MemberPointer a type whose specific kind is not exposed via this interface.
	Type_MemberPointer = C.CXType_MemberPointer
	// Type_Auto a type whose specific kind is not exposed via this interface.
	Type_Auto = C.CXType_Auto
	// Type_Elaborated represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_Elaborated = C.CXType_Elaborated
	// Type_Pipe represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_Pipe = C.CXType_Pipe
	// Type_OCLImage1dRO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage1dRO = C.CXType_OCLImage1dRO
	// Type_OCLImage1dArrayRO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage1dArrayRO = C.CXType_OCLImage1dArrayRO
	// Type_OCLImage1dBufferRO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage1dBufferRO = C.CXType_OCLImage1dBufferRO
	// Type_OCLImage2dRO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dRO = C.CXType_OCLImage2dRO
	// Type_OCLImage2dArrayRO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dArrayRO = C.CXType_OCLImage2dArrayRO
	// Type_OCLImage2dDepthRO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dDepthRO = C.CXType_OCLImage2dDepthRO
	// Type_OCLImage2dArrayDepthRO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dArrayDepthRO = C.CXType_OCLImage2dArrayDepthRO
	// Type_OCLImage2dMSAARO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dMSAARO = C.CXType_OCLImage2dMSAARO
	// Type_OCLImage2dArrayMSAARO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dArrayMSAARO = C.CXType_OCLImage2dArrayMSAARO
	// Type_OCLImage2dMSAADepthRO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dMSAADepthRO = C.CXType_OCLImage2dMSAADepthRO
	// Type_OCLImage2dArrayMSAADepthRO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dArrayMSAADepthRO = C.CXType_OCLImage2dArrayMSAADepthRO
	// Type_OCLImage3dRO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage3dRO = C.CXType_OCLImage3dRO
	// Type_OCLImage1dWO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage1dWO = C.CXType_OCLImage1dWO
	// Type_OCLImage1dArrayWO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage1dArrayWO = C.CXType_OCLImage1dArrayWO
	// Type_OCLImage1dBufferWO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage1dBufferWO = C.CXType_OCLImage1dBufferWO
	// Type_OCLImage2dWO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dWO = C.CXType_OCLImage2dWO
	// Type_OCLImage2dArrayWO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dArrayWO = C.CXType_OCLImage2dArrayWO
	// Type_OCLImage2dDepthWO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dDepthWO = C.CXType_OCLImage2dDepthWO
	// Type_OCLImage2dArrayDepthWO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dArrayDepthWO = C.CXType_OCLImage2dArrayDepthWO
	// Type_OCLImage2dMSAAWO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dMSAAWO = C.CXType_OCLImage2dMSAAWO
	// Type_OCLImage2dArrayMSAAWO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dArrayMSAAWO = C.CXType_OCLImage2dArrayMSAAWO
	// Type_OCLImage2dMSAADepthWO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dMSAADepthWO = C.CXType_OCLImage2dMSAADepthWO
	// Type_OCLImage2dArrayMSAADepthWO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dArrayMSAADepthWO = C.CXType_OCLImage2dArrayMSAADepthWO
	// Type_OCLImage3dWO represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage3dWO = C.CXType_OCLImage3dWO
	// Type_OCLImage1dRW represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage1dRW = C.CXType_OCLImage1dRW
	// Type_OCLImage1dArrayRW represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage1dArrayRW = C.CXType_OCLImage1dArrayRW
	// Type_OCLImage1dBufferRW represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage1dBufferRW = C.CXType_OCLImage1dBufferRW
	// Type_OCLImage2dRW represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dRW = C.CXType_OCLImage2dRW
	// Type_OCLImage2dArrayRW represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dArrayRW = C.CXType_OCLImage2dArrayRW
	// Type_OCLImage2dDepthRW represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dDepthRW = C.CXType_OCLImage2dDepthRW
	// Type_OCLImage2dArrayDepthRW represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dArrayDepthRW = C.CXType_OCLImage2dArrayDepthRW
	// Type_OCLImage2dMSAARW represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dMSAARW = C.CXType_OCLImage2dMSAARW
	// Type_OCLImage2dArrayMSAARW represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dArrayMSAARW = C.CXType_OCLImage2dArrayMSAARW
	// Type_OCLImage2dMSAADepthRW represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dMSAADepthRW = C.CXType_OCLImage2dMSAADepthRW
	// Type_OCLImage2dArrayMSAADepthRW represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage2dArrayMSAADepthRW = C.CXType_OCLImage2dArrayMSAADepthRW
	// Type_OCLImage3dRW represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLImage3dRW = C.CXType_OCLImage3dRW
	// Type_OCLSampler represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLSampler = C.CXType_OCLSampler
	// Type_OCLEvent represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLEvent = C.CXType_OCLEvent
	// Type_OCLQueue represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLQueue = C.CXType_OCLQueue
	// Type_OCLReserveID represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLReserveID = C.CXType_OCLReserveID
	// Type_ObjCObject represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_ObjCObject = C.CXType_ObjCObject
	// Type_ObjCTypeParam represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_ObjCTypeParam = C.CXType_ObjCTypeParam
	// Type_Attributed represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_Attributed = C.CXType_Attributed
	// Type_OCLIntelSubgroupAVCMcePayload represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLIntelSubgroupAVCMcePayload = C.CXType_OCLIntelSubgroupAVCMcePayload
	// Type_OCLIntelSubgroupAVCImePayload represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLIntelSubgroupAVCImePayload = C.CXType_OCLIntelSubgroupAVCImePayload
	// Type_OCLIntelSubgroupAVCRefPayload represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLIntelSubgroupAVCRefPayload = C.CXType_OCLIntelSubgroupAVCRefPayload
	// Type_OCLIntelSubgroupAVCSicPayload represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLIntelSubgroupAVCSicPayload = C.CXType_OCLIntelSubgroupAVCSicPayload
	// Type_OCLIntelSubgroupAVCMceResult represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLIntelSubgroupAVCMceResult = C.CXType_OCLIntelSubgroupAVCMceResult
	// Type_OCLIntelSubgroupAVCImeResult represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLIntelSubgroupAVCImeResult = C.CXType_OCLIntelSubgroupAVCImeResult
	// Type_OCLIntelSubgroupAVCRefResult represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLIntelSubgroupAVCRefResult = C.CXType_OCLIntelSubgroupAVCRefResult
	// Type_OCLIntelSubgroupAVCSicResult represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLIntelSubgroupAVCSicResult = C.CXType_OCLIntelSubgroupAVCSicResult
	// Type_OCLIntelSubgroupAVCImeResultSingleRefStreamout represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLIntelSubgroupAVCImeResultSingleRefStreamout = C.CXType_OCLIntelSubgroupAVCImeResultSingleRefStreamout
	// Type_OCLIntelSubgroupAVCImeResultDualRefStreamout represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLIntelSubgroupAVCImeResultDualRefStreamout = C.CXType_OCLIntelSubgroupAVCImeResultDualRefStreamout
	// Type_OCLIntelSubgroupAVCImeSingleRefStreamin represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLIntelSubgroupAVCImeSingleRefStreamin = C.CXType_OCLIntelSubgroupAVCImeSingleRefStreamin
	// Type_OCLIntelSubgroupAVCImeDualRefStreamin represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_OCLIntelSubgroupAVCImeDualRefStreamin = C.CXType_OCLIntelSubgroupAVCImeDualRefStreamin
	// Type_ExtVector represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_ExtVector = C.CXType_ExtVector
	// Type_Atomic represents a type that was referred to using an elaborated type keyword.
	//
	// E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	Type_Atomic = C.CXType_Atomic
)

// GetTypeKindSpelling retrieve the spelling of a given CXTypeKind.
func (tk TypeKind) Spelling() string {
	o := cxstring{C.clang_getTypeKindSpelling(C.enum_CXTypeKind(tk))}
	defer o.Dispose()

	return o.String()
}

func (tk TypeKind) String() string {
	return tk.Spelling()
}
