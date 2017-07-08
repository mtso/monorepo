package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
)

var summaryPattern = regexp.MustCompile("(?:\\n\\[summary\\]::[^\\n]*\\n)([^\\n]*)(?:\\n)")
var ErrNoSummary = errors.New("No summary tag found")

func GetSummary(readme []byte) ([]byte, error) {
	match := summaryPattern.FindSubmatch(readme)
	if len(match) > 1 {
		return match[1], nil
	}
	return nil, ErrNoSummary
}

func main() {
	rm, err := ioutil.ReadFile("./README.md")
	if err != nil {
		panic(err)
	}

	sum, err := GetSummary(rm)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", sum)

}
