package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// CursorKind describes the kind of entity that a cursor refers to.
type CursorKind uint32

const (
	// Cursor_UnexposedDecl a declaration whose specific kind is not exposed via this
	// interface.
	//
	// Unexposed declarations have the same operations as any other kind
	// of declaration; one can extract their location information,
	// spelling, find their definitions, etc. However, the specific kind
	// of the declaration is not reported.
	Cursor_UnexposedDecl CursorKind = C.CXCursor_UnexposedDecl
	// Cursor_StructDecl a C or C++ struct.
	Cursor_StructDecl = C.CXCursor_StructDecl
	// Cursor_UnionDecl a C or C++ union.
	Cursor_UnionDecl = C.CXCursor_UnionDecl
	// Cursor_ClassDecl a C++ class.
	Cursor_ClassDecl = C.CXCursor_ClassDecl
	// Cursor_EnumDecl an enumeration.
	Cursor_EnumDecl = C.CXCursor_EnumDecl
	// Cursor_FieldDecl a field (in C) or non-static data member (in C++) in a struct, union, or C++ class.
	Cursor_FieldDecl = C.CXCursor_FieldDecl
	// Cursor_EnumConstantDecl an enumerator constant.
	Cursor_EnumConstantDecl = C.CXCursor_EnumConstantDecl
	// Cursor_FunctionDecl a function.
	Cursor_FunctionDecl = C.CXCursor_FunctionDecl
	// Cursor_VarDecl a variable.
	Cursor_VarDecl = C.CXCursor_VarDecl
	// Cursor_ParmDecl a function or method parameter.
	Cursor_ParmDecl = C.CXCursor_ParmDecl
	// Cursor_ObjCInterfaceDecl an Objective-C \@interface.
	Cursor_ObjCInterfaceDecl = C.CXCursor_ObjCInterfaceDecl
	// Cursor_ObjCCategoryDecl an Objective-C \@interface for a category.
	Cursor_ObjCCategoryDecl = C.CXCursor_ObjCCategoryDecl
	// Cursor_ObjCProtocolDecl an Objective-C \@protocol declaration.
	Cursor_ObjCProtocolDecl = C.CXCursor_ObjCProtocolDecl
	// Cursor_ObjCPropertyDecl an Objective-C \@property declaration.
	Cursor_ObjCPropertyDecl = C.CXCursor_ObjCPropertyDecl
	// Cursor_ObjCIvarDecl an Objective-C instance variable.
	Cursor_ObjCIvarDecl = C.CXCursor_ObjCIvarDecl
	// Cursor_ObjCInstanceMethodDecl an Objective-C instance method.
	Cursor_ObjCInstanceMethodDecl = C.CXCursor_ObjCInstanceMethodDecl
	// Cursor_ObjCClassMethodDecl an Objective-C class method.
	Cursor_ObjCClassMethodDecl = C.CXCursor_ObjCClassMethodDecl
	// Cursor_ObjCImplementationDecl an Objective-C \@implementation.
	Cursor_ObjCImplementationDecl = C.CXCursor_ObjCImplementationDecl
	// Cursor_ObjCCategoryImplDecl an Objective-C \@implementation for a category.
	Cursor_ObjCCategoryImplDecl = C.CXCursor_ObjCCategoryImplDecl
	// Cursor_TypedefDecl a typedef.
	Cursor_TypedefDecl = C.CXCursor_TypedefDecl
	// Cursor_CXXMethod a C++ class method.
	Cursor_CXXMethod = C.CXCursor_CXXMethod
	// Cursor_Namespace a C++ namespace.
	Cursor_Namespace = C.CXCursor_Namespace
	// Cursor_LinkageSpec a linkage specification, e.g. 'extern "C"'.
	Cursor_LinkageSpec = C.CXCursor_LinkageSpec
	// Cursor_Constructor a C++ constructor.
	Cursor_Constructor = C.CXCursor_Constructor
	// Cursor_Destructor a C++ destructor.
	Cursor_Destructor = C.CXCursor_Destructor
	// Cursor_ConversionFunction a C++ conversion function.
	Cursor_ConversionFunction = C.CXCursor_ConversionFunction
	// Cursor_TemplateTypeParameter a C++ template type parameter.
	Cursor_TemplateTypeParameter = C.CXCursor_TemplateTypeParameter
	// Cursor_NonTypeTemplateParameter a C++ non-type template parameter.
	Cursor_NonTypeTemplateParameter = C.CXCursor_NonTypeTemplateParameter
	// Cursor_TemplateTemplateParameter a C++ template template parameter.
	Cursor_TemplateTemplateParameter = C.CXCursor_TemplateTemplateParameter
	// Cursor_FunctionTemplate a C++ function template.
	Cursor_FunctionTemplate = C.CXCursor_FunctionTemplate
	// Cursor_ClassTemplate a C++ class template.
	Cursor_ClassTemplate = C.CXCursor_ClassTemplate
	// Cursor_ClassTemplatePartialSpecialization a C++ class template partial specialization.
	Cursor_ClassTemplatePartialSpecialization = C.CXCursor_ClassTemplatePartialSpecialization
	// Cursor_NamespaceAlias a C++ namespace alias declaration.
	Cursor_NamespaceAlias = C.CXCursor_NamespaceAlias
	// Cursor_UsingDirective a C++ using directive.
	Cursor_UsingDirective = C.CXCursor_UsingDirective
	// Cursor_UsingDeclaration a C++ using declaration.
	Cursor_UsingDeclaration = C.CXCursor_UsingDeclaration
	// Cursor_TypeAliasDecl a C++ alias declaration
	Cursor_TypeAliasDecl = C.CXCursor_TypeAliasDecl
	// Cursor_ObjCSynthesizeDecl an Objective-C \@synthesize definition.
	Cursor_ObjCSynthesizeDecl = C.CXCursor_ObjCSynthesizeDecl
	// Cursor_ObjCDynamicDecl an Objective-C \@dynamic definition.
	Cursor_ObjCDynamicDecl = C.CXCursor_ObjCDynamicDecl
	// Cursor_CXXAccessSpecifier an access specifier.
	Cursor_CXXAccessSpecifier = C.CXCursor_CXXAccessSpecifier
	// Cursor_FirstDecl an access specifier.
	Cursor_FirstDecl = C.CXCursor_FirstDecl
	// Cursor_LastDecl an access specifier.
	Cursor_LastDecl = C.CXCursor_LastDecl
	// Cursor_FirstRef an access specifier.
	Cursor_FirstRef = C.CXCursor_FirstRef
	// Cursor_ObjCSuperClassRef an access specifier.
	Cursor_ObjCSuperClassRef = C.CXCursor_ObjCSuperClassRef
	// Cursor_ObjCProtocolRef an access specifier.
	Cursor_ObjCProtocolRef = C.CXCursor_ObjCProtocolRef
	// Cursor_ObjCClassRef an access specifier.
	Cursor_ObjCClassRef = C.CXCursor_ObjCClassRef
	// Cursor_TypeRef a reference to a type declaration.
	//
	// A type reference occurs anywhere where a type is named but not
	// declared. For example, given:
	//
	//  typedef unsigned size_type;
	//  size_type size;
	//
	// The typedef is a declaration of size_type (CXCursor_TypedefDecl),
	// while the type of the variable "size" is referenced. The cursor
	// referenced by the type of size is the typedef for size_type.
	Cursor_TypeRef = C.CXCursor_TypeRef
	// Cursor_CXXBaseSpecifier a reference to a type declaration.
	//
	// A type reference occurs anywhere where a type is named but not
	// declared. For example, given:
	//
	//  typedef unsigned size_type;
	//  size_type size;
	//
	// The typedef is a declaration of size_type (CXCursor_TypedefDecl),
	// while the type of the variable "size" is referenced. The cursor
	// referenced by the type of size is the typedef for size_type.
	Cursor_CXXBaseSpecifier = C.CXCursor_CXXBaseSpecifier
	// Cursor_TemplateRef a reference to a class template, function template, template template parameter, or class template partial specialization.
	Cursor_TemplateRef = C.CXCursor_TemplateRef
	// Cursor_NamespaceRef a reference to a namespace or namespace alias.
	Cursor_NamespaceRef = C.CXCursor_NamespaceRef
	// Cursor_MemberRef a reference to a member of a struct, union, or class that occurs in some non-expression context, e.g., a designated initializer.
	Cursor_MemberRef = C.CXCursor_MemberRef
	// Cursor_LabelRef a reference to a labeled statement.
	//
	// This cursor kind is used to describe the jump to "start_over" in the
	// goto statement in the following example:
	//
	//  start_over:
	//  ++counter;
	//
	//  goto start_over;
	//
	// A label reference cursor refers to a label statement.
	Cursor_LabelRef = C.CXCursor_LabelRef
	// Cursor_OverloadedDeclRef a reference to a set of overloaded functions or function templates
	// that has not yet been resolved to a specific function or function template.
	//
	// An overloaded declaration reference cursor occurs in C++ templates where
	// a dependent name refers to a function. For example:
	//
	//  template<typename T> void swap(T&, T&);
	//
	//  struct X { ... };
	//  void swap(X&, X&);
	//
	//  template<typename T>
	//  void reverse(T* first, T* last) {
	//  while (first < last - 1) {
	//  swap(*first, *--last);
	//  ++first;
	//  }
	//  }
	//
	//  struct Y { };
	//  void swap(Y&, Y&);
	//
	// Here, the identifier "swap" is associated with an overloaded declaration
	// reference. In the template definition, "swap" refers to either of the two
	// "swap" functions declared above, so both results will be available. At
	// instantiation time, "swap" may also refer to other functions found via
	// argument-dependent lookup (e.g., the "swap" function at the end of the
	// example).
	//
	// The functions clang_getNumOverloadedDecls() and
	// clang_getOverloadedDecl() can be used to retrieve the definitions
	// referenced by this cursor.
	Cursor_OverloadedDeclRef = C.CXCursor_OverloadedDeclRef
	// Cursor_VariableRef a reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_VariableRef = C.CXCursor_VariableRef
	// Cursor_LastRef a reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_LastRef = C.CXCursor_LastRef
	// Cursor_FirstInvalid a reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_FirstInvalid = C.CXCursor_FirstInvalid
	// Cursor_InvalidFile a reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_InvalidFile = C.CXCursor_InvalidFile
	// Cursor_NoDeclFound a reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_NoDeclFound = C.CXCursor_NoDeclFound
	// Cursor_NotImplemented a reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_NotImplemented = C.CXCursor_NotImplemented
	// Cursor_InvalidCode a reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_InvalidCode = C.CXCursor_InvalidCode
	// Cursor_LastInvalid a reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_LastInvalid = C.CXCursor_LastInvalid
	// Cursor_FirstExpr a reference to a variable that occurs in some non-expression context, e.g., a C++ lambda capture list.
	Cursor_FirstExpr = C.CXCursor_FirstExpr
	// Cursor_UnexposedExpr an expression whose specific kind is not exposed via this
	// interface.
	//
	// Unexposed expressions have the same operations as any other kind
	// of expression; one can extract their location information,
	// spelling, children, etc. However, the specific kind of the
	// expression is not reported.
	Cursor_UnexposedExpr = C.CXCursor_UnexposedExpr
	// Cursor_DeclRefExpr an expression that refers to some value declaration, such as a function, variable, or enumerator.
	Cursor_DeclRefExpr = C.CXCursor_DeclRefExpr
	// Cursor_MemberRefExpr an expression that refers to a member of a struct, union, class, Objective-C class, etc.
	Cursor_MemberRefExpr = C.CXCursor_MemberRefExpr
	// Cursor_CallExpr an expression that calls a function.
	Cursor_CallExpr = C.CXCursor_CallExpr
	// Cursor_ObjCMessageExpr an expression that sends a message to an Objective-C  object or class.
	Cursor_ObjCMessageExpr = C.CXCursor_ObjCMessageExpr
	// Cursor_BlockExpr an expression that represents a block literal.
	Cursor_BlockExpr = C.CXCursor_BlockExpr
	// Cursor_IntegerLiteral an integer literal.
	Cursor_IntegerLiteral = C.CXCursor_IntegerLiteral
	// Cursor_FloatingLiteral a floating point number literal.
	Cursor_FloatingLiteral = C.CXCursor_FloatingLiteral
	// Cursor_ImaginaryLiteral an imaginary number literal.
	Cursor_ImaginaryLiteral = C.CXCursor_ImaginaryLiteral
	// Cursor_StringLiteral a string literal.
	Cursor_StringLiteral = C.CXCursor_StringLiteral
	// Cursor_CharacterLiteral a character literal.
	Cursor_CharacterLiteral = C.CXCursor_CharacterLiteral
	// Cursor_ParenExpr a parenthesized expression, e.g. "(1)".
	//
	// This AST node is only formed if full location information is requested.
	Cursor_ParenExpr = C.CXCursor_ParenExpr
	// Cursor_UnaryOperator this represents the unary-expression's (except sizeof and alignof).
	Cursor_UnaryOperator = C.CXCursor_UnaryOperator
	// Cursor_ArraySubscriptExpr [C99 6.5.2.1] Array Subscripting.
	Cursor_ArraySubscriptExpr = C.CXCursor_ArraySubscriptExpr
	// Cursor_BinaryOperator a builtin binary operation expression such as "x + y" or "x <= y".
	Cursor_BinaryOperator = C.CXCursor_BinaryOperator
	// Cursor_CompoundAssignOperator compound assignment such as "+=".
	Cursor_CompoundAssignOperator = C.CXCursor_CompoundAssignOperator
	// Cursor_ConditionalOperator the ?: ternary operator.
	Cursor_ConditionalOperator = C.CXCursor_ConditionalOperator
	// Cursor_CStyleCastExpr an explicit cast in C (C99 6.5.4) or a C-style cast in C++
	// (C++ [expr.cast]), which uses the syntax (Type)expr.
	//
	// For example: (int)f.
	Cursor_CStyleCastExpr = C.CXCursor_CStyleCastExpr
	// Cursor_CompoundLiteralExpr [C99 6.5.2.5]
	Cursor_CompoundLiteralExpr = C.CXCursor_CompoundLiteralExpr
	// Cursor_InitListExpr describes an C or C++ initializer list.
	Cursor_InitListExpr = C.CXCursor_InitListExpr
	// Cursor_AddrLabelExpr the GNU address of label extension, representing &&label.
	Cursor_AddrLabelExpr = C.CXCursor_AddrLabelExpr
	// Cursor_StmtExpr this is the GNU Statement Expression extension: ({int X=4; X;})
	Cursor_StmtExpr = C.CXCursor_StmtExpr
	// Cursor_GenericSelectionExpr represents a C11 generic selection.
	Cursor_GenericSelectionExpr = C.CXCursor_GenericSelectionExpr
	// Cursor_GNUNullExpr implements the GNU __null extension, which is a name for a null
	// pointer constant that has integral type (e.g., int or long) and is the same
	// size and alignment as a pointer.
	//
	// The __null extension is typically only used by system headers, which define
	// NULL as __null in C++ rather than using 0 (which is an integer that may not
	// match the size of a pointer).
	Cursor_GNUNullExpr = C.CXCursor_GNUNullExpr
	// Cursor_CXXStaticCastExpr c++'s static_cast<> expression.
	Cursor_CXXStaticCastExpr = C.CXCursor_CXXStaticCastExpr
	// Cursor_CXXDynamicCastExpr c++'s dynamic_cast<> expression.
	Cursor_CXXDynamicCastExpr = C.CXCursor_CXXDynamicCastExpr
	// Cursor_CXXReinterpretCastExpr c++'s reinterpret_cast<> expression.
	Cursor_CXXReinterpretCastExpr = C.CXCursor_CXXReinterpretCastExpr
	// Cursor_CXXConstCastExpr c++'s const_cast<> expression.
	Cursor_CXXConstCastExpr = C.CXCursor_CXXConstCastExpr
	// Cursor_CXXFunctionalCastExpr represents an explicit C++ type conversion that uses "functional"
	// notion (C++ [expr.type.conv]).
	//
	// Example:
	//  x = int(0.5);
	Cursor_CXXFunctionalCastExpr = C.CXCursor_CXXFunctionalCastExpr
	// Cursor_CXXTypeidExpr a C++ typeid expression (C++ [expr.typeid]).
	Cursor_CXXTypeidExpr = C.CXCursor_CXXTypeidExpr
	// Cursor_CXXBoolLiteralExpr [C++ 2.13.5] C++ Boolean Literal.
	Cursor_CXXBoolLiteralExpr = C.CXCursor_CXXBoolLiteralExpr
	// Cursor_CXXNullPtrLiteralExpr [C++0x 2.14.7] C++ Pointer Literal.
	Cursor_CXXNullPtrLiteralExpr = C.CXCursor_CXXNullPtrLiteralExpr
	// Cursor_CXXThisExpr represents the "this" expression in C++
	Cursor_CXXThisExpr = C.CXCursor_CXXThisExpr
	// Cursor_CXXThrowExpr [C++ 15] C++ Throw Expression.
	//
	// This handles 'throw' and 'throw' assignment-expression. When
	// assignment-expression isn't present, Op will be null.
	Cursor_CXXThrowExpr = C.CXCursor_CXXThrowExpr
	// Cursor_CXXNewExpr a new expression for memory allocation and constructor calls, e.g: "new CXXNewExpr(foo)".
	Cursor_CXXNewExpr = C.CXCursor_CXXNewExpr
	// Cursor_CXXDeleteExpr a delete expression for memory deallocation and destructor calls, e.g. "delete[] pArray".
	Cursor_CXXDeleteExpr = C.CXCursor_CXXDeleteExpr
	// Cursor_UnaryExpr a unary expression. (noexcept, sizeof, or other traits)
	Cursor_UnaryExpr = C.CXCursor_UnaryExpr
	// Cursor_ObjCStringLiteral an Objective-C string literal i.e. @"foo".
	Cursor_ObjCStringLiteral = C.CXCursor_ObjCStringLiteral
	// Cursor_ObjCEncodeExpr an Objective-C \@encode expression.
	Cursor_ObjCEncodeExpr = C.CXCursor_ObjCEncodeExpr
	// Cursor_ObjCSelectorExpr an Objective-C \@selector expression.
	Cursor_ObjCSelectorExpr = C.CXCursor_ObjCSelectorExpr
	// Cursor_ObjCProtocolExpr an Objective-C \@protocol expression.
	Cursor_ObjCProtocolExpr = C.CXCursor_ObjCProtocolExpr
	// Cursor_ObjCBridgedCastExpr an Objective-C "bridged" cast expression, which casts between
	// Objective-C pointers and C pointers, transferring ownership in the process.
	//
	//  NSString *str = (__bridge_transfer NSString *)CFCreateString();
	Cursor_ObjCBridgedCastExpr = C.CXCursor_ObjCBridgedCastExpr
	// Cursor_PackExpansionExpr represents a C++0x pack expansion that produces a sequence of
	// expressions.
	//
	// A pack expansion expression contains a pattern (which itself is an
	// expression) followed by an ellipsis. For example:
	//
	//  template<typename F, typename ...Types>
	//  void forward(F f, Types &&...args) {
	//  f(static_cast<Types&&>(args)...);
	//  }
	Cursor_PackExpansionExpr = C.CXCursor_PackExpansionExpr
	// Cursor_SizeOfPackExpr represents an expression that computes the length of a parameter
	// pack.
	//
	//  template<typename ...Types>
	//  struct count {
	//  static const unsigned value = sizeof...(Types);
	//  };
	Cursor_SizeOfPackExpr = C.CXCursor_SizeOfPackExpr
	Cursor_LambdaExpr     = C.CXCursor_LambdaExpr
	// Cursor_ObjCBoolLiteralExpr objective-c Boolean Literal.
	Cursor_ObjCBoolLiteralExpr = C.CXCursor_ObjCBoolLiteralExpr
	// Cursor_ObjCSelfExpr represents the "self" expression in an Objective-C method.
	Cursor_ObjCSelfExpr = C.CXCursor_ObjCSelfExpr
	// Cursor_OMPArraySectionExpr openMP 5.0 [2.1.5, Array Section].
	Cursor_OMPArraySectionExpr = C.CXCursor_OMPArraySectionExpr
	// Cursor_ObjCAvailabilityCheckExpr represents an @available(...) check.
	Cursor_ObjCAvailabilityCheckExpr = C.CXCursor_ObjCAvailabilityCheckExpr
	// Cursor_FixedPointLiteral fixed point literal
	Cursor_FixedPointLiteral = C.CXCursor_FixedPointLiteral
	// Cursor_OMPArrayShapingExpr openMP 5.0 [2.1.4, Array Shaping].
	Cursor_OMPArrayShapingExpr = C.CXCursor_OMPArrayShapingExpr
	// Cursor_OMPIteratorExpr openMP 5.0 [2.1.6 Iterators]
	Cursor_OMPIteratorExpr = C.CXCursor_OMPIteratorExpr
	// Cursor_CXXAddrspaceCastExpr openCL's addrspace_cast<> expression.
	Cursor_CXXAddrspaceCastExpr = C.CXCursor_CXXAddrspaceCastExpr
	// Cursor_LastExpr openCL's addrspace_cast<> expression.
	Cursor_LastExpr = C.CXCursor_LastExpr
	// Cursor_FirstStmt openCL's addrspace_cast<> expression.
	Cursor_FirstStmt = C.CXCursor_FirstStmt
	// Cursor_UnexposedStmt a statement whose specific kind is not exposed via this
	// interface.
	//
	// Unexposed statements have the same operations as any other kind of
	// statement; one can extract their location information, spelling,
	// children, etc. However, the specific kind of the statement is not
	// reported.
	Cursor_UnexposedStmt = C.CXCursor_UnexposedStmt
	// Cursor_LabelStmt a labelled statement in a function.
	//
	// This cursor kind is used to describe the "start_over:" label statement in
	// the following example:
	//
	//  start_over:
	//  ++counter;
	Cursor_LabelStmt = C.CXCursor_LabelStmt
	// Cursor_CompoundStmt a group of statements like { stmt stmt }.
	//
	// This cursor kind is used to describe compound statements, e.g. function
	// bodies.
	Cursor_CompoundStmt = C.CXCursor_CompoundStmt
	// Cursor_CaseStmt a case statement.
	Cursor_CaseStmt = C.CXCursor_CaseStmt
	// Cursor_DefaultStmt a default statement.
	Cursor_DefaultStmt = C.CXCursor_DefaultStmt
	// Cursor_IfStmt an if statement
	Cursor_IfStmt = C.CXCursor_IfStmt
	// Cursor_SwitchStmt a switch statement.
	Cursor_SwitchStmt = C.CXCursor_SwitchStmt
	// Cursor_WhileStmt a while statement.
	Cursor_WhileStmt = C.CXCursor_WhileStmt
	// Cursor_DoStmt a do statement.
	Cursor_DoStmt = C.CXCursor_DoStmt
	// Cursor_ForStmt a for statement.
	Cursor_ForStmt = C.CXCursor_ForStmt
	// Cursor_GotoStmt a goto statement.
	Cursor_GotoStmt = C.CXCursor_GotoStmt
	// Cursor_IndirectGotoStmt an indirect goto statement.
	Cursor_IndirectGotoStmt = C.CXCursor_IndirectGotoStmt
	// Cursor_ContinueStmt a continue statement.
	Cursor_ContinueStmt = C.CXCursor_ContinueStmt
	// Cursor_BreakStmt a break statement.
	Cursor_BreakStmt = C.CXCursor_BreakStmt
	// Cursor_ReturnStmt a return statement.
	Cursor_ReturnStmt = C.CXCursor_ReturnStmt
	// Cursor_GCCAsmStmt a GCC inline assembly statement extension.
	Cursor_GCCAsmStmt = C.CXCursor_GCCAsmStmt
	// Cursor_AsmStmt a GCC inline assembly statement extension.
	Cursor_AsmStmt = C.CXCursor_AsmStmt
	// Cursor_ObjCAtTryStmt objective-C's overall try-catch-finallystatement.
	Cursor_ObjCAtTryStmt = C.CXCursor_ObjCAtTryStmt
	// Cursor_ObjCAtCatchStmt objective-C's catch statement.
	Cursor_ObjCAtCatchStmt = C.CXCursor_ObjCAtCatchStmt
	// Cursor_ObjCAtFinallyStmt objective-C's finallystatement.
	Cursor_ObjCAtFinallyStmt = C.CXCursor_ObjCAtFinallyStmt
	// Cursor_ObjCAtThrowStmt objective-C's \@throw statement.
	Cursor_ObjCAtThrowStmt = C.CXCursor_ObjCAtThrowStmt
	// Cursor_ObjCAtSynchronizedStmt objective-C's \@synchronized statement.
	Cursor_ObjCAtSynchronizedStmt = C.CXCursor_ObjCAtSynchronizedStmt
	// Cursor_ObjCAutoreleasePoolStmt objective-C's autorelease pool statement.
	Cursor_ObjCAutoreleasePoolStmt = C.CXCursor_ObjCAutoreleasePoolStmt
	// Cursor_ObjCForCollectionStmt objective-C's collection statement.
	Cursor_ObjCForCollectionStmt = C.CXCursor_ObjCForCollectionStmt
	// Cursor_CXXCatchStmt c++'s catch statement.
	Cursor_CXXCatchStmt = C.CXCursor_CXXCatchStmt
	// Cursor_CXXTryStmt c++'s try statement.
	Cursor_CXXTryStmt = C.CXCursor_CXXTryStmt
	// Cursor_CXXForRangeStmt c++'s for (* : *) statement.
	Cursor_CXXForRangeStmt = C.CXCursor_CXXForRangeStmt
	// Cursor_SEHTryStmt windows Structured Exception Handling's try statement.
	Cursor_SEHTryStmt = C.CXCursor_SEHTryStmt
	// Cursor_SEHExceptStmt windows Structured Exception Handling's except statement.
	Cursor_SEHExceptStmt = C.CXCursor_SEHExceptStmt
	// Cursor_SEHFinallyStmt windows Structured Exception Handling's finally statement.
	Cursor_SEHFinallyStmt = C.CXCursor_SEHFinallyStmt
	// Cursor_MSAsmStmt a MS inline assembly statement extension.
	Cursor_MSAsmStmt = C.CXCursor_MSAsmStmt
	// Cursor_NullStmt the null statement ";": C99 6.8.3p3.
	//
	// This cursor kind is used to describe the null statement.
	Cursor_NullStmt = C.CXCursor_NullStmt
	// Cursor_DeclStmt adaptor class for mixing declarations with statements and expressions.
	Cursor_DeclStmt = C.CXCursor_DeclStmt
	// Cursor_OMPParallelDirective openMP parallel directive.
	Cursor_OMPParallelDirective = C.CXCursor_OMPParallelDirective
	// Cursor_OMPSimdDirective openMP SIMD directive.
	Cursor_OMPSimdDirective = C.CXCursor_OMPSimdDirective
	// Cursor_OMPForDirective openMP for directive.
	Cursor_OMPForDirective = C.CXCursor_OMPForDirective
	// Cursor_OMPSectionsDirective openMP sections directive.
	Cursor_OMPSectionsDirective = C.CXCursor_OMPSectionsDirective
	// Cursor_OMPSectionDirective openMP section directive.
	Cursor_OMPSectionDirective = C.CXCursor_OMPSectionDirective
	// Cursor_OMPSingleDirective openMP single directive.
	Cursor_OMPSingleDirective = C.CXCursor_OMPSingleDirective
	// Cursor_OMPParallelForDirective openMP parallel for directive.
	Cursor_OMPParallelForDirective = C.CXCursor_OMPParallelForDirective
	// Cursor_OMPParallelSectionsDirective openMP parallel sections directive.
	Cursor_OMPParallelSectionsDirective = C.CXCursor_OMPParallelSectionsDirective
	// Cursor_OMPTaskDirective openMP task directive.
	Cursor_OMPTaskDirective = C.CXCursor_OMPTaskDirective
	// Cursor_OMPMasterDirective openMP master directive.
	Cursor_OMPMasterDirective = C.CXCursor_OMPMasterDirective
	// Cursor_OMPCriticalDirective openMP critical directive.
	Cursor_OMPCriticalDirective = C.CXCursor_OMPCriticalDirective
	// Cursor_OMPTaskyieldDirective openMP taskyield directive.
	Cursor_OMPTaskyieldDirective = C.CXCursor_OMPTaskyieldDirective
	// Cursor_OMPBarrierDirective openMP barrier directive.
	Cursor_OMPBarrierDirective = C.CXCursor_OMPBarrierDirective
	// Cursor_OMPTaskwaitDirective openMP taskwait directive.
	Cursor_OMPTaskwaitDirective = C.CXCursor_OMPTaskwaitDirective
	// Cursor_OMPFlushDirective openMP flush directive.
	Cursor_OMPFlushDirective = C.CXCursor_OMPFlushDirective
	// Cursor_SEHLeaveStmt windows Structured Exception Handling's leave statement.
	Cursor_SEHLeaveStmt = C.CXCursor_SEHLeaveStmt
	// Cursor_OMPOrderedDirective openMP ordered directive.
	Cursor_OMPOrderedDirective = C.CXCursor_OMPOrderedDirective
	// Cursor_OMPAtomicDirective openMP atomic directive.
	Cursor_OMPAtomicDirective = C.CXCursor_OMPAtomicDirective
	// Cursor_OMPForSimdDirective openMP for SIMD directive.
	Cursor_OMPForSimdDirective = C.CXCursor_OMPForSimdDirective
	// Cursor_OMPParallelForSimdDirective openMP parallel for SIMD directive.
	Cursor_OMPParallelForSimdDirective = C.CXCursor_OMPParallelForSimdDirective
	// Cursor_OMPTargetDirective openMP target directive.
	Cursor_OMPTargetDirective = C.CXCursor_OMPTargetDirective
	// Cursor_OMPTeamsDirective openMP teams directive.
	Cursor_OMPTeamsDirective = C.CXCursor_OMPTeamsDirective
	// Cursor_OMPTaskgroupDirective openMP taskgroup directive.
	Cursor_OMPTaskgroupDirective = C.CXCursor_OMPTaskgroupDirective
	// Cursor_OMPCancellationPointDirective openMP cancellation point directive.
	Cursor_OMPCancellationPointDirective = C.CXCursor_OMPCancellationPointDirective
	// Cursor_OMPCancelDirective openMP cancel directive.
	Cursor_OMPCancelDirective = C.CXCursor_OMPCancelDirective
	// Cursor_OMPTargetDataDirective openMP target data directive.
	Cursor_OMPTargetDataDirective = C.CXCursor_OMPTargetDataDirective
	// Cursor_OMPTaskLoopDirective openMP taskloop directive.
	Cursor_OMPTaskLoopDirective = C.CXCursor_OMPTaskLoopDirective
	// Cursor_OMPTaskLoopSimdDirective openMP taskloop simd directive.
	Cursor_OMPTaskLoopSimdDirective = C.CXCursor_OMPTaskLoopSimdDirective
	// Cursor_OMPDistributeDirective openMP distribute directive.
	Cursor_OMPDistributeDirective = C.CXCursor_OMPDistributeDirective
	// Cursor_OMPTargetEnterDataDirective openMP target enter data directive.
	Cursor_OMPTargetEnterDataDirective = C.CXCursor_OMPTargetEnterDataDirective
	// Cursor_OMPTargetExitDataDirective openMP target exit data directive.
	Cursor_OMPTargetExitDataDirective = C.CXCursor_OMPTargetExitDataDirective
	// Cursor_OMPTargetParallelDirective openMP target parallel directive.
	Cursor_OMPTargetParallelDirective = C.CXCursor_OMPTargetParallelDirective
	// Cursor_OMPTargetParallelForDirective openMP target parallel for directive.
	Cursor_OMPTargetParallelForDirective = C.CXCursor_OMPTargetParallelForDirective
	// Cursor_OMPTargetUpdateDirective openMP target update directive.
	Cursor_OMPTargetUpdateDirective = C.CXCursor_OMPTargetUpdateDirective
	// Cursor_OMPDistributeParallelForDirective openMP distribute parallel for directive.
	Cursor_OMPDistributeParallelForDirective = C.CXCursor_OMPDistributeParallelForDirective
	// Cursor_OMPDistributeParallelForSimdDirective openMP distribute parallel for simd directive.
	Cursor_OMPDistributeParallelForSimdDirective = C.CXCursor_OMPDistributeParallelForSimdDirective
	// Cursor_OMPDistributeSimdDirective openMP distribute simd directive.
	Cursor_OMPDistributeSimdDirective = C.CXCursor_OMPDistributeSimdDirective
	// Cursor_OMPTargetParallelForSimdDirective openMP target parallel for simd directive.
	Cursor_OMPTargetParallelForSimdDirective = C.CXCursor_OMPTargetParallelForSimdDirective
	// Cursor_OMPTargetSimdDirective openMP target simd directive.
	Cursor_OMPTargetSimdDirective = C.CXCursor_OMPTargetSimdDirective
	// Cursor_OMPTeamsDistributeDirective openMP teams distribute directive.
	Cursor_OMPTeamsDistributeDirective = C.CXCursor_OMPTeamsDistributeDirective
	// Cursor_OMPTeamsDistributeSimdDirective openMP teams distribute simd directive.
	Cursor_OMPTeamsDistributeSimdDirective = C.CXCursor_OMPTeamsDistributeSimdDirective
	// Cursor_OMPTeamsDistributeParallelForSimdDirective openMP teams distribute parallel for simd directive.
	Cursor_OMPTeamsDistributeParallelForSimdDirective = C.CXCursor_OMPTeamsDistributeParallelForSimdDirective
	// Cursor_OMPTeamsDistributeParallelForDirective openMP teams distribute parallel for directive.
	Cursor_OMPTeamsDistributeParallelForDirective = C.CXCursor_OMPTeamsDistributeParallelForDirective
	// Cursor_OMPTargetTeamsDirective openMP target teams directive.
	Cursor_OMPTargetTeamsDirective = C.CXCursor_OMPTargetTeamsDirective
	// Cursor_OMPTargetTeamsDistributeDirective openMP target teams distribute directive.
	Cursor_OMPTargetTeamsDistributeDirective = C.CXCursor_OMPTargetTeamsDistributeDirective
	// Cursor_OMPTargetTeamsDistributeParallelForDirective openMP target teams distribute parallel for directive.
	Cursor_OMPTargetTeamsDistributeParallelForDirective = C.CXCursor_OMPTargetTeamsDistributeParallelForDirective
	// Cursor_OMPTargetTeamsDistributeParallelForSimdDirective openMP target teams distribute parallel for simd directive.
	Cursor_OMPTargetTeamsDistributeParallelForSimdDirective = C.CXCursor_OMPTargetTeamsDistributeParallelForSimdDirective
	// Cursor_OMPTargetTeamsDistributeSimdDirective openMP target teams distribute simd directive.
	Cursor_OMPTargetTeamsDistributeSimdDirective = C.CXCursor_OMPTargetTeamsDistributeSimdDirective
	// Cursor_BuiltinBitCastExpr c++2a std::bit_cast expression.
	Cursor_BuiltinBitCastExpr = C.CXCursor_BuiltinBitCastExpr
	// Cursor_OMPMasterTaskLoopDirective openMP master taskloop directive.
	Cursor_OMPMasterTaskLoopDirective = C.CXCursor_OMPMasterTaskLoopDirective
	// Cursor_OMPParallelMasterTaskLoopDirective openMP parallel master taskloop directive.
	Cursor_OMPParallelMasterTaskLoopDirective = C.CXCursor_OMPParallelMasterTaskLoopDirective
	// Cursor_OMPMasterTaskLoopSimdDirective openMP master taskloop simd directive.
	Cursor_OMPMasterTaskLoopSimdDirective = C.CXCursor_OMPMasterTaskLoopSimdDirective
	// Cursor_OMPParallelMasterTaskLoopSimdDirective openMP parallel master taskloop simd directive.
	Cursor_OMPParallelMasterTaskLoopSimdDirective = C.CXCursor_OMPParallelMasterTaskLoopSimdDirective
	// Cursor_OMPParallelMasterDirective openMP parallel master directive.
	Cursor_OMPParallelMasterDirective = C.CXCursor_OMPParallelMasterDirective
	// Cursor_OMPDepobjDirective openMP depobj directive.
	Cursor_OMPDepobjDirective = C.CXCursor_OMPDepobjDirective
	// Cursor_OMPScanDirective openMP scan directive.
	Cursor_OMPScanDirective = C.CXCursor_OMPScanDirective
	// Cursor_OMPTileDirective openMP tile directive.
	Cursor_OMPTileDirective = C.CXCursor_OMPTileDirective
	// Cursor_OMPCanonicalLoop openMP canonical loop.
	Cursor_OMPCanonicalLoop = C.CXCursor_OMPCanonicalLoop
	// Cursor_OMPInteropDirective openMP interop directive.
	Cursor_OMPInteropDirective = C.CXCursor_OMPInteropDirective
	// Cursor_OMPDispatchDirective openMP dispatch directive.
	Cursor_OMPDispatchDirective = C.CXCursor_OMPDispatchDirective
	// Cursor_OMPMaskedDirective openMP masked directive.
	Cursor_OMPMaskedDirective = C.CXCursor_OMPMaskedDirective
	// Cursor_OMPUnrollDirective openMP unroll directive.
	Cursor_OMPUnrollDirective = C.CXCursor_OMPUnrollDirective
	// Cursor_LastStmt openMP unroll directive.
	Cursor_LastStmt = C.CXCursor_LastStmt
	// Cursor_TranslationUnit cursor that represents the translation unit itself.
	//
	// The translation unit cursor exists primarily to act as the root
	// cursor for traversing the contents of a translation unit.
	Cursor_TranslationUnit = C.CXCursor_TranslationUnit
	// Cursor_FirstAttr cursor that represents the translation unit itself.
	//
	// The translation unit cursor exists primarily to act as the root
	// cursor for traversing the contents of a translation unit.
	Cursor_FirstAttr = C.CXCursor_FirstAttr
	// Cursor_UnexposedAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_UnexposedAttr = C.CXCursor_UnexposedAttr
	// Cursor_IBActionAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_IBActionAttr = C.CXCursor_IBActionAttr
	// Cursor_IBOutletAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_IBOutletAttr = C.CXCursor_IBOutletAttr
	// Cursor_IBOutletCollectionAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_IBOutletCollectionAttr = C.CXCursor_IBOutletCollectionAttr
	// Cursor_CXXFinalAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_CXXFinalAttr = C.CXCursor_CXXFinalAttr
	// Cursor_CXXOverrideAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_CXXOverrideAttr = C.CXCursor_CXXOverrideAttr
	// Cursor_AnnotateAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_AnnotateAttr = C.CXCursor_AnnotateAttr
	// Cursor_AsmLabelAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_AsmLabelAttr = C.CXCursor_AsmLabelAttr
	// Cursor_PackedAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_PackedAttr = C.CXCursor_PackedAttr
	// Cursor_PureAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_PureAttr = C.CXCursor_PureAttr
	// Cursor_ConstAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_ConstAttr = C.CXCursor_ConstAttr
	// Cursor_NoDuplicateAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_NoDuplicateAttr = C.CXCursor_NoDuplicateAttr
	// Cursor_CUDAConstantAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_CUDAConstantAttr = C.CXCursor_CUDAConstantAttr
	// Cursor_CUDADeviceAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_CUDADeviceAttr = C.CXCursor_CUDADeviceAttr
	// Cursor_CUDAGlobalAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_CUDAGlobalAttr = C.CXCursor_CUDAGlobalAttr
	// Cursor_CUDAHostAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_CUDAHostAttr = C.CXCursor_CUDAHostAttr
	// Cursor_CUDASharedAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_CUDASharedAttr = C.CXCursor_CUDASharedAttr
	// Cursor_VisibilityAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_VisibilityAttr = C.CXCursor_VisibilityAttr
	// Cursor_DLLExport an attribute whose specific kind is not exposed via this interface.
	Cursor_DLLExport = C.CXCursor_DLLExport
	// Cursor_DLLImport an attribute whose specific kind is not exposed via this interface.
	Cursor_DLLImport = C.CXCursor_DLLImport
	// Cursor_NSReturnsRetained an attribute whose specific kind is not exposed via this interface.
	Cursor_NSReturnsRetained = C.CXCursor_NSReturnsRetained
	// Cursor_NSReturnsNotRetained an attribute whose specific kind is not exposed via this interface.
	Cursor_NSReturnsNotRetained = C.CXCursor_NSReturnsNotRetained
	// Cursor_NSReturnsAutoreleased an attribute whose specific kind is not exposed via this interface.
	Cursor_NSReturnsAutoreleased = C.CXCursor_NSReturnsAutoreleased
	// Cursor_NSConsumesSelf an attribute whose specific kind is not exposed via this interface.
	Cursor_NSConsumesSelf = C.CXCursor_NSConsumesSelf
	// Cursor_NSConsumed an attribute whose specific kind is not exposed via this interface.
	Cursor_NSConsumed = C.CXCursor_NSConsumed
	// Cursor_ObjCException an attribute whose specific kind is not exposed via this interface.
	Cursor_ObjCException = C.CXCursor_ObjCException
	// Cursor_ObjCNSObject an attribute whose specific kind is not exposed via this interface.
	Cursor_ObjCNSObject = C.CXCursor_ObjCNSObject
	// Cursor_ObjCIndependentClass an attribute whose specific kind is not exposed via this interface.
	Cursor_ObjCIndependentClass = C.CXCursor_ObjCIndependentClass
	// Cursor_ObjCPreciseLifetime an attribute whose specific kind is not exposed via this interface.
	Cursor_ObjCPreciseLifetime = C.CXCursor_ObjCPreciseLifetime
	// Cursor_ObjCReturnsInnerPointer an attribute whose specific kind is not exposed via this interface.
	Cursor_ObjCReturnsInnerPointer = C.CXCursor_ObjCReturnsInnerPointer
	// Cursor_ObjCRequiresSuper an attribute whose specific kind is not exposed via this interface.
	Cursor_ObjCRequiresSuper = C.CXCursor_ObjCRequiresSuper
	// Cursor_ObjCRootClass an attribute whose specific kind is not exposed via this interface.
	Cursor_ObjCRootClass = C.CXCursor_ObjCRootClass
	// Cursor_ObjCSubclassingRestricted an attribute whose specific kind is not exposed via this interface.
	Cursor_ObjCSubclassingRestricted = C.CXCursor_ObjCSubclassingRestricted
	// Cursor_ObjCExplicitProtocolImpl an attribute whose specific kind is not exposed via this interface.
	Cursor_ObjCExplicitProtocolImpl = C.CXCursor_ObjCExplicitProtocolImpl
	// Cursor_ObjCDesignatedInitializer an attribute whose specific kind is not exposed via this interface.
	Cursor_ObjCDesignatedInitializer = C.CXCursor_ObjCDesignatedInitializer
	// Cursor_ObjCRuntimeVisible an attribute whose specific kind is not exposed via this interface.
	Cursor_ObjCRuntimeVisible = C.CXCursor_ObjCRuntimeVisible
	// Cursor_ObjCBoxable an attribute whose specific kind is not exposed via this interface.
	Cursor_ObjCBoxable = C.CXCursor_ObjCBoxable
	// Cursor_FlagEnum an attribute whose specific kind is not exposed via this interface.
	Cursor_FlagEnum = C.CXCursor_FlagEnum
	// Cursor_ConvergentAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_ConvergentAttr = C.CXCursor_ConvergentAttr
	// Cursor_WarnUnusedAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_WarnUnusedAttr = C.CXCursor_WarnUnusedAttr
	// Cursor_WarnUnusedResultAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_WarnUnusedResultAttr = C.CXCursor_WarnUnusedResultAttr
	// Cursor_AlignedAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_AlignedAttr = C.CXCursor_AlignedAttr
	// Cursor_LastAttr an attribute whose specific kind is not exposed via this interface.
	Cursor_LastAttr = C.CXCursor_LastAttr
	// Cursor_PreprocessingDirective an attribute whose specific kind is not exposed via this interface.
	Cursor_PreprocessingDirective = C.CXCursor_PreprocessingDirective
	// Cursor_MacroDefinition an attribute whose specific kind is not exposed via this interface.
	Cursor_MacroDefinition = C.CXCursor_MacroDefinition
	// Cursor_MacroExpansion an attribute whose specific kind is not exposed via this interface.
	Cursor_MacroExpansion = C.CXCursor_MacroExpansion
	// Cursor_MacroInstantiation an attribute whose specific kind is not exposed via this interface.
	Cursor_MacroInstantiation = C.CXCursor_MacroInstantiation
	// Cursor_InclusionDirective an attribute whose specific kind is not exposed via this interface.
	Cursor_InclusionDirective = C.CXCursor_InclusionDirective
	// Cursor_FirstPreprocessing an attribute whose specific kind is not exposed via this interface.
	Cursor_FirstPreprocessing = C.CXCursor_FirstPreprocessing
	// Cursor_LastPreprocessing an attribute whose specific kind is not exposed via this interface.
	Cursor_LastPreprocessing = C.CXCursor_LastPreprocessing
	// Cursor_ModuleImportDecl a module import declaration.
	Cursor_ModuleImportDecl = C.CXCursor_ModuleImportDecl
	// Cursor_TypeAliasTemplateDecl a module import declaration.
	Cursor_TypeAliasTemplateDecl = C.CXCursor_TypeAliasTemplateDecl
	// Cursor_StaticAssert a static_assert or _Static_assert node
	Cursor_StaticAssert = C.CXCursor_StaticAssert
	// Cursor_FriendDecl a friend declaration.
	Cursor_FriendDecl = C.CXCursor_FriendDecl
	// Cursor_FirstExtraDecl a friend declaration.
	Cursor_FirstExtraDecl = C.CXCursor_FirstExtraDecl
	// Cursor_LastExtraDecl a friend declaration.
	Cursor_LastExtraDecl = C.CXCursor_LastExtraDecl
	// Cursor_OverloadCandidate a code completion overload candidate.
	Cursor_OverloadCandidate = C.CXCursor_OverloadCandidate
)

// IsDeclaration determine whether the given cursor kind represents a declaration.
func (ck CursorKind) IsDeclaration() bool {
	o := C.clang_isDeclaration(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// IsReference determine whether the given cursor kind represents a simple
// reference.
//
// Note that other kinds of cursors (such as expressions) can also refer to
// other cursors. Use clang_getCursorReferenced() to determine whether a
// particular cursor refers to another entity.
func (ck CursorKind) IsReference() bool {
	o := C.clang_isReference(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// IsExpression determine whether the given cursor kind represents an expression.
func (ck CursorKind) IsExpression() bool {
	o := C.clang_isExpression(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// IsStatement determine whether the given cursor kind represents a statement.
func (ck CursorKind) IsStatement() bool {
	o := C.clang_isStatement(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// IsAttribute determine whether the given cursor kind represents an attribute.
func (ck CursorKind) IsAttribute() bool {
	o := C.clang_isAttribute(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// IsInvalid determine whether the given cursor kind represents an invalid cursor.
func (ck CursorKind) IsInvalid() bool {
	o := C.clang_isInvalid(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// IsTranslationUnit determine whether the given cursor kind represents a translation unit.
func (ck CursorKind) IsTranslationUnit() bool {
	o := C.clang_isTranslationUnit(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// IsPreprocessing * Determine whether the given cursor represents a preprocessing element, such as a preprocessor directive or macro instantiation.
func (ck CursorKind) IsPreprocessing() bool {
	o := C.clang_isPreprocessing(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

// IsUnexposed * Determine whether the given cursor represents a currently unexposed piece of the AST (e.g., CXCursor_UnexposedStmt).
func (ck CursorKind) IsUnexposed() bool {
	o := C.clang_isUnexposed(C.enum_CXCursorKind(ck))

	return o != C.uint(0)
}

func (ck CursorKind) Spelling() string {
	o := cxstring{C.clang_getCursorKindSpelling(C.enum_CXCursorKind(ck))}
	defer o.Dispose()

	return o.String()
}

func (ck CursorKind) String() string {
	return ck.Spelling()
}
