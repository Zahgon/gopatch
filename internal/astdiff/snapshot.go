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

package astdiff

import (
	"go/ast"
	"go/token"
	"reflect"
)

type value struct {
	t reflect.Type

	// Only one of the following three is set.
	isNil    bool
	value    any
	Elem     *value
	Children []*value

	// Set only if this is a Node.
	IsNode   bool
	pos, end token.Pos
	Comments []*ast.CommentGroup
}

func (v *value) Type() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
func (v *value) Kind() reflect.Kind { _ = "STUB: not implemented"; return *new(reflect.Kind) }
func (v *value) Pos() token.Pos     { _ = "STUB: not implemented"; return *new(token.Pos) }
func (v *value) End() token.Pos     { _ = "STUB: not implemented"; return *new(token.Pos) }
func (v *value) Interface() any     { _ = "STUB: not implemented"; return *new(any) }
func (v *value) Len() int           { _ = "STUB: not implemented"; return 0 }
func (v *value) IsNil() bool        { _ = "STUB: not implemented"; return false }

func snapshot(v reflect.Value, cmap ast.CommentMap) (val *value) {
	_ = "STUB: not implemented"
	return nil
}

// Ident.obj forms a cycle.

// Snapshots don't care about comment groups.

func minPos(l, r token.Pos) token.Pos { _ = "STUB: not implemented"; return *new(token.Pos) }

func maxPos(l, r token.Pos) token.Pos { _ = "STUB: not implemented"; return *new(token.Pos) }
