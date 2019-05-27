package main

import (
	"net/http"

	"github.com/goincremental/negroni-sessions/cookiestore"

	sessions "github.com/goincremental/negroni-sessions"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var renderer *render.Render
var dbid = "myid"
var dbpw = "password"
var sessionSecret = "sessionSecret"
var sessionKey = "sessionKey"
var userSession = "myidssession"
var userSessionVal = "userSessionVal"

func init() {
	renderer = render.New()
}

func main() {
	//router
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		session := sessions.GetSession(r)
		if sessionVal := session.Get(userSession); sessionVal == userSessionVal {
			renderer.HTML(w, http.StatusOK, "index", map[string]string{"message": "환영합니다."})
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	})

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		renderer.HTML(w, http.StatusOK, "login", map[string]string{"message": "login이 필요합니다."})
	}).Methods("GET")

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("username")
		pw := r.FormValue("password")
		if id == dbid && pw == dbpw {
			//upload session
			session := sessions.GetSession(r)
			session.Set(userSession, userSessionVal)
			//redirect
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		renderer.HTML(w, http.StatusOK, "login", map[string]string{"message": "ip 혹은 pw가 틀렸습니다."})
	}).Methods("POST")

	//middleware
	n := negroni.Classic()
	store := cookiestore.New([]byte(sessionSecret))
	n.Use(sessions.Sessions(sessionKey, store))

	n.UseHandler(router)
	//run
	n.Run(":5000")
}
