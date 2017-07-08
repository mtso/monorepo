package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/mtso/monorepo/src/elo-service/models"
)

type JSON map[string]interface{}

var ErrNoTitle = errors.New("Request is missing 'title' field")

type App struct {
	Db      *sql.DB
	Handler http.Handler
}

func newApp() App {
	db, err := models.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	return App{
		Db:      db,
		Handler: NewRouter(),
	}
}

func main() {
	db, err := models.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3750"
	}

	http.ListenAndServe(":"+port, NewRouter())
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello~ " + models.RandomId()))
	})

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
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

	}).Methods("POST")

	return router
}

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
