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

// SearchResult contains information about search results found by a
// SearchMatcher.
type SearchResult struct {
	// Object containing the matched node.
	parent ast.Node

	// Name of the field of the parent referring to the matched node.
	name string

	// Non-negative if parent.name is a slice. The match node is at this
	// index in parent.name.
	index int

	data   data.Data
	region Region
}

// SearchNode provides access to an AST node, its parent, and its positional
// information during a traversal.
type SearchNode interface {
	Node() ast.Node
	Parent() ast.Node
	Name() string
	Index() int
}

// Searcher inspects the given Node using the given Matcher and returns a
// non-nil SearchResult if it matched.
type Searcher func(SearchNode, Matcher, data.Data) *SearchResult

// SearchMatcher runs a Matcher on descendants of an AST, producing
// SearchResults into Data.
//
// The corresponding replacer applies a Replacer to these matched descendants.
type SearchMatcher struct {
	Search  Searcher
	Matcher Matcher
}

// Match runs the matcher on the provided ast.Node.
func (m SearchMatcher) Match(got reflect.Value, d data.Data, _ Region) (data.Data, bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}

// keep looking
/* post func */

// SearchReplacer replaces nodes found by a SearchMatcher.
type SearchReplacer struct {
	Replacer Replacer
}

// Replace replaces nodes found by a SearchMatcher.
func (r SearchReplacer) Replace(d data.Data, cl Changelog, pos token.Pos) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// This is a bug in our code.

// If the generated value isn't assignable to the target, the match
// was too eager. For example, trying to place "foo.Bar"
// (SelectorExpr) where only an identifier is allowed (in a variable
// declaration name, for example).

type _searchResultKey struct{}

var searchResultKey _searchResultKey

type searchResultData struct {
	Root    reflect.Value
	Results []*SearchResult
}

func pushSearchResults(d data.Data, root reflect.Value, results []*SearchResult) data.Data {
	_ = "STUB: not implemented"
	return *new(data.Data)
}

func lookupSearchResults(d data.Data) (root reflect.Value, results []*SearchResult) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// TODO(abg): Handle !ok
