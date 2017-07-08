package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/mtso/monorepo/src/elo-service/models"
)

type JSON map[string]interface{}

var ErrNoTitle = errors.New("Request is missing 'title' field")

func NewLeague(w http.ResponseWriter, r *http.Request) {
	body, err := ParseBody(r.Body)
	if err != nil {
		log.Println(err)
		WriteResponse(w, err)
		return
	}

	title, ok := body["title"]
	if !ok {
		resp := &JSON{
			"ok":      false,
			"message": ErrNoTitle,
		}
		WriteResponse(w, resp)
		return
	}

	id := models.RandomId()
	newleague, err := models.CreateLeague(id, title)
	if err != nil {
		WriteResponse(w, err)
		return
	}

	resp := &JSON{
		"ok":     true,
		"league": newleague,
	}
	WriteResponse(w, resp)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	body, err := ParseBody(r.Body)
	if err != nil {
		log.Println(err)
		WriteResponse(w, err)
		return
	}

	_, _ = body["title"]
}
