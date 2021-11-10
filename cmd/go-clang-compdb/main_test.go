package main

import (
	"testing"
)

func TestGoClangCompDB(t *testing.T) {
	for _, path := range []string{
		"../../testdata",
	} {
		if got, want := cmd([]string{path}), 0; got != want {
			t.Fatal("got %v but want %v", got, want)
		}
	}
}
