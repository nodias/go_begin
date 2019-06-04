package main

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/unrolled/render"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var renderer *render.Render

const (
	sessionKey    = "mysessionkey"
	sessionSecret = "mysessionSecret"
)

func init() {
	renderer = render.New()
}

func main() {
	//router
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/login", login)

	//session
	server := negroni.Classic()
	store := cookiestore.New([]byte(sessionSecret))

	//middleware
	server.UseHandler(router)
	server.Use(sessions.Sessions(sessionKey, store))

	server.Run(":5000")
}

func index(w http.ResponseWriter, req *http.Request) {

	renderer.HTML(w, http.StatusOK, "index", map[string]string{"userId": "nodias"})
}

func login(w http.ResponseWriter, req *http.Request) {

	renderer.HTML(w, http.StatusOK, "login", map[string]string{"message": "Login is required"})
}

func loginPost(w http.ResponseWriter, req *http.Request) {
	// r.FormValue("userId")
	// r.FormValue("userPassword")
	// GetSession()
	// http.Redirect(w, r, "/index", http.StatusFound)
}
