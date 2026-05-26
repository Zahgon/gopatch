// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package pgo

import (
	"go/ast"
	"go/token"

	"github.com/uber-go/gopatch/internal/pgo/augment"
)

// Parse parses a pgo file. AST nodes in the returned File reference a newly
// added token.File in the given FileSet.
func Parse(fset *token.FileSet, filename string, src []byte) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: The current system doesn't work for just top-level types.
// For example, you can't parse just "[]T" with this because the generated
// "func _fake() { []T }" is not valid Go code.
//
// One option is to use parser.ParseExpr for top-level expressions, in
// which case augmentations, imports, etc. have to be handled separately.
// So augment should probably return a File object with package name
// already parsed, text for the unparsed imports, and the source. We can
// then try ParseExpr on the source.

// We'll use adjuster to map positions from the augmented file back to the
// user-provided file.

// f.Decls includes import declarations.

// FakePackage will always be the first augmentation if it's there. If we
// have a FakePackage, we should ignore the parsed package name.

// We allow only one declaration in the patch.

// This is impossible. Even for an empty string, we'll generate an
// empty function which will translate to a declaration.

// This is what we expect.

// FakeFunc will always be the next augmentation if it's there. If we're
// using a FakeFunc, the part of the AST under consideration in ast.Source
// is the body of the function.

// If the body contains a single expression, it was from a
// top-level expression.

// If the body is empty, the user didn't provide anything
// after the package/imports.

// TODO(abg): This should be the position after the package
// and imports.

// TODO(abg): We should support zero declarations for
// import/package-only transforms.

// If the body contains a single expression, we want to do an
// expression transformation.

// TODO: check for empty BlockStmt

// We get back an AST with positions inside the .augmented file. Map them
// back to the pgoFile we created above.

func filterImports(decls []ast.Decl) []ast.Decl { _ = "STUB: not implemented"; return nil }

// posAdjuster maps positions in the augmented file back to the unaugmented
// patch file.
type posAdjuster struct {
	Fset *token.FileSet
	File *token.File
	Adjs []augment.PosAdjustment // sorted by offset
}

// Pos maps the provided position to the File associated with the Adjuster,
// using the provided PosAdjustments to offset it appropriately.
func (a *posAdjuster) Pos(pos token.Pos) token.Pos {
	_ = "STUB: not implemented"
	return *new(token.Pos)
}

// Position returns the full positional information for the given token.Pos
// after adjustment.
func (a *posAdjuster) Position(pos token.Pos) token.Position {
	_ = "STUB: not implemented"
	return *new(token.Position)
}

type byOffset []augment.PosAdjustment

func (a byOffset) Len() int { _ = "STUB: not implemented"; return 0 }

func (a byOffset) Less(i int, j int) bool { _ = "STUB: not implemented"; return false }

func (a byOffset) Swap(i int, j int) { _ = "STUB: not implemented"; return }
