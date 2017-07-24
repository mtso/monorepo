package main

import (
	"io"
	"net/http"
)

const destinationURL = "http://localhost:80"

// CopyHeader copies header values from one to another.
func CopyHeader(from http.Header, to http.Header) {
	for key, slice := range from {
		for _, value := range slice {
			to.Add(key, value)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	req, err := http.NewRequest(r.Method, destinationURL, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer r.Body.Close()

	CopyHeader(r.Header, req.Header)
	req.Header.Add("TACKDB-API-VERSION", "1")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	CopyHeader(resp.Header, w.Header())
	w.WriteHeader(resp.StatusCode)

	// Copy stream from response body into response writer.
	io.Copy(w, resp.Body)
	defer resp.Body.Close()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3750", nil)
}
