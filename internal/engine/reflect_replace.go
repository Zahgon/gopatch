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

// compileGeneric compiles a Replacer for arbitrary values inside a Go AST.
func (c *replacerCompiler) compileGeneric(v reflect.Value) (r Replacer) {
	_ = "STUB: not implemented"
	return *new(Replacer)
}

// PtrReplacer replaces a pointer type.
type PtrReplacer struct {
	Replacer // underlying replacer

	Type reflect.Type // type of the pointer
}

func (c *replacerCompiler) compilePtr(v reflect.Value) Replacer {
	_ = "STUB: not implemented"
	return *new(Replacer)
}

// Replace replaces a pointer type.
func (r PtrReplacer) Replace(d data.Data, cl Changelog, pos token.Pos) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// SliceReplacer replaces a slice of values.
//
// This replacer does not support generating values elided with "...".
type SliceReplacer struct {
	Type reflect.Type // type of slice

	// Replacers for individual items in the slice.
	Items []Replacer
}

func (c *replacerCompiler) compileSlice(v reflect.Value) Replacer {
	_ = "STUB: not implemented"
	return *new(Replacer)
}

// Replace replaces a slice.
func (r SliceReplacer) Replace(d data.Data, cl Changelog, pos token.Pos) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// StructReplacer replaces a struct.
type StructReplacer struct {
	Type reflect.Type // type of the struct

	// Replacers for individual fields of the struct.
	Fields []Replacer
}

func (c *replacerCompiler) compileStruct(v reflect.Value) Replacer {
	_ = "STUB: not implemented"
	return *new(Replacer)
}

// Replace replaces a struct value.
func (r StructReplacer) Replace(d data.Data, cl Changelog, pos token.Pos) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// InterfaceReplacer replaces an interface value.
type InterfaceReplacer struct {
	Replacer // underlying replacer

	Type reflect.Type // type of the interface
}

func (c *replacerCompiler) compileInterface(v reflect.Value) Replacer {
	_ = "STUB: not implemented"
	return *new(Replacer)
}

// Replace replaces an interface value.
func (r InterfaceReplacer) Replace(d data.Data, cl Changelog, pos token.Pos) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// ValueReplacer replace a value as-is.
type ValueReplacer struct{ Value reflect.Value }

// Replace replaces a value as-is.
func (r ValueReplacer) Replace(data.Data, Changelog, token.Pos) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}
