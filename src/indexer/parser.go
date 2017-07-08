package main

import (
    "errors"
    "regexp"
    "fmt"
    "io/ioutil"
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

	match := summaryPattern.FindSubmatch(rm)

	for i, v := range match {
		fmt.Printf("%d: %s\n", i, v)
	}

	// idx := summaryPattern.FindSubmatchIndex(rm)
	// fmt.Printf("%s", rm)
	// fmt.Printf("%s\n", rm[idx[2]:idx[3]])
	// summaryPattern.FindAllSubmatch(rm, -1)[0][1]
	// fmt.Printf("%s\n", )
}
