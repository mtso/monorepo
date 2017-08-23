package gateway

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

const (
	SessionId = "sessid"
	JwtKey    = "JWT"
)

var ErrUnrecognizedToken = errors.New("Unrecognized token")
var ErrNoToken = errors.New("No token found")
var ErrLogout = errors.New("Logged out")

var userstore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

// AttachJWT adds the JWT from session store to the request.
func AttachJWT(r *http.Request) error {
	session, err := userstore.Get(r, SessionId)
	if err != nil {
		return err
	}

	if jwt, ok := session.Values[JwtKey]; ok {
		if jwtstring, ok := jwt.(string); ok {
			r.Header.Set("Authorization", jwtstring)
		} else {
			return ErrUnrecognizedToken
		}
	}
	return nil
}

// StripJWT strips the Authorization token from the header.
func StripJWT(r *http.Response) (string, error) {
	jwtstring := r.Header.Get("Authorization")
	if jwtstring == "" {
		return "", ErrNoToken
	} else if jwtstring == "LOGOUT" {
		return "", ErrLogout
	}
	r.Header.Del("Authorization")
	return jwtstring, nil
}

// MakeHandlerTo creates a handler that forwards the request to destination string.
func MakeHandlerTo(destination string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := userstore.Get(r, SessionId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := &http.Client{}

		// Clone request data.
		req, err := http.NewRequest(r.Method, destination+r.URL.Path, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// Add JWT token if any.
		AttachJWT(r)
		copyHeader(r.Header, req.Header)

		// Make proxy request.
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Remove JWT token, if any.
		// Should not respond with the same headers unless using echo server.
		switch tokenstring, err := StripJWT(resp); err {
		case ErrLogout:
			delete(session.Values, JwtKey)
			log.Println(session.Values)
			userstore.Save(r, w, session)
		case ErrNoToken:
			break
		case nil:
			session.Values[JwtKey] = tokenstring
			userstore.Save(r, w, session)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Copy header and body data
		copyHeader(resp.Header, w.Header())
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
		defer resp.Body.Close()
	}
}

func notNil(err error) bool {
	if err != nil {
		log.Println(err)
	}
	return err != nil
}

func copyHeader(from, to http.Header) {
	for key, slice := range from {
		for _, value := range slice {
			to.Add(key, value)
		}
	}
}
