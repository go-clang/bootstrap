package main

import (
	"testing"
)

func TestGoClangDump(t *testing.T) {
	for _, fname := range []string{
		"../../testdata/basicparsing.c",
	} {
		if got, want := cmd([]string{"-fname", fname}), 0; got != want {
			t.Fatalf("got %v but want %v", got, want)
		}
	}
}
