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

package augment

import (
	"go/scanner"
	"go/token"
)

// find looks for augmentations to the Go source inside the given patch
// source.
func find(src []byte) ([]Augmentation, error) { _ = "STUB: not implemented"; return nil, nil }

/* flags */

// read first token

type finder struct {
	file    *token.File
	scanner *scanner.Scanner

	tok token.Token // current token
	pos token.Pos   // position of current token

	// Offset of tok inside the original source file. This is equal to
	// file.Offset(pos).
	offset int

	// Augmentations and errors recorded so far.
	augs   []Augmentation
	errors []error
}

// Called by go/scanner in case of errors.
func (f *finder) onError(pos token.Position, msg string) { _ = "STUB: not implemented"; return }

func (f *finder) append(aug Augmentation) { _ = "STUB: not implemented"; return }

// Advances the scanner.
func (f *finder) next() { _ = "STUB: not implemented"; return }

// Returns the line number for the provided token.Pos.
func (f *finder) line(pos token.Pos) int { _ = "STUB: not implemented"; return 0 }

func (f *finder) find() []Augmentation { _ = "STUB: not implemented"; return nil }

func (f *finder) process() { _ = "STUB: not implemented"; return }

// Ensures that we have a package clause, recording the need for one if not.
func (f *finder) pkg() { _ = "STUB: not implemented"; return }

// Missing a package clause. Generate a fake one.

// package
// package_name
// ;

// Skips over the imports, if any.
func (f *finder) imports() { _ = "STUB: not implemented"; return }

// import

// import group (import (...))
// Skip until )

// )
// ;

// dot import (import . "foo")
// .

// named import (import x "foo")
// x

// invalid syntax; parser will handle this

// "foo"
// ;

// Ensures that we have a valid top-level decl.
func (f *finder) topLevelDecl() { _ = "STUB: not implemented"; return }

// We can parse these as GenDecls. These transformations will
// apply to both, top-level GenDecls and GenDecls nested
// inside DeclStmts.
//
// If users want to place multiple of these in the patch, they
// should use {} to ensure that the patch is interpreted as a
// list of statemetns.
// type/const/var

// Add a fake func()

// {

// Add a fake func() { ... }

func (f *finder) ident() {
	_ = "STUB: not implemented"
	// IDENT

	// foo...
	return
}

// leave unchanged

func (f *finder) ellipsis() { _ = "STUB: not implemented"; return }

// ...

// The scanner tracks some parsing-related state, implicitly inserting
// SEMICOLON tokens when newlines are encountered at the end of a
// statement.
//
// This is problematic because ELLIPSIS is not a valid end of a statement.
// So the following pgo code,
//
//   ...
//   foo
//
// Will be scanned as, [ELLIPSIS, IDENT] rather than [ELLIPSIS, SEMICOLON,
// IDENT], which makes it impossible to differentiate between that and
// just "...foo" based on just tokens alone.
//
// To work around this, we need to check if the ELLIPSIS and IDENT are on
// the same line.

// ...foo

// leave unchanged

// ...

// Processes a top-level function or method declaration.
func (f *finder) funcDecl() {
	_ = "STUB: not implemented"
	// func

	// handle receiver if present
	return
}

// (

// )

// func name

// Processes a function literal.
func (f *finder) function() {
	_ = "STUB: not implemented"
	// func
	return
}

func (f *finder) params() { _ = "STUB: not implemented"; return }

func (f *finder) results() { _ = "STUB: not implemented"; return }

// return to process loop for unwrapped results

// Processes a argument or results lists.
func (f *finder) fieldList() {
	_ = "STUB: not implemented"
	// (
	return
}

// list of offsets at which ellipses were found

// ident

// ident was beginning of selector expression; pop off
// the "." and next ident before checking if named.
// .
// ident

// For the next token,
//
// - comma indicates that a new parameter is beginning
// - rparen indicates the end of the parameter list
//
// If either token appears after a single identifier, this was an unnamed
// parameter. So anything else as the next token indicates a named
// parameter.

// ...

// ellipsis was variadic operator, continue without
// augmenting this ellipsis.

// * or , e.g.

// )
