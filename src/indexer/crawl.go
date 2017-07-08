package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
)

var readmePattern = regexp.MustCompile("(?i)README")

type Content struct {
	Dirname string
	Summary string
}

func main() {
	content := make([]Content, 0)

	crawlDir(".", func(dirname, filename string) {
		// Read contents if it is a README file.
		if readmePattern.MatchString(filename) {
			c, err := readContent(dirname, filename)
			if err == ErrNoSummary {
				return
			}
			content = append(content, c)
		}
	})

	PrintContents(content)
}

func PrintContents(table []Content) {
	for _, content := range table {
		fmt.Printf("%s/ | %s\n", content.Dirname, content.Summary)
	}
}

func readContent(dirname, filename string) (c Content, err error) {
	readme, err := ioutil.ReadFile(path.Join(dirname, filename))
	if err != nil {
		return c, err
	}

	summary, err := GetSummary(readme)
	if err != nil {
		return c, err
	}

	return Content{dirname, string(summary)}, nil
}

// Recursively reads filepaths from a given path and invokes
// a callback func(dirname, filename string)
func crawlDir(dirname string, callback func(string, string)) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		name := file.Name()

		if file.IsDir() {
			dirpath := path.Join(dirname, name)
			crawlDir(dirpath, callback)
			continue
		}

		callback(dirname, name)
	}
	return
}
