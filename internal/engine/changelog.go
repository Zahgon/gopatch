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

	"github.com/google/go-intervals/intervalset"
)

// Changelog records the ranges of positions changed over the course of
// successive patches.
type Changelog struct {
	// we maintain separate plus and minus sets because we want to be able to
	// submit changed/unchanged requests out of order.
	plus  *intervalset.Set
	minus *intervalset.Set
}

// NewChangelog builds a new, empty Changelog
func NewChangelog() Changelog { _ = "STUB: not implemented"; return *new(Changelog) }

// Changed records the portion of code altered in this match to the Changelog.
func (c Changelog) Changed(start, end token.Pos) { _ = "STUB: not implemented"; return }

// Unchanged records the portion of unaltered code to the Changelog.
func (c Changelog) Unchanged(start, end token.Pos) { _ = "STUB: not implemented"; return }

// Interval represents a consecutive set of positions in the source file.
type Interval struct {
	Start, End token.Pos
}

// ChangedIntervals returns the sum of the Intervals recorded to the Changelog.
func (c Changelog) ChangedIntervals() []Interval { _ = "STUB: not implemented"; return nil }

type span struct{ Start, End token.Pos }

func minPos(l, r token.Pos) token.Pos { _ = "STUB: not implemented"; return *new(token.Pos) }

func maxPos(l, r token.Pos) token.Pos { _ = "STUB: not implemented"; return *new(token.Pos) }

func zeroSpan() *span { _ = "STUB: not implemented"; return nil }

func validOrZero(s *span) *span { _ = "STUB: not implemented"; return nil }

var _ intervalset.Interval = (*span)(nil)

func (s *span) String() string { _ = "STUB: not implemented"; return "" }

func (s *span) Intersect(other intervalset.Interval) intervalset.Interval {
	_ = "STUB: not implemented"
	return *new(intervalset.Interval)
}

func (s *span) Before(i intervalset.Interval) bool {
	_ = "STUB: not implemented"

	// XXX: intervalset passes in nil when an interval is out of range (I
	// think?). There isn't a lot of visibility; it runs stuff in separate
	// goroutines for some reason.
	return false
}

func (s *span) IsZero() bool { _ = "STUB: not implemented"; return false }

func (s *span) Bisect(i intervalset.Interval) (intervalset.Interval, intervalset.Interval) {
	_ = "STUB: not implemented"
	return *new(intervalset.Interval), *new(intervalset.Interval)
}

func (s *span) Adjoin(other intervalset.Interval) intervalset.Interval {
	_ = "STUB: not implemented"
	return *new(intervalset.Interval)
}

func (s *span) Encompass(other intervalset.Interval) intervalset.Interval {
	_ = "STUB: not implemented"
	return *new(intervalset.Interval)
}

func (s *span) AsSet() *intervalset.ImmutableSet { _ = "STUB: not implemented"; return nil }
