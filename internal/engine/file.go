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
	"github.com/uber-go/gopatch/internal/pgo"
)

// FileMatcher matches Go files.
type FileMatcher struct {
	// Matches the package name, if any.
	Package string

	// Imports in the file.
	Imports ImportsMatcher

	// Matches nodes in the file.
	NodeMatcher Matcher
}

func (c *matcherCompiler) compileFile(file *pgo.File) FileMatcher {
	_ = "STUB: not implemented"
	return *new(FileMatcher)
}

// Match matches against the file, recording information about all matches
// found in it.
func (m FileMatcher) Match(file *ast.File, d data.Data) (data.Data, bool) {
	_ = "STUB: not implemented"
	// Match package name.
	return *new(data.Data), false
}

// TODO(abg): Use an identMatcher with a constraint.

// To match the body, we use astutil.Apply which traverses the AST and
// provides a replaceable pointer to each node so that we can rewrite
// the AST in-place.

// TODO(abg): Support nil NodeMatcher for when a patch is matching on
// just the package name or import paths.

// don't change outer d

// keep looking
/* post func */

// FileReplacer replaces an ast.File.
type FileReplacer struct {
	Fset *token.FileSet

	// Package name to change to, if any.
	Package string

	// Imports in the file.
	Imports ImportsReplacer

	// Replaces matched nodes in the file.
	NodeReplacer Replacer
}

func (c *replacerCompiler) compileFile(file *pgo.File) FileReplacer {
	_ = "STUB: not implemented"
	return *new(FileReplacer)
}

// Replace replaces a file using the provided Match data.
func (r FileReplacer) Replace(d data.Data, cl Changelog) (*ast.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is a bug in our code.

// If the generated value isn't assignable to the target, the match
// was too eager. For example, trying to place "foo.Bar"
// (SelectorExpr) where only an identifier is allowed (in a variable
// declaration name, for example).

type _fileMatchKey struct{}

var fileMatchKey _fileMatchKey

type fileMatchData struct {
	File    *ast.File
	Matches []*SearchResult
}
