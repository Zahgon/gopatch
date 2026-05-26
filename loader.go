package main

import (
	"go/token"
	"io"

	"github.com/uber-go/gopatch/internal/engine"
)

// patchLoader loads patches from varying sources
// and compiles them into a series of programs.
type patchLoader struct {
	fset  *token.FileSet
	progs []*engine.Program

	// Pointer to parseAndCompile function,
	// which we can use to swap out this logic.
	parseAndCompile func(*token.FileSet, string, []byte) (*engine.Program, error)
}

func newPatchLoader(fset *token.FileSet) *patchLoader { _ = "STUB: not implemented"; return nil }

func (l *patchLoader) Programs() []*engine.Program {
	_ = "STUB: not implemented"

	// LoadReader loads a patch from an io.Reader.
	return nil
}

func (l *patchLoader) LoadReader(name string, r io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadFile loads a patch from the given file.
func (l *patchLoader) LoadFile(path string) (err error) { _ = "STUB: not implemented"; return nil }

// LoadFileList loads patches specified in a file
// that contains a list of file paths to other patches.
func (l *patchLoader) LoadFileList(patchList string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// parseAndCompile parses the given patch contents,
// and compiles them into a gopatch program.
func parseAndCompile(fset *token.FileSet, name string, src []byte) (*engine.Program, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
