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

package data

// Data stores arbitrary data for Matchers and Replacers as they traverse or
// build the Go AST.
//
// Data provides an interface very similar to context.Context.
type Data interface {
	// Keys returns all known keys for this Data object in an unspecified
	// order.
	Keys() []any

	// Values retrives the value associated with the given key or nil.
	Value(k any) any
}

// New builds a new empty Data.
func New() Data { _ = "STUB: not implemented"; return *new(Data) }

type emptyData struct{}

func (d *emptyData) Keys() []any { _ = "STUB: not implemented"; return nil }
func (d *emptyData) Value(any) any {
	_ = "STUB: not implemented"

	// WithValue returns a new Data with the given key-value pair associated with
	// it.
	//
	//	d = data.WithValue(d, x, 42)
	//	d.Value(x) // == 42
	//
	// The original object is left unmodified so if the returned Data object is
	// discarded, its value will not be made available.
	//
	// Retreiving the value from this Data object is a linear time operation. To
	// optimize for read-heavy use cases, use the Index function.
	//
	// Panics if either the key or the value are nil.
	return *new(any)
}

func WithValue(d Data, k, v any) Data { _ = "STUB: not implemented"; return *new(Data) }

type valueData struct {
	Data

	k, v any
}

func (d *valueData) Keys() []any { _ = "STUB: not implemented"; return nil }

func (d *valueData) Value(k any) any { _ = "STUB: not implemented"; return *new(any) }

// Lookup retrieves the value associated with the given key and stores it into
// the pointer that vptr points to.
//
// Reports whether a value was found.
//
//	d = data.WithValue(d, x, 42)
//
//	var out int
//	data.Lookup(k, x, &out) // == true
//
// Panics if the type of the value for the pointer is not compatible with the
// value associated with the key.
func Lookup(d Data, k, vptr any) (ok bool) { _ = "STUB: not implemented"; return false }

// Index returns a copy of the provided Data object where items are indexed
// for fast lookup.
//
// Lookups with Value or the Lookup function are constant time in the returned
// data object.
//
// Use this if your workload is divided between write-heavy and read-heavy
// portions.
func Index(d Data) Data {
	_ = "STUB: not implemented"
	// Already indexed.
	return *new(Data)
}

type indexedData struct {
	keys  []any
	items map[any]any
}

func (d *indexedData) Keys() []any { _ = "STUB: not implemented"; return nil }

func (d *indexedData) Value(k any) any { _ = "STUB: not implemented"; return *new(any) }
