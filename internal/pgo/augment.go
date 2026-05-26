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
	"golang.org/x/tools/go/ast/astutil"
)

// augmentAST replaces portions of this AST with pgo nodes based on the
// provided augmentations.
func augmentAST(file *token.File, n ast.Node, augs []augment.Augmentation, adj *posAdjuster) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

// Augmenter is used to traverse the Go AST and replace sections of it.
type augmenter struct {
	file   *token.File
	errors []error

	// Used to adjust positions back to the original file.
	adj *posAdjuster

	// augmentations are indexed by position so that we can match them to
	// nodes as we traverse.
	augs map[augPos]augment.Augmentation
}

// augPos is the position of an augmentation in the file.
type augPos struct {
	// Start and end offsets of the augmentation.
	start, end int
}

func newAugmenter(file *token.File, augs []augment.Augmentation, adj *posAdjuster) *augmenter {
	_ = "STUB: not implemented"
	return nil
}

// errf logs errors at the given position. Position must be in the augmented
// file, not the original file.
func (a *augmenter) errf(pos token.Pos, msg string, args ...any) { _ = "STUB: not implemented"; return }

// pop retrieves the augmentation associated with the provided node, or nil if
// no augmentation applies to this node. If an augmentation was found, it is
// removed from the list of known augmentations.
func (a *augmenter) pop(n ast.Node) augment.Augmentation {
	_ = "STUB: not implemented"
	return *new(augment.Augmentation)
}

// Augmentations are deleted so that we can track unused augmentations
// later.

// Apply visits the given AST cursor. Use this with astutil.Apply.
func (a *augmenter) Apply(cursor *astutil.Cursor) bool { _ = "STUB: not implemented"; return false }

// keep looking

// We can't rely on cursor.Replace because we want to inspect if the type
// of the target field matches the kind of expression we are placing
// there. We can't do this with type switches because for interfaces, type
// switches will match if the target implements the interface.
//
// For example, for a ast.TypeSpec, the Name field is of type *ast.Ident.
// To determine if we can place a &Dots{} there, we need ask the type
// system if that field is an ast.Expr (which it's not). The type system
// will respond with yes because *ast.Ident implements ast.Expr.
//
// Instead, we'll use reflection to inspect the type of the field and then
// let cursor.Replace handle it.

// Err returns an error if this augmenter encountered any errors or if any of
// the augmentations were unused.
func (a *augmenter) Err() error { _ = "STUB: not implemented"; return nil }
