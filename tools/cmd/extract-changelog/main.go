// extract-changelog extracts the release notes for a specific version from a
// file matching the format prescribed by https://keepachangelog.com/en/1.0.0/.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	cmd := mainCmd{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if err := cmd.Run(os.Args[1:]); err != nil && err != flag.ErrHelp {
		fmt.Fprintln(cmd.Stderr, err)
		os.Exit(1)
	}
}

type mainCmd struct {
	Stdout io.Writer
	Stderr io.Writer
}

const _usage = `USAGE

	%v [OPTIONS] VERSION

Retrieves the release notes for VERSION from a CHANGELOG.md file and prints
them to stdout.

EXAMPLES

  extract-changelog -i CHANGELOG.md v1.2.3
  extract-changelog 0.2.5

OPTIONS
`

func (cmd *mainCmd) Run(args []string) error { _ = "STUB: not implemented"; return nil }

func extract(r io.Reader, version string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Version headers take one of the following forms:
//
//   ## 0.1.3 - 2021-08-18
//   ## [0.1.3] - 2021-08-18

// Found a new version header. Stop extracting.

// unreachable but guard against it.

// always end with a single newline
