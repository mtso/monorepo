package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, js interface{}, code ...int) {
	switch js.(type) {
	case error:
		http.Error(w, js.(error).Error(), http.StatusInternalServerError)
	default:
		break
	}

	resp, err := json.Marshal(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if len(code) > 0 {
		w.WriteHeader(code[0])
	}
	w.Write(resp)
}

func ParseBody(reader io.ReadCloser) (JSON, error) {
	decoder := json.NewDecoder(reader)
	var raw interface{}

	err := decoder.Decode(&raw)
	if err != nil {
		return nil, err
	}

	js := raw.(map[string]interface{})
	return js, nil
}
