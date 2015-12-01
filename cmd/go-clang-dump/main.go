// go-clang-dump shows how to dump the AST of a C/C++ file via the Cursor
// visitor API.
//
// ex:
// $ go-clang-dump -fname=foo.cxx
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-clang/bootstrap/clang"
)

var fname = flag.String("fname", "", "the file to analyze")

func main() {
	fmt.Printf(":: go-clang-dump...\n")
	flag.Parse()
	fmt.Printf(":: fname: %s\n", *fname)
	fmt.Printf(":: args: %v\n", flag.Args())

	if *fname == "" {
		flag.Usage()
		fmt.Printf("please provide a file name to analyze\n")

		os.Exit(1)
	}

	idx := clang.NewIndex(0, 1)
	defer idx.Dispose()

	args := []string{}
	if len(flag.Args()) > 0 && flag.Args()[0] == "-" {
		args = make([]string, len(flag.Args()[1:]))
		copy(args, flag.Args()[1:])
	}

	tu := idx.ParseTranslationUnit(*fname, args, nil, 0)
	defer tu.Dispose()

	fmt.Printf("tu: %s\n", tu.Spelling())
	cursor := tu.TranslationUnitCursor()
	fmt.Printf("cursor-isnull: %v\n", cursor.IsNull())
	fmt.Printf("cursor: %s\n", cursor.Spelling())
	fmt.Printf("cursor-kind: %s\n", cursor.Kind().Spelling())

	fmt.Printf("tu-fname: %s\n", tu.File(*fname).Name())

	cursor.Visit(func(cursor, parent clang.Cursor) clang.ChildVisitResult {
		if cursor.IsNull() {
			fmt.Printf("cursor: <none>\n")

			return clang.ChildVisit_Continue
		}

		fmt.Printf("%s: %s (%s)\n", cursor.Kind().Spelling(), cursor.Spelling(), cursor.USR())

		switch cursor.Kind() {
		case clang.Cursor_ClassDecl, clang.Cursor_EnumDecl, clang.Cursor_StructDecl, clang.Cursor_Namespace:
			return clang.ChildVisit_Recurse
		}

		return clang.ChildVisit_Continue
	})

	fmt.Printf(":: bye.\n")
}
