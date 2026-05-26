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

	"github.com/uber-go/gopatch/internal/parse"
)

// Program is a collection of compiled changes.
type Program struct {
	Changes []*Change
}

// Compile compiles a parsed gopatch Program.
func Compile(fset *token.FileSet, p *parse.Program) (*Program, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type compiler struct {
	fset   *token.FileSet
	errors []error
}

func newCompiler(fset *token.FileSet) *compiler { _ = "STUB: not implemented"; return nil }

// Convenience function to build error messages with positioning data.
func (c *compiler) errf(pos token.Pos, msg string, args ...any) { _ = "STUB: not implemented"; return }

// Err collates all the errors encountered during compilation and returns
// them.
func (c *compiler) Err() error { _ = "STUB: not implemented"; return nil }

// Compiles a Program.
func (c *compiler) compileProgram(aprogram *parse.Program) *Program {
	_ = "STUB: not implemented"
	return nil
}
