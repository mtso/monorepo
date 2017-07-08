package main

import (
	"errors"
	// "io/ioutil"
	"regexp"
)

// ErrNoSummary is returned by GetSummary if the `[summary]::` pattern is not found.
var ErrNoSummary = errors.New("No summary tag found")

// SummaryPattern is used to match one line of text after the tag `[summary]::`.
var SummaryPattern = regexp.MustCompile("(?:\\n\\[summary\\]::[^\\n]*\\n)([^\\n]*)(?:\\n)")

// GetSummary finds the one-line summary from a byte-slice of
// a text file, the line must be preceded by the `[summary]::`
// tag and be on a new line after the tag.
func GetSummary(readme []byte) ([]byte, error) {
	match := SummaryPattern.FindSubmatch(readme)
	if len(match) > 1 {
		return match[1], nil
	}
	return nil, ErrNoSummary
}
