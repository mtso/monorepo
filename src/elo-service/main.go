package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/mtso/monorepo/src/elo-service/models"
)

type App struct {
	Db      *sql.DB
	Handler http.Handler
}

func newApp() App {
	db, err := models.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	r := LogHandler(NewRouter())

	return App{
		Db:      db,
		Handler: r,
	}
}

func main() {
	app := newApp()
	defer app.Db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3750"
	}

	http.ListenAndServe(":"+port, app.Handler)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello~ " + models.RandomId()))
	})

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/new", NewLeague).Methods("POST")
	api.HandleFunc("/{id:[a-f0-9]{24}}/games", GetGames).Methods("GET")
	api.HandleFunc("/{id:[a-f0-9]{24}}/players", GetPlayers).Methods("GET")
	api.HandleFunc("/{id:[a-f0-9]{24}}", AddGame).Methods("POST")

	return router
}
