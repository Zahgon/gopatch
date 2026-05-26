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
	"reflect"

	"github.com/uber-go/gopatch/internal/data"
)

// GenericNodeMatcher is the top-level matcher for ast.Node objects.
type GenericNodeMatcher struct {
	Matcher // underlying matcher
}

// compileGeneric compiles a Matcher for arbitrary values inside a Go AST.
func (c *matcherCompiler) compileGeneric(v reflect.Value) (m Matcher) {
	_ = "STUB: not implemented"

	// Wrap with GenericNodeMatcher only if the type is a Go AST node.
	return *new(Matcher)
}

// Match matches an ast.Node.
func (m GenericNodeMatcher) Match(got reflect.Value, d data.Data, r Region) (data.Data, bool) {
	_ = "STUB: not implemented"
	// Collapse the region under consideration down to the region covered by
	// this node.
	return *new(data.Data), false
}

// PtrMatcher matches a non-nil pointer in the AST.
type PtrMatcher struct {
	Matcher // underlying matcher
}

func (c *matcherCompiler) compilePtr(v reflect.Value) Matcher {
	_ = "STUB: not implemented"
	// If the value is nil, we don't need to build the PtrMatcher.
	return *new(Matcher)
}

// Match matches a non-nil pointer.
func (m PtrMatcher) Match(got reflect.Value, d data.Data, r Region) (data.Data, bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}

// SliceMatcher matches a slice of values exactly.
//
// This matcher does not support eliding values with "...".
type SliceMatcher struct {
	// Matchers for individual fields of the slice.
	Items []Matcher
}

func (c *matcherCompiler) compileSlice(v reflect.Value) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

// Match mathces a slice of values.
func (m SliceMatcher) Match(got reflect.Value, d data.Data, r Region) (data.Data, bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}

// StructMatcher matches a struct.
type StructMatcher struct {
	Type reflect.Type // type of the struct

	// Matchers for individual fields of the struct.
	Fields []Matcher
}

func (c *matcherCompiler) compileStruct(v reflect.Value) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

// Match matches a struct.
func (m StructMatcher) Match(got reflect.Value, d data.Data, r Region) (data.Data, bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}

// InterfaceMatcher matches an interface value.
type InterfaceMatcher struct {
	Matcher // underlying matcher
}

func (c *matcherCompiler) compileInterface(v reflect.Value) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

// Match matches non-nil interface nalues.
func (m InterfaceMatcher) Match(got reflect.Value, d data.Data, r Region) (data.Data, bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}

// ValueMatcher matches a value as-is.
type ValueMatcher struct {
	Type  reflect.Type // underlying type
	Value any          // value to match
}

// Match matches a value as-is.
func (m ValueMatcher) Match(got reflect.Value, d data.Data, _ Region) (data.Data, bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}
