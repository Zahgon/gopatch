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

// Package astdiff provides a means of diffing Go AST nodes that may be
// mutated in-place by taking snapshots of their state between successive
// patch executions.
//
//	snap := astdiff.Before(node)
//	modify(node)
//	snap = snap.Diff(node, changelog)
//	modify(node)
//	snap = snap.Diff(node, changelog)
package astdiff

import (
	"go/ast"
	"go/token"

	"github.com/uber-go/gopatch/internal/diff"
)

// Changelog records the ranges of positions changed over the course
// of successive patches.
type Changelog interface {
	// Changed informs the changelog that the code in the range [pos, end)
	// has been modified.
	Changed(pos, end token.Pos)
}

// Before starts an AST diff.
//
// Given the unchanged node, it creates a Snapshot of the initial state before
// applying any patches.
func Before(n ast.Node, comments ast.CommentMap) *Snapshot { _ = "STUB: not implemented"; return nil }

// Snapshot maintains the current view of the ast.Node being altered over the
// course of patching.
type Snapshot struct {
	value *value
}

// Diff compares the current Snapshot to the altered Node, updates the provided
// Changelog, and returns an updated Snapshot.
func (s *Snapshot) Diff(n ast.Node, cl Changelog) *Snapshot { _ = "STUB: not implemented"; return nil }

// Region denotes the consecutive positions
type Region struct{ Pos, End token.Pos }

type changeFinder struct {
	Region

	cl Changelog
}

func (f changeFinder) unchanged(from, to *value) { _ = "STUB: not implemented"; return }

func (f changeFinder) changed() { _ = "STUB: not implemented"; return }

func (f changeFinder) commentsFor(n *value) (before, after []*ast.Comment) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f changeFinder) Walk(from, to *value) (equal bool) { _ = "STUB: not implemented"; return false }

// Can't traverse into this subtree.

// Ident.obj forms a cycle.

// If LHS didn't exist, whether RHS exists or not doesn't affect the
// Changed sections.

// If LHS was deleted, this whole section has changed.

// Dereferencing a pointer or interface doesn't affect region.

func (f changeFinder) walkStruct(from, to *value) bool {
	_ = "STUB: not implemented"
	// The order of fields in AST structs matches how the elements appear in
	// the code so we can treat fields as siblings
	return false
}

// If the field is a Node, its range begins when the Node starts.

// If the field is a token.Pos, its range begins based on whatever
// its value is.

// Otherwise the start position is the end position of the last
// valid node.

// If the field is a Node, its range ends where the Node ends.

// The field that preceds this field should use this field's start
// position as its end position if it's not a Node.

func (f changeFinder) walkSlice(from, to *value) bool {
	_ = "STUB: not implemented"
	// If this isn't a slice of nodes, require that they're equal in length
	// and check all elements.
	return false
}

// Extend to sibling beinning and end.

// If the previous node has any trailing comments, maintain
// whitespace between them and us.

// If the next node has leading comments, maintain whitespace
// between them and us.

// If we have any comments associated with us, clamp to them.

// deleted

// new item, nothing to do

type nodeComparer struct{ diff.Result }

func compareNodes(from, to *value) diff.Result { _ = "STUB: not implemented"; return *new(diff.Result) }

func (c *nodeComparer) Walk(from, to *value) { _ = "STUB: not implemented"; return }

// not equal or similar

// Ident.obj forms a cycle.
