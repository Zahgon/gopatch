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
)

// ImportMatcher matches a single import.
type ImportMatcher struct {
	NameS         string  // TODO: better naming
	Name          Matcher // named import, if any
	NameIsMetavar bool    // records wehther the named import was a metavariable
	// TODO: This is janky.

	Path string // import path as a string
}

func (c *matcherCompiler) compileImport(imp *ast.ImportSpec) ImportMatcher {
	_ = "STUB: not implemented"
	return *new(ImportMatcher)
}

// TODO: In the future, we should try to determine the package name
// automatically if the named import is not provided. This is
// expensive because we'd have to resolve the package name from the
// import path, so the code would have to actually exist on-disk.

// Match matches an import in a file. If the matcher was built with a named
// import where the name is a metavariable, then this will match both, named
// and unnamed imports and record that information in the patch data.
func (m ImportMatcher) Match(file *ast.File, d data.Data) (_ data.Data, ok bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}

// We need to account for four cases here:
//
// +--------------+-------------+-----------------------+
// | Patch import | File import | Behavior              |
// +--------------+-------------+-----------------------+
// | unnamed      | unnamed     | match                 |
// | unnamed      | named       | no match              |
// | named        | unnamed     | match if metavariable |
// | named        | named       | match name            |
// +--------------+-------------+-----------------------+

// Patch import is unnamed. Match only if the file import is also
// unnamed.

// If the patch import is not a metavar, then we meant
// to match it verbatim, so this is not a match.

// If the patch import is a metavar, then record a fake name
// for the metavar so that there's a value associated with
// "foo" in "foo.X", but also record that the real name for
// this import metavar is empty.

// Both are named. Match as-is and also associate the package
// name with the import path so that we can delete it later.

// Both imports are named. Match the name and move on.

type importMetavarKey string

type importMetavarData struct{ Unnamed bool }

type importKey string // import path

type importData struct {
	Name string // package name of the import

	// Name of the import metavar we used to match this value, if any.
	MetavarKey importMetavarKey
}

// ImportsMatcher matches multiple imports in a Go file.
type ImportsMatcher struct {
	Imports []ImportMatcher
}

func (c *matcherCompiler) compileImports(imps []*ast.ImportSpec) ImportsMatcher {
	_ = "STUB: not implemented"
	return *new(ImportsMatcher)
}

// Match matches a block of imports in a file.
func (m ImportsMatcher) Match(file *ast.File, d data.Data) (_ data.Data, ok bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}

type _importsKey string

var importsKey _importsKey

type importsData struct {
	MatchedImports []string // import paths
}

// ImportReplacer replaces imports in a file.
type ImportReplacer struct {
	Name          Replacer // named import, if any
	NameS         string
	NameIsMetavar bool

	Path string // import path as a string
	Fset *token.FileSet
}

func (c *replacerCompiler) compileImport(imp *ast.ImportSpec) ImportReplacer {
	_ = "STUB: not implemented"
	return *new(ImportReplacer)
}

// TODO: Same as matcher, maybe we should attempt to determine the
// package name.

// Replace adds a single import. Returns the name of the import that was
// added.
func (r ImportReplacer) Replace(d data.Data, cl Changelog, f *ast.File) (string, error) {
	_ = "STUB: not implemented"
	// name is the name we want to use for the named import, and pkgName is
	// how the rest of the file references this import.
	return "", nil
}

// The name replacer will produce the value for the named
// import specified in the patch as-is, or if it was a
// metavariable, using the matched value.
//
// This is undesirable for the case where the named import
// matched an unnamed import. For that case, we want to ignore
// the recorded name. So, if the named import is a
// metavariable that matched an unnamed import, don't look up
// its recorded value.

// default to metavar name

// pos is irrelevant

// TODO: more sophisticated package name guessing logic here
// and below.

// ImportsReplacer replaces a block of imports.
type ImportsReplacer struct {
	Imports []ImportReplacer
	Fset    *token.FileSet
}

func (c *replacerCompiler) compileImports(imps []*ast.ImportSpec) ImportsReplacer {
	_ = "STUB: not implemented"
	return *new(ImportsReplacer)
}

// Replace adds zero or more imports t a file.
//
// Returns a list of the names of the imports that were added, if known.
func (r ImportsReplacer) Replace(d data.Data, cl Changelog, f *ast.File) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cleanup cleans up unused imports. newNames is a list of names of imports
// that were added by the plus sections.
func (r ImportsReplacer) Cleanup(d data.Data, f *ast.File, newNames []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete matched imports that are no longer used.

// If we used a metavariable to match this import,
// record if its name was actually empty.

// If this import was replaced by an added import, kill it.

// For each import decl, if the import is the last in the group,
// delete the parens around it.

// To make the AST happy, if f.Imports is empty, explicitly nil it.

// TODO: This is probably not the best place or method to implement this.
func usesNameAsTopLevel(f *ast.File, name string) bool { _ = "STUB: not implemented"; return false }

// keep looking

// keep looking
