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

package engine

import (
	"go/token"
	"reflect"

	"github.com/uber-go/gopatch/internal/data"
)

// Replacer generates portions of the Go AST meant to replace sections matched
// by a Matcher. A Replacer is built from the "+" portion of a patch.
type Replacer interface {
	// Replace generates values for a Go AST provided prior match data.
	Replace(d data.Data, cl Changelog, pos token.Pos) (v reflect.Value, err error)
}

// replacerCompiler compiles the "+" portion of a patch into a Replacer which
// will generate the portions to fill in the original AST.
type replacerCompiler struct {
	fset     *token.FileSet
	meta     *Meta
	dots     []token.Pos
	dotAssoc map[token.Pos]token.Pos

	patchStart, patchEnd token.Pos
}

func newReplacerCompiler(fset *token.FileSet, meta *Meta, patchStart, patchEnd token.Pos) *replacerCompiler {
	_ = "STUB: not implemented"
	return nil
}

func (c *replacerCompiler) compile(v reflect.Value) Replacer {
	_ = "STUB: not implemented"
	return *new(Replacer)
}

// TODO(abg): pgo.Parse should probably replace this with a DotsField.

// TODO: We're currently ignoring comments in the replacement patch.
// We should probably record them and report them in the top-level
// file.

// Ident.Obj forms a cycle so we'll replace it with a nil pointer.

// ZeroReplacer replaces with a zero value.
type ZeroReplacer struct{ Type reflect.Type }

// Replace replaces with a zero value.
func (r ZeroReplacer) Replace(data.Data, Changelog, token.Pos) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}
