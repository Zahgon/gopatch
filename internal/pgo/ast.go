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

package pgo

import (
	"go/ast"
	"go/token"
)

// File is a single pgo file. This is analogous to go/ast's File.
type File struct {
	// Package name specified in the file, if any.
	Package string

	// Imports made in the source.
	Imports []*ast.ImportSpec

	// Comments declared in the file.
	Comments []*ast.CommentGroup

	// Top-level node declared in the file.
	//
	// pgo allows only one node at the top-level. It must be either a standard
	// Go top-level declaration, an expression, or a list of statements.
	Node Node
}

// Node unifies the custom node types we introduce to the Go AST.
type Node interface {
	ast.Node

	pgoNode()
}

var (
	_ Node = (*Dots)(nil)
	_ Node = (*Expr)(nil)
	_ Node = (*FuncDecl)(nil)
	_ Node = (*GenDecl)(nil)
	_ Node = (*StmtList)(nil)
)

// Expr is a Go expression at the top-level in pgo.
type Expr struct{ ast.Expr }

func (*Expr) pgoNode() {
	_ = "STUB: not implemented"

	// StmtList is a list of statements at the top-level of a pgo file.
	return
}

type StmtList struct {
	List []ast.Stmt // inv: len > 0
}

func (*StmtList) pgoNode() {
	_ = "STUB: not implemented"

	// Pos returns the start position of the statement list or NoPos if there are
	// no statements in it.
	return
}

func (l *StmtList) Pos() token.Pos {
	_ = "STUB: not implemented"
	return *

	// End returns the position of the character immediately after this statement
	// list, or NoPos if there are no statements in this list.
	new(token.Pos)
}

func (l *StmtList) End() token.Pos { _ = "STUB: not implemented"; return *new(token.Pos) }

// FuncDecl is a Go function declaration at the top-level in pgo.
type FuncDecl struct{ *ast.FuncDecl }

func (*FuncDecl) pgoNode() {
	_ = "STUB: not implemented"

	// GenDecl is a Go general declaration at the top-level in pgo.
	return
}

type GenDecl struct{ *ast.GenDecl }

func (*GenDecl) pgoNode() {
	_ = "STUB: not implemented"

	// Dots is a "..." used as an expression.
	//
	// If used as a statement, Dots will be inside an ExprStmt.
	return
}

type Dots struct {
	ast.Expr

	Dots token.Pos // position of dots
}

func (*Dots) pgoNode() {
	_ = "STUB: not implemented"

	// Pos returns the start position of "...".
	return
}

func (d *Dots) Pos() token.Pos {
	_ = "STUB: not implemented"

	// End returns the position after "...".
	return *new(token.Pos)
}

func (d *Dots) End() token.Pos { _ = "STUB: not implemented"; return *new(token.Pos) }
