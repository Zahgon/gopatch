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

// MetavarMatcher is compiled from a metavarible occurring in the minus
// section of the patch.
//
//	@@
//	var x expression
//	@@
//	-foo(x, x)
//
// The first time a metavariable occurs in the patch, it matches and captures
// a value of the requested type. For any consecutive appearances of the
// metavariable, the previously captured value is expected to match.
//
// For example, the patch above will match any expression for the first "x"
// and the second occurrence will require the previously captured expression
// to match.
type MetavarMatcher struct {
	Fset *token.FileSet

	// Name of the metavariable.
	Name string

	// Reports whether the provided type matches the metavariable declaration.
	TypeMatches func(reflect.Type) bool
}

func (c *matcherCompiler) compileIdent(v reflect.Value) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

// Not a metavariable. Match the identifer as-is.

// Match matches a metavariable from the patch in the AST.
func (m MetavarMatcher) Match(got reflect.Value, d data.Data, r Region) (data.Data, bool) {
	_ = "STUB: not implemented"
	return *new(data.Data), false
}

// We've already seen this metavariable. Match the value without
// altering captured data.

// We're seeing this for the first time. Capture it into a compiler and
// replacer so we can match and reproduce it later.

type metavarKey string

type metavarData struct {
	Matcher
	Replacer
}

func isExpression(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func isIdent(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

// MetavarReplacer is compiled from a metavarible occurring in the plus
// section of the patch.
//
//	@@
//	var x expression
//	@@
//	-foo(x)
//	+bar(x)
//
// A metavariable cannot be referenced in the plus secton of the patch if it
// wasn't in the minus section. For example, the following is invalid.
//
//	@@
//	var x expression
//	@@
//	-foo()
//	+foo(x)
//
// Each occurrence of the metavariable in the plus section of the patch is
// replaced with the value originally captured for it by the Matcher.
type MetavarReplacer struct {
	Name string
}

func (c *replacerCompiler) compileIdent(v reflect.Value) Replacer {
	_ = "STUB: not implemented"
	return *new(Replacer)
}

// Not a metavariable. Reproduce the identifier as-is.

// Replace reproduces the value of a matched metavariable.
func (m MetavarReplacer) Replace(d data.Data, cl Changelog, pos token.Pos) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// This will happen only if a metavariable was referenced in the plus
// section without being referenced in the minus section.
// TODO(abg): Guard against that during compilation instead.
