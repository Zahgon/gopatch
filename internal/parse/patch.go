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

package parse

import (
	"github.com/uber-go/gopatch/internal/parse/section"
	"github.com/uber-go/gopatch/internal/pgo"
)

// parsePatch parses a Patch from the given source.
func (p *parser) parsePatch(i int, c *section.Change) (*Patch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME: Hack: If one of minus and plus believes their side is an
// statement and the other believes it's an expression, make them both
// expresisons.

// parses one version of the unified diff of a file.
func (p *parser) parsePatchVersion(name string, f patchVersion) (*pgo.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: fill position data for lines before msrc.Node.

// patchVersion is one of the versions of a patch specified in a unified diff.
type patchVersion struct {
	// Contents of the file.
	Contents []byte

	// Positional information for each line in Contents.
	//
	// Each LinePos contains matches an offset in Contents to a token.Pos in
	// the original patch file.
	Lines []section.LinePos
}

// splitPatch splits a patch into the before and after versions of the
// file.
//
// Given the unified diff,
//
//	 foo
//	-bar
//	+baz
//	 qux
//
// This functions splits it into,
//
//	Before  After
//	------  -----
//	foo     foo
//	bar     baz
//	qux     qux
func splitPatch(patch section.Section) (before, after patchVersion) {
	_ = "STUB: not implemented"
	return *new(patchVersion), *new(patchVersion)
}

// If true, the corresponding item won't get a LinePos entry because
// it wasn't written to this time.

// '-'

// '+'
