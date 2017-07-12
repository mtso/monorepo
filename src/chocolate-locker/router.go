package main

import (
	// "github.com/mtso/monorepo/src/chocolate-locker/models"
	"./models"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", models.Login).Methods("POST")
	auth.HandleFunc("/signup", models.Signup).Methods("POST")
	auth.HandleFunc("/logout", models.Logout).Methods("GET")

	return router
}
