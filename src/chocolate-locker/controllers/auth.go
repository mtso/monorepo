package controllers

import (
	"net/http"

	"github.com/gorilla/session"
)

// POST /api/auth/login
func Login(w http.ResponseWriter, r *http.Request) {

}

// POST /api/auth/signup
func Signup(w http.ResponseWriter, r *http.Request) {

}

// GET /api/auth/logout
func Logout(w http.ResponseWriter, r *http.Request) {
	sess, err := session.Get(r)
	if err != nil {
		WriteJSON(w, err)
	}

	resp := &JSON{
		"ok": true,
	}
	WriteJSON(w, resp)
}
