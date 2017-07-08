package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mtso/monorepo/src/elo-service/models"
)

type JSON map[string]interface{}

var ErrNoTitle = errors.New("Request is missing 'title' field")
var ErrNoWinner = errors.New("Request is missing 'winner' field")
var ErrNoLoser = errors.New("Request is missing 'loser' field")
var ErrNoId = errors.New("Path is missing 'id' param")

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

func AddGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leagueid, ok := vars["id"]
	if !ok {
		WriteResponse(w, ErrNoId, http.StatusBadRequest)
		return
	}

	body, err := ParseBody(r.Body)
	if err != nil {
		log.Println(err)
		WriteResponse(w, err)
		return
	}

	winner, ok := body["winner"]
	if !ok {
		WriteResponse(w, ErrNoWinner, http.StatusBadRequest)
		return
	}

	loser, ok := body["loser"]
	if !ok {
		WriteResponse(w, ErrNoLoser, http.StatusBadRequest)
		return
	}

	calcElo := func(w, l models.Player) int64 {
		return TransferPoints(w.Elo, l.Elo, BASE)
	}

	// If params are ok (has leagueid, winner, loser)
	// Get winner and loser
	gm, err := models.AddGame(leagueid, winner, loser, calcElo)
	if err != nil {
		WriteResponse(w, err)
		return
	}

	resp := JSON{
		"ok":   true,
		"game": gm,
	}
	WriteResponse(w, resp)
}

func GetGames(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		WriteResponse(w, ErrNoId, http.StatusBadRequest)
		return
	}

	gms, err := models.GetGames(id)
	if err != nil {
		WriteResponse(w, err, http.StatusInternalServerError)
		return
	}

	resp := JSON{
		"ok":    true,
		"games": gms,
	}
	WriteResponse(w, resp)
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		WriteResponse(w, ErrNoId, http.StatusBadRequest)
		return
	}

	players, err := models.GetPlayers(id)
	if err != nil {
		WriteResponse(w, err, http.StatusInternalServerError)
		return
	}

	resp := JSON{
		"ok":      true,
		"players": players,
	}
	WriteResponse(w, resp)
}
