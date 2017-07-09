package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

const IndexHtml = `<!doctype html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>ELO</title>
	<link rel="stylesheet" href="/static/style.css">
</head>
<body>
	<div id="app"></div>
	<script src="/static/bundle.js"></script>
</body>
</html>`

var indexTemplate = template.Must(template.New("").Parse(IndexHtml))

var serveStatic = http.StripPrefix("/static/", http.FileServer(http.Dir("./dist/")))

func serveApp(w http.ResponseWriter, r *http.Request) {
	err := indexTemplate.Execute(w, nil)
	if err != nil {
		WriteResponse(w, err)
	}
}

func catchall(*http.Request, *mux.RouteMatch) bool {
	return true
}
