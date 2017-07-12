package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSON map[string]interface{}

func WriteError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	log.Println(err)
	resp := &JSON{
		"ok":    false,
		"error": err.Error(),
	}
	WriteJSON(w, resp, 500)
	return true
}

func WriteJSON(w http.ResponseWriter, data interface{}, code ...int) error {
	jsondata, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsondata)
	return nil
}

func ParseBody(body io.ReadCloser) (map[string]interface{}, error) {

}
