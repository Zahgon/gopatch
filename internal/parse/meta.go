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

package parse

import (
	"go/ast"
	"go/scanner"
	"go/token"

	"github.com/uber-go/gopatch/internal/parse/section"
)

// Parses the metavariables section of the change at index i.
func (p *parser) parseMeta(i int, c *section.Change) (*Meta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We will create a new File with the contents of the metavariables
// section and map positions in it back to the original file for error
// messages.

// Generate a fake name for the File.

/* mode */

// read the first token

type metaParser struct {
	scanner *scanner.Scanner

	fset *token.FileSet
	pos  token.Pos   // current token position
	tok  token.Token // current token
	text string      // current token contents

	failed bool
	errors []error
}

// This function is called by go/scanner when errors are encountered. We
// connect it in the Init call above.
func (p *metaParser) onError(pos token.Position, msg string) { _ = "STUB: not implemented"; return }

// Posts a formatted error message to the parser.
func (p *metaParser) errf(msg string, args ...any) { _ = "STUB: not implemented"; return }

// Advances to the next token.
func (p *metaParser) next() { _ = "STUB: not implemented"; return }

// Parses the metavariables section.
func (p *metaParser) parse() *Meta { _ = "STUB: not implemented"; return nil }

// Parses and returns a VarDecl.
//
//	var x, y, z Foo
func (p *metaParser) parseDecl() *VarDecl { _ = "STUB: not implemented"; return nil }

// skip var/,

// A type name is expected after list of variables.

// go/scanner implicitly inserts SEMICOLON when a newline is found where a
// semicolon would be accepted. So we expect a semicolon after every var
// declaration.

// Reads and returns an identifier, advancing the parser to the next token.
// Fails the parser and returns nil if an identifier was not found.
func (p *metaParser) parseIdent() *ast.Ident { _ = "STUB: not implemented"; return nil }
