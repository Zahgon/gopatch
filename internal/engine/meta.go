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
	"github.com/uber-go/gopatch/internal/parse"
)

// MetavarType defines the different types of metavariables accepted in a
// '@@' section.
type MetavarType int

// Supported metavariable types.
const (
	ExprMetavarType  MetavarType = iota + 1 // expression
	IdentMetavarType                        // identifier
)

// Meta is the compiled representaton of a Meta section.
type Meta struct {
	// Variables defined in this Meta section and their types.
	Vars map[string]MetavarType
}

// LookupVar returns the type of the given metavariable or zero value if it
// wasn't found.
func (m *Meta) LookupVar(name string) MetavarType {
	_ = "STUB: not implemented"
	return *new(MetavarType)
}

func (c *compiler) compileMeta(m *parse.Meta) *Meta { _ = "STUB: not implemented"; return nil }

// Underscore isn't a variable declaration.
