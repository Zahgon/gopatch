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

	"github.com/uber-go/gopatch/internal/data"
	"github.com/uber-go/gopatch/internal/parse"
)

// Change is a single Change in a program.
type Change struct {
	Name string
	Meta *Meta

	Comments []string
	fset     *token.FileSet
	matcher  FileMatcher
	replacer FileReplacer
}

func (c *compiler) compileChange(achange *parse.Change) *Change {
	_ = "STUB: not implemented"
	return nil
}

// TODO(abg): validate name

// Match matches this change in the given Go AST and returns captured match
// information it a data.Data object.
func (c *Change) Match(f *ast.File) (d data.Data, ok bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}

// Replace generates a replacement File based on previously captured match
// data.
func (c *Change) Replace(d data.Data, cl Changelog) (*ast.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func connectDots(fset *token.FileSet, lhs, rhs []token.Pos, conns map[token.Pos]token.Pos) error {
	_ = "STUB: not implemented"
	return nil
}

// Descending order.

// Ascending order.
