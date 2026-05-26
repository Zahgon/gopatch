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
	"github.com/uber-go/gopatch/internal/pgo"
)

// stmtSliceContainerMatcher matches AST nodes that contain statement slices
// ([]Stmt) anywhere in a Go AST.
//
// Specifically, it matches BlockStmt, CaseClause, and CommClause nodes.
type stmtSliceContainerMatcher struct {
	Stmts Matcher // matcher for a []Stmt
}

// Compiles a Matcher from a pgo.StmtList. When a list of statements is
// provided at the top level in the minus section of the patch, we should
// match anywhere in the AST where a []ast.Stmt can be present. We'll use
// stmtSliceContainerMatcher for this.
func (c *matcherCompiler) compilePGoStmtList(slist *pgo.StmtList) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

func (m stmtSliceContainerMatcher) Match(v reflect.Value, d data.Data, r Region) (data.Data, bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}

// Instead of copying individual fields of BlockStmt, CaseClause, and
// CommClause, we will match against the statements (present under
// .List in BlockStmt and .Body under CaseClause and CommClause) and
// make a shallow copy of all other attributes of the object, to be
// replicated in the Replacer.

// Position of the end of the text right before statements
// start. For block statements, this will be the position of
// "{", for case clauses, it will be the position of ":".

// Fields besides the one containing []Stmt.

// Information about the field containing []Stmt.

// stmtSliceContainerReplacer reproduces an AST node for which a statement
// list was previously matched.
//
// For example, if we previously matched a CaseClause, this will reproduce the
// original CaseClause but with its body replaced with the output of the Stmts
// replacer.
type stmtSliceContainerReplacer struct {
	Stmts Replacer // replacer for []Stmt
}

// Compiles a Replacer from a pgo.StmtList. When a list of statements is
// provided at the top level in the plus section of teh patch, we should be
// able to reproduce the original container for these statements (BlockStmt,
// CaseClause, CommClause) as-is with only the statement list modified.
func (c *replacerCompiler) compilePGoStmtList(slist *pgo.StmtList) Replacer {
	_ = "STUB: not implemented"
	return *new(Replacer)
}

func (r stmtSliceContainerReplacer) Replace(d data.Data, cl Changelog, pos token.Pos) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// Reproduce the original struct without setting Stmts.

type _stmtListKey struct{}

var stmtListKey _stmtListKey

type stmtListData struct {
	// Type of statement that we matched.
	//
	// This is one of BlockStmt, CaseClause, or CommClause.
	Type reflect.Type

	// Index in Type at which the []ast.Stmt can be found.
	//
	// That is, Type.Field(StmtFieldIdx) should be []ast.Stmt.
	StmtFieldIdx int

	// Indexes and values of the other fields in the type.
	OtherFields []stmtListField

	// Region of the original statement (block, case, or comm) that was
	// unmodified.
	//
	// For block, it's up to the "{", for case and comm, it's up to the
	// ":" after the label.
	UnchangedRegion Region
}

type stmtListField struct {
	// Index of the field in Type.
	FieldIdx int

	// Captured value of the field.
	Value reflect.Value
}

func dotsStmt(pos token.Pos) ast.Stmt { _ = "STUB: not implemented"; return *new(ast.Stmt) }
