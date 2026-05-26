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

// TODO(abg): This has a fair amount of logic similar to stmt_list's
// reproduction logic. We can probably generalize this.

// ForDotsMatcher represents a "for ..." stmt, matching both, for and range
// statements.
type ForDotsMatcher struct {
	Dots token.Pos
	Body Matcher
}

func (c *matcherCompiler) compileForStmt(v reflect.Value) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

// Not a "for ...". Fall back to the usual logic.

// Match matches either a for statement or a range statement.
func (m ForDotsMatcher) Match(got reflect.Value, d data.Data, r Region) (data.Data, bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}

// The BlockStmt is named Body on both types.

// r tracks the unchanged region. In this case, it
// ends when the body starts.

// ForDotsReplacer replaces a "for ...".
type ForDotsReplacer struct {
	Dots token.Pos
	Body Replacer

	dotAssoc map[token.Pos]token.Pos
}

func (c *replacerCompiler) compileForStmt(v reflect.Value) Replacer {
	_ = "STUB: not implemented"
	return *new(Replacer)
}

// Not a "for ...". Fall back to the usual logic.

// Replace rebuilds a For or Range statement from the originally captured
// fields.
func (r ForDotsReplacer) Replace(d data.Data, cl Changelog, pos token.Pos) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// Reproduce fields besides the body as-is.

type forDotsKey token.Pos

type forDotsField struct {
	Idx   int           // index of the field in Type
	Value reflect.Value // captured value of the field
}

type forDotsData struct {
	// Type of statement that we matched.
	//
	// This is one of ForStmt or RangeStmt.
	Type reflect.Type

	// Field index in Type at which the "Body *ast.BlockStmt" field is
	// present.
	BodyFieldIdx int

	OtherFields []forDotsField

	Region Region
}
