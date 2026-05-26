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

// PosMatcher matches token.Pos fields in an AST.
//
// token.Pos fields are not matched by value, but by validity. A token.Pos
// in the patch matches a token.Pos in the source file if they're both
// valid (non-zero) or they're both invalid (zero).
//
// PosMatcher records the matched position from the source file so that it
// can be reproduced by the PosReplacer later.
type PosMatcher struct {
	Fset *token.FileSet

	// Pos in the patch file.
	Pos token.Pos
}

func (c *matcherCompiler) compilePosMatcher(v reflect.Value) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

// Match matches a position in a file.
//
// If the position matches and is valid, this records this match in Data for
// later retrieval.
func (m PosMatcher) Match(v reflect.Value, d data.Data, _ Region) (data.Data, bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}

// PosReplacer replaces token.Pos fields.
//
// For valid positions, the replacer will use the previously recorded
// position for the field, if any. This ensures that whitespace between
// tokens from the original file are preserved as much as possible.
type PosReplacer struct {
	Fset *token.FileSet
	Pos  token.Pos
}

func (c *replacerCompiler) compilePosReplacer(v reflect.Value) Replacer {
	_ = "STUB: not implemented"
	return *new(Replacer)
}

// TODO DATA POSITION NEEDS TO BE MUTABLE. SEND POSITIONS UP FOR GENERATED
// CODE AS WE PROGRESS THROUGH MORE OF THE ORIGINAL POSITIONS

// Replace replaces position nodes.
func (r PosReplacer) Replace(d data.Data, cl Changelog, pos token.Pos) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// For positions, use the position associated with this item in the match
// data, falling back to the most recent position recorded in Data.

type posMatchKey struct {
	Line   int
	Column int
}

// Records a position match. Both positions MUST be valid.
func pushPosMatch(fset *token.FileSet, d data.Data, patchPos, matchedPos token.Pos) data.Data {
	_ = "STUB: not implemented"
	// TODO(abg): This can be cheaper if the positions in the PGo AST point
	// back to the patch file rather than the intermediate files.
	return *new(data.Data)
}

// Retrieves the position associated with the given patch position or an
// invalid position if none.
func lookupPosMatch(fset *token.FileSet, d data.Data, patchPos token.Pos) (pos token.Pos) {
	_ = "STUB: not implemented"
	return *new(token.Pos)
}

// TODO(abg): Handle !ok
