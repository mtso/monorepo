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
		if !readmePattern.MatchString(filename) {
			return
		}

		// Read contents if it is a README file.
		readme, err := ioutil.ReadFile(path.Join(dirname, filename))
		if err != nil {
			return
		}

		summary, err := GetSummary(readme)
		if err == ErrNoSummary {
			return
		}

		content = append(content, Content{dirname, string(summary)})
	})

	print(content)
}

func print(table []Content) {
	for _, content := range table {
		fmt.Printf("`%s/` | %s\n", content.Dirname, content.Summary)
	}
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
