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

// SliceDotsMatcher implements support for "..." in portions of the AST where
// slices of values are expected.
type SliceDotsMatcher struct {
	// TODO: Type

	// List of contiguous sections to match against.
	Sections [][]Matcher // inv: len > 0

	// Positions at which dots were found.
	Dots []token.Pos // inv: len(dots) = len(sections) - 1
}

func (c *matcherCompiler) compileSliceDots(items reflect.Value, isDots func(ast.Node) bool) Matcher {
	_ = "STUB: not implemented"
	// TODO(abg): Validate that this is called with ast.Node-compatible types
	// only.
	return *new(Matcher)
}

// Optimization: If there are no "..."s, we can use the faster slice
// matcher.

// Match matches
func (m SliceDotsMatcher) Match(got reflect.Value, d data.Data, r Region) (data.Data, bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}

// The first section must match in-place.

// Returns Region for items[start:end].
func sectionRegion(items []reflect.Value, r Region, start, end int) Region {
	_ = "STUB: not implemented"

	// If this isn't the first item, use the end position of the previous
	// sibling as the start of the region.
	return *new(Region)
}

// If this isn't the last item, use the start position of the next
// sibling as the end of this region.

// matchPrefix matches all items in want starting at got[idx]. If all items
// were matched, the new index for the remaining matches is returned.
func matchPrefix(want []Matcher, got []reflect.Value, d data.Data, r Region, idx int) (newIdx int, _ data.Data, ok bool) {
	_ = "STUB: not implemented"
	return 0, *new(data.Data), false
}

// findSection attempts to match want starting at got[idx], moving onto idx+1,
// idx+2, and so on until a match is found. Returns the new index for the
// remaining matches.
//
// Invariant: If ok is true, a list of skipped items will have been pushed to
// Data.
func findSection(dots token.Pos, want []Matcher, got []reflect.Value, d data.Data, r Region, idx int) (newIdx int, _ data.Data, ok bool) {
	_ = "STUB: not implemented"
	// Special case: Looking for "..." at the end of the list. Skip everything
	// in got.
	return 0, *new(data.Data), false
}

// SliceDotsReplacer replaces target nodes and reproduces the values captured by
// "..." in places which the AST expects a slice.
type SliceDotsReplacer struct {
	Type     reflect.Type
	Sections [][]Replacer // inv: len > 0

	// Positions at which dots were found.
	Dots []token.Pos // inv: len(dots) = len(sections) - 1

	dotAssoc map[token.Pos]token.Pos
}

func (c *replacerCompiler) compileSliceDots(items reflect.Value, isDots func(ast.Node) bool) Replacer {
	_ = "STUB: not implemented"
	return *new(Replacer)
}

// Optimization: If there are no "..."s, we can use the faster slice
// replacer.

// TODO: we probably just need some sort of "what's my associated
// dot pos/ID" functionality injected here instead of
// grabbing the map from compiler.

// Replace replaces target Nodes in slices where elements may have been elided
// in the patch.
func (r SliceDotsReplacer) Replace(d data.Data, cl Changelog, pos token.Pos) (reflect.Value, error) {
	_ = "STUB: not implemented"
	// TODO: Recurse into Cursor
	return *new(reflect.Value), nil
}

// TODO: Need to explicitly handle nil vs empty

type sliceDotsKey token.Pos

type sliceDotsData struct {
	Skipped []reflect.Value
	Region  Region
}

// Pushes the list of skipped items to Data. Skipped items are got[start:end].
func pushSliceDotsSkipped(d data.Data, dots token.Pos, got []reflect.Value, r Region) data.Data {
	_ = "STUB: not implemented"
	return *new(data.Data)
}

func lookupSliceDotsSkipped(d data.Data, dots token.Pos) (result []reflect.Value, region Region) {
	_ = "STUB: not implemented"
	return nil,

		// TODO(abg): If ok is false, we couldn't find data for this "...". That's
		// invalid and we should report that.
		*new(Region)
}
