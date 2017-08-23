package gateway

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

type Reloadable interface {
	Reload(w http.ResponseWriter, r *http.Request)
}

type ReloadingRouter struct {
	router *mux.Router
}

// Routes maps paths to destinations.
type Routes map[string]string

func (rr *ReloadingRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rr.router.ServeHTTP(w, r)
}

func (rr *ReloadingRouter) Reload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	if err := rr.Load(name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write([]byte("Successfully registered new routes."))
	}
}

func (rr *ReloadingRouter) Load(name string) error {
	data, err := ioutil.ReadFile(path.Join(ConfigDir, name))
	if err != nil {
		return err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	rr.RegisterRoutes(config.Routes)
	return nil
}

func (rr *ReloadingRouter) RegisterRoutes(routes Routes) {
	router := mux.NewRouter()
	var root string
	var hasRoot bool
	if root, hasRoot = routes["/"]; hasRoot {
		delete(routes, "/")
	}

	router.PathPrefix("/reload/{name}").HandlerFunc(rr.Reload)
	for path, destination := range routes {
		router.PathPrefix(path).HandlerFunc(MakeHandlerTo(destination))
	}
	if hasRoot {
		router.PathPrefix("/").HandlerFunc(MakeHandlerTo(root))
	}
	rr.router = router
}
