package main

import (
	"log"
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

func init() {
	renderer = render.New()
}

func main() {
	//router
	router := mux.NewRouter()

	//index Page
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := GetCurrentUser(r)
		if user == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		renderer.HTML(w, http.StatusOK, "index", map[string]string{"userId": user.Uid})
		return
	})

	//login Page
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		renderer.HTML(w, http.StatusOK, "login", map[string]string{"message": "login이 필요합니다."})
	}).Methods("GET")

	//login POST
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		userId := r.FormValue("userId")
		password := r.FormValue("password")
		if userId == dbid && password == dbpw {
			log.Println("login Success")
			user := User{Uid: userId}
			SetCurrentUser(r, &user)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		log.Println("login Fail")
		renderer.HTML(w, http.StatusOK, "login", map[string]string{"message": "아이디 또는 암호를 확인하세요."})
		return
	}).Methods("POST")

	router.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		user := GetCurrentUser(r)
		if user != nil {
			DeleteCurrentUser(r, user)
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}).Methods("GET")

	//middleware
	n := negroni.Classic()
	store := cookiestore.New([]byte(sessionSecret))
	n.Use(sessions.Sessions(sessionKey, store))

	n.UseHandler(router)
	//run
	n.Run(":5000")
}
