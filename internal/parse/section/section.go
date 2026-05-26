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

// Package section is responsible for splitting a program into its different
// sections without attempting to parse the contents.
package section

import (
	"go/ast"
	"go/token"
)

// Program is a single .patch file consisting of one or more changes.
type Program []*Change

// Change is a single change in a program.
type Change struct {
	// Position at which the first @ of the header occurs.
	HeaderPos token.Pos

	// Changes can optionally have a name.
	//
	// If any, it is specified between the first pair of @@s in the change.
	Name string

	// Metavariables section of the change.
	Meta Section

	// Position of the second "@@".
	AtPos token.Pos

	// Patch is the patch section of the change. This is the code after the
	// second @@.
	Patch Section

	// Comments in the patch
	Comments []string
}

var _ ast.Node = (*Change)(nil)

// Pos returns the position at which this change begins.
func (c *Change) Pos() token.Pos {
	_ = "STUB: not implemented"

	// End returns the position of the first character after this change.
	return *new(token.Pos)
}

func (c *Change) End() token.Pos { _ = "STUB: not implemented"; return *new(token.Pos) }

// An emty change is effectively a no-op but that's not relevant here.
// The End position for an empty change is when the second pair of "@@"s
// ends.

// Section is a section of the change.
type Section []*Line

// Line is a single line from the patch.
type Line struct {
	// Position at which this line begins.
	StartPos token.Pos

	// Contents of the line.
	Text []byte
}

var _ ast.Node = (*Line)(nil)

// Pos returns the position at which this line begins.
func (l *Line) Pos() token.Pos {
	_ = "STUB: not implemented"

	// End returns the position of the character just past this line.
	return *new(token.Pos)
}

func (l *Line) End() token.Pos { _ = "STUB: not implemented"; return *new(token.Pos) }

// Split splits a Program into sections.
func Split(fset *token.FileSet, filename string, content []byte) (Program, error) {
	_ = "STUB: not implemented"
	return *new(Program), nil
}

// read the first line

type programSplitter struct {
	file    *token.File // file to feed newline information
	content []byte      // raw source

	text         []byte    // contents of the current line
	pos          token.Pos // position at which text begins
	lastComments []string  // last comment stored before the section

	eof bool // whether we've reached EOF

	startOffset int // offset at the start of the current line
	offset      int // current position in content

	errors []error
}

// Posts an error message with positional information.
func (p *programSplitter) errf(off int, msg string, args ...any) { _ = "STUB: not implemented"; return }

// Skips to the end of line. This may be a newline character or EOF.
func (p *programSplitter) skipUntilEOL() { _ = "STUB: not implemented"; return }

// Advances the scanner to the next non-comment line and collects the last comments.
func (p *programSplitter) next() {
	_ = "STUB: not implemented"
	// setting last comments to empty string whenever encounter non comment line
	return
}

// Reached EOF.

// Comments are supported only on their own lines.
func isComment(s []byte) bool { _ = "STUB: not implemented"; return false }

func (p *programSplitter) readProgram() Program { _ = "STUB: not implemented"; return *new(Program) }

// Read and return a Change, or nil if EOF was reached.
func (p *programSplitter) readChange() *Change {
	_ = "STUB: not implemented"
	// Can't use a struct literal here because readName and readMeta advance
	// p.pos between HeaderPos and AtPos.
	return nil
}

// Reads the name of a change.
func (p *programSplitter) readName() string { _ = "STUB: not implemented"; return "" }

// unnamed

// named

// leading @
// trailing @

// Number of bytes shaved off the front of text. We'll use this to
// mark the position in the error message in case of an invalid name.
// leading @

// Manually trim the left so that we can keep track of the number of
// bytes we're shifting.

// Reads the metavariables section of the change.
func (p *programSplitter) readMeta() Section { _ = "STUB: not implemented"; return *new(Section) }

// Reads the patch section of a change, stopping when a new change is
// encountered or the end of the file is reached.
func (p *programSplitter) readPatch() Section {
	_ = "STUB: not implemented"
	// skip past "@@" marking the end of metavariables section
	return *new(Section)
}

// new change begins

// Validates that the given non-empty string is a valid Go identifier. If the
// name is invalid, the first invalid character and the index at which it
// occurs is returned.
func validateChangeName(s string) (i int, ch rune, ok bool) {
	_ = "STUB: not implemented"
	return 0,

		// Only letters and underscores are allowed.
		0, false
}

// ...unless this is past the first character, in which case numbers
// are allowed too.

func notIsSpace(ch rune) bool { _ = "STUB: not implemented"; return false }
