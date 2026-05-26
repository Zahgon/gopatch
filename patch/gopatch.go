package patch

import (
	"go/ast"
	"go/token"

	"github.com/uber-go/gopatch/internal/engine"
)

// File is a patch difference file that can be applied to Go file.
type File struct {
	fset *token.FileSet
	prog *engine.Program
}

// Parse the patch file and creates data that can be applied to the Go file.
func Parse(patchFileName string, src []byte) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply takes the Go file name and its contents and returns a Go file with the patch applied.
func (f *File) Apply(filename string, src []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This patch didn't modify the file. Try the next one.

func cleanupFilePos(tfile *token.File, cl engine.Changelog, comments []*ast.CommentGroup) {
	_ = "STUB: not implemented"
	return
}

// Remove comments in the changed sections of the code.
