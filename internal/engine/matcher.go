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
	"go/ast"
	"go/token"
	"reflect"

	"github.com/uber-go/gopatch/internal/data"
)

// Region denotes the portion of the code being matched, i.e. the start and end
// position of the given Node.
type Region struct{ Pos, End token.Pos }

// nodeRegion returns the Region occupied by a given node.
func nodeRegion(n ast.Node) Region { _ = "STUB: not implemented"; return *new(Region) }

// Matcher matches values in a Go AST. It is built from the "-" portion of a
// patch.
type Matcher interface {
	// Match is called with the a value from the AST being compared with the
	// match data captured so far.
	//
	// Match reports whether the match succeeded and if so, returns the
	// original or a different Data object containing additional match data.
	Match(reflect.Value, data.Data, Region) (_ data.Data, ok bool)
}

// matcherCompiler compiles the "-" portion of a patch into a Matcher which
// will report whether another Go AST matches it.
type matcherCompiler struct {
	fset *token.FileSet
	meta *Meta

	// All dots found during match compilation.
	dots []token.Pos

	patchStart, patchEnd token.Pos
}

func newMatcherCompiler(fset *token.FileSet, meta *Meta, patchStart, patchEnd token.Pos) *matcherCompiler {
	_ = "STUB: not implemented"
	return nil
}

func (c *matcherCompiler) compile(v reflect.Value) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

// TODO(abg): pgo.Parse should probably replace this with a DotsField.

// TODO: Dedupe

// Comments shouldn't affect match.

// Ident.Obj forms a cycle. We'll consider Object pointers to always
// match because the entites they point to will be matched separately
// anyway.

type matcherFunc func(reflect.Value) bool

func (f matcherFunc) Match(v reflect.Value, d data.Data, _ Region) (data.Data, bool) {
	_ = "STUB: not implemented"
	return *

	// nilMatcher is a Matcher that only matches nil values.
	new(data.Data), false
}

var (
	nilMatcher Matcher = matcherFunc(func(got reflect.Value) bool { return got.IsNil() })

	// successMatcher always return true.
	successMatcher Matcher = matcherFunc(func(reflect.Value) bool { return true })
)
