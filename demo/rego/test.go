package main

import (
	"fmt"

	"github.com/mtso/rego"
)

type Hello struct {
	rego.Component
}

func (h Hello) Render() rego.Component {
	return rego.CreateComponent(
		"div",
		rego.Props{
			"name": "world",
		},
	)

	name, ok := h.Props["name"]
	if ok {
		return "<div>hello, " + fmt.Sprintf("%s", name) + "</div>"
	}

	return "<div>hello, noname</div>"
}

type name string

func (n name) String() string {
	return string(n)
}

func main() {

	component := Hello{}
	component.Props = rego.Props{
		"name": "world",
	}

	markup := rego.Render(component)
	fmt.Println(markup)
}
