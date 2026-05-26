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

package main

import (
	"go/ast"
	"go/token"
	"io"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/uber-go/gopatch/internal/engine"
)

func main() {
	os.Exit(runMain())
}

type arguments struct {
	Patterns []string `positional-arg-name:"pattern"`
}

type options struct {
	Patches              []string  `short:"p" long:"patch" value-name:"file"`
	PatchesFile          string    `short:"P" long:"patches-file" value-name:"file"`
	Diff                 bool      `short:"d" long:"diff"`
	DisplayVersion       bool      `long:"version"`
	Print                bool      `long:"print-only"`
	SkipImportProcessing bool      `long:"skip-import-processing"`
	SkipGenerated        bool      `long:"skip-generated"`
	Args                 arguments `positional-args:"yes"`
	Verbose              bool      `short:"v" long:"verbose"`
}

func newArgParser() (*flags.Parser, *options) { _ = "STUB: not implemented"; return nil, nil }

// The following is more readable than long descriptions in struct
// tags.

// loadPatches loads patches specified by command line options.
func loadPatches(fset *token.FileSet, opts *options, stdin io.Reader) ([]*engine.Program, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If -p and -P are unset, read from stdin.

// sourcePath is the path to a Go source file.
type sourcePath struct {
	// Form closest to what was provided by the user.
	// If they provided a relative path, this will be relative.
	Provided string

	// Absolute path to the file.
	Absolute string
}

func findGoFiles(cwd, path string) (_ []sourcePath, err error) {
	_ = "STUB: not implemented"
	// Users may expect "./..."-stlye patterns to work.
	return nil, nil
}

// empty if path was absolute

// drop extraneous ., .., etc.

func findFiles(cwd string, patterns []string) (_ []sourcePath, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type mainCmd struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	Getwd func() (string, error) // == os.Getwd
}

func runMain() (exitCode int) { _ = "STUB: not implemented"; return 0 }

func (cmd *mainCmd) Run(args []string) error { _ = "STUB: not implemented"; return nil }

/* src */

// If at least one patch didn't match, there's nothing to do.
// If --print-only was passed, print the contents out as-is.

// This error shouldn't occur due to checks in
// findFiles, loadPatches and format.Node()

func checkGeneratedCode(f *ast.File) bool { _ = "STUB: not implemented"; return false }

func (cmd *mainCmd) preview(
	filename string,
	originalContent, modifiedContent []byte,
	comments []string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *mainCmd) printComments(filename string, comments []string) {
	_ = "STUB: not implemented"
	return
}

type patchRunner struct {
	fset    *token.FileSet
	patches []*engine.Program
	errors  []error
}

func newPatchRunner(fset *token.FileSet, patches []*engine.Program) *patchRunner {
	_ = "STUB: not implemented"
	return nil
}

func (r *patchRunner) Apply(filename string, f *ast.File) (fout *ast.File, comments []string, matched bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// This patch didn't modify the file. Try the next one.

func cleanupFilePos(tfile *token.File, cl engine.Changelog, comments []*ast.CommentGroup) {
	_ = "STUB: not implemented"
	return
}

// Remove comments in the changed sections of the code.
