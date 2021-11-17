package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

/*
	Properties for the printing policy.

	See clang::PrintingPolicy for more information.
*/
type PrintingPolicyProperty uint32

const (
	PrintingPolicy_Indentation                           PrintingPolicyProperty = C.CXPrintingPolicy_Indentation
	PrintingPolicy_SuppressSpecifiers                                           = C.CXPrintingPolicy_SuppressSpecifiers
	PrintingPolicy_SuppressTagKeyword                                           = C.CXPrintingPolicy_SuppressTagKeyword
	PrintingPolicy_IncludeTagDefinition                                         = C.CXPrintingPolicy_IncludeTagDefinition
	PrintingPolicy_SuppressScope                                                = C.CXPrintingPolicy_SuppressScope
	PrintingPolicy_SuppressUnwrittenScope                                       = C.CXPrintingPolicy_SuppressUnwrittenScope
	PrintingPolicy_SuppressInitializers                                         = C.CXPrintingPolicy_SuppressInitializers
	PrintingPolicy_ConstantArraySizeAsWritten                                   = C.CXPrintingPolicy_ConstantArraySizeAsWritten
	PrintingPolicy_AnonymousTagLocations                                        = C.CXPrintingPolicy_AnonymousTagLocations
	PrintingPolicy_SuppressStrongLifetime                                       = C.CXPrintingPolicy_SuppressStrongLifetime
	PrintingPolicy_SuppressLifetimeQualifiers                                   = C.CXPrintingPolicy_SuppressLifetimeQualifiers
	PrintingPolicy_SuppressTemplateArgsInCXXConstructors                        = C.CXPrintingPolicy_SuppressTemplateArgsInCXXConstructors
	PrintingPolicy_Bool                                                         = C.CXPrintingPolicy_Bool
	PrintingPolicy_Restrict                                                     = C.CXPrintingPolicy_Restrict
	PrintingPolicy_Alignof                                                      = C.CXPrintingPolicy_Alignof
	PrintingPolicy_UnderscoreAlignof                                            = C.CXPrintingPolicy_UnderscoreAlignof
	PrintingPolicy_UseVoidForZeroParams                                         = C.CXPrintingPolicy_UseVoidForZeroParams
	PrintingPolicy_TerseOutput                                                  = C.CXPrintingPolicy_TerseOutput
	PrintingPolicy_PolishForDeclaration                                         = C.CXPrintingPolicy_PolishForDeclaration
	PrintingPolicy_Half                                                         = C.CXPrintingPolicy_Half
	PrintingPolicy_MSWChar                                                      = C.CXPrintingPolicy_MSWChar
	PrintingPolicy_IncludeNewlines                                              = C.CXPrintingPolicy_IncludeNewlines
	PrintingPolicy_MSVCFormatting                                               = C.CXPrintingPolicy_MSVCFormatting
	PrintingPolicy_ConstantsAsWritten                                           = C.CXPrintingPolicy_ConstantsAsWritten
	PrintingPolicy_SuppressImplicitBase                                         = C.CXPrintingPolicy_SuppressImplicitBase
	PrintingPolicy_FullyQualifiedName                                           = C.CXPrintingPolicy_FullyQualifiedName
	PrintingPolicy_LastProperty                                                 = C.CXPrintingPolicy_LastProperty
)

func (ppp PrintingPolicyProperty) Spelling() string {
	switch ppp {
	case PrintingPolicy_Indentation:
		return "PrintingPolicy=Indentation"
	case PrintingPolicy_SuppressSpecifiers:
		return "PrintingPolicy=SuppressSpecifiers"
	case PrintingPolicy_SuppressTagKeyword:
		return "PrintingPolicy=SuppressTagKeyword"
	case PrintingPolicy_IncludeTagDefinition:
		return "PrintingPolicy=IncludeTagDefinition"
	case PrintingPolicy_SuppressScope:
		return "PrintingPolicy=SuppressScope"
	case PrintingPolicy_SuppressUnwrittenScope:
		return "PrintingPolicy=SuppressUnwrittenScope"
	case PrintingPolicy_SuppressInitializers:
		return "PrintingPolicy=SuppressInitializers"
	case PrintingPolicy_ConstantArraySizeAsWritten:
		return "PrintingPolicy=ConstantArraySizeAsWritten"
	case PrintingPolicy_AnonymousTagLocations:
		return "PrintingPolicy=AnonymousTagLocations"
	case PrintingPolicy_SuppressStrongLifetime:
		return "PrintingPolicy=SuppressStrongLifetime"
	case PrintingPolicy_SuppressLifetimeQualifiers:
		return "PrintingPolicy=SuppressLifetimeQualifiers"
	case PrintingPolicy_SuppressTemplateArgsInCXXConstructors:
		return "PrintingPolicy=SuppressTemplateArgsInCXXConstructors"
	case PrintingPolicy_Bool:
		return "PrintingPolicy=Bool"
	case PrintingPolicy_Restrict:
		return "PrintingPolicy=Restrict"
	case PrintingPolicy_Alignof:
		return "PrintingPolicy=Alignof"
	case PrintingPolicy_UnderscoreAlignof:
		return "PrintingPolicy=UnderscoreAlignof"
	case PrintingPolicy_UseVoidForZeroParams:
		return "PrintingPolicy=UseVoidForZeroParams"
	case PrintingPolicy_TerseOutput:
		return "PrintingPolicy=TerseOutput"
	case PrintingPolicy_PolishForDeclaration:
		return "PrintingPolicy=PolishForDeclaration"
	case PrintingPolicy_Half:
		return "PrintingPolicy=Half"
	case PrintingPolicy_MSWChar:
		return "PrintingPolicy=MSWChar"
	case PrintingPolicy_IncludeNewlines:
		return "PrintingPolicy=IncludeNewlines"
	case PrintingPolicy_MSVCFormatting:
		return "PrintingPolicy=MSVCFormatting"
	case PrintingPolicy_ConstantsAsWritten:
		return "PrintingPolicy=ConstantsAsWritten"
	case PrintingPolicy_SuppressImplicitBase:
		return "PrintingPolicy=SuppressImplicitBase"
	case PrintingPolicy_FullyQualifiedName:
		return "PrintingPolicy=FullyQualifiedName, LastProperty"
	}

	return fmt.Sprintf("PrintingPolicyProperty unknown %d", int(ppp))
}

func (ppp PrintingPolicyProperty) String() string {
	return ppp.Spelling()
}
